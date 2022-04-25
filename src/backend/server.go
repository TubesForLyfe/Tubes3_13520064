package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
	"strconv"

	sm "backend/stringMatching"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Penyakit struct {
	NamaPenyakit string
	DNA          string
}

type HasilPrediksi struct {
	TanggalPrediksi  string
	NamaPasien       string
	PenyakitPrediksi string
	TingkatKemiripan int
	Status           int
}

func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func openDatabase() *sql.DB {
	// Open database connection.
	db, err := sql.Open("mysql", getEnv("DATABASE_USERNAME")+":"+getEnv("DATABASE_PASSWORD")+"@tcp("+getEnv("DATABASE_PORT")+")/"+getEnv("DATABASE_NAME"))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getDetailPrediction(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	string_body := string(body)
	if strings.Split(string_body, ":")[0] != "" {
		data := strings.Split(string_body, ":")[1]
		input := ""
		two_arguments := false
		for i := 1; i < len(data)-2; i++ {
			input += string(data[i])
			if string(data[i]) == " " {
				two_arguments = true
			}
		}

		db := openDatabase()
		result := []HasilPrediksi{}
		if two_arguments {
			date := strings.Split(input, " ")[0]
			disease := strings.Split(input, " ")[1]

			// Db query for hasilprediksi table
			db_result, err := db.Query("SELECT * FROM hasilprediksi WHERE TanggalPrediksi = '" + date + "' AND PenyakitPrediksi = '" + disease + "'")
			if err != nil {
				panic(err.Error())
			}

			for db_result.Next() {
				var hasil HasilPrediksi

				// Get hasil for each row
				err = db_result.Scan(&hasil.TanggalPrediksi, &hasil.NamaPasien, &hasil.PenyakitPrediksi, &hasil.TingkatKemiripan, &hasil.Status)
				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}
				// Append hasil to result
				result = append(result, hasil)
			}
		} else {
			// Db query for hasilprediksi table
			db_result, err := db.Query("SELECT * FROM hasilprediksi WHERE TanggalPrediksi = '" + input + "' OR PenyakitPrediksi = '" + input + "'")
			if err != nil {
				panic(err.Error())
			}

			for db_result.Next() {
				var hasil HasilPrediksi

				// Get hasil for each row
				err = db_result.Scan(&hasil.TanggalPrediksi, &hasil.NamaPasien, &hasil.PenyakitPrediksi, &hasil.TingkatKemiripan, &hasil.Status)
				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}
				// Append hasil to result
				result = append(result, hasil)
			}
		}
		// Convert result to []byte
		marshal, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// Send to frontend
		res.Write(marshal)
	}
}

func getDiseasePrediction(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	string_body := string(body)

	if strings.Split(string_body, ":")[0] != "" {

		output := []HasilPrediksi{}

		s1 := strings.Split(string_body, "------WebKitFormBoundary")
		s1_dna := strings.Split(s1[1], "\n")
		DNA := s1_dna[4]
		DNA = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, DNA)
		s1_nama := strings.Split(s1[2], "\n")
		Nama := s1_nama[3]
		Nama = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, Nama)
		s1_penyakit := strings.Split(s1[3], "\n")
		Penyakit := s1_penyakit[3]
		Penyakit = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, Penyakit)
		s1_tanggalsplit := strings.Split(s1[4], "\n")
		s1_tanggal := strings.Replace(s1_tanggalsplit[3], ",", "", -1)
		tanggalsplit := strings.Split(s1_tanggal, "/")
		year := tanggalsplit[2][:len(tanggalsplit[2])-1]
		Tanggal := year + "/" + tanggalsplit[0] +"/"+ tanggalsplit[1] 
		Tanggal = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, Tanggal)


		if sm.Regex(DNA) {

			db := openDatabase()
			db_result, err := db.Query("SELECT DNA FROM jenispenyakit WHERE NamaPenyakit = '" + Penyakit + "'")
			if err != nil {
				panic(err.Error())
			}
			defer db_result.Close()

			empty := true

			var pDNA string;

			for db_result.Next() {
				db_result.Scan(&pDNA)
				empty = false
			}

			if !empty {
				Percentage := sm.Lcs(DNA, pDNA)
				var stat int

				if (Percentage > 80) {
					stat = 1
				} else {
					stat = 0
				}

				db := openDatabase()
				db_result, err := db.Query("INSERT INTO hasilprediksi VALUES ('"+ Tanggal +"','"+ Nama +"','"+ Penyakit +"','" + strconv.Itoa(Percentage) + "','" + strconv.Itoa(stat) +"')")
				if err != nil {
					outputisi := HasilPrediksi{
						NamaPasien : Nama,
						PenyakitPrediksi : Penyakit,
						TanggalPrediksi : Tanggal,
						TingkatKemiripan : Percentage,
						Status : stat,
					}
	
					output = append(output, outputisi)
	
					marshal, err := json.Marshal(output)
					if err != nil {
						fmt.Println(err)
					}
	
					res.Write(marshal)
					return
				}

				defer db_result.Close()

				outputisi := HasilPrediksi{
					NamaPasien : Nama,
					PenyakitPrediksi : Penyakit,
					TanggalPrediksi : Tanggal,
					TingkatKemiripan : Percentage,
					Status : stat,
				}

				output = append(output, outputisi)

				marshal, err := json.Marshal(output)
				if err != nil {
					fmt.Println(err)
				}

				res.Write(marshal)
			} else {
				outputisi := HasilPrediksi{
					NamaPasien : Nama,
					PenyakitPrediksi : Penyakit,
					TanggalPrediksi : Tanggal,
					TingkatKemiripan : 0,
					Status : -2,
				}
	
				output = append(output, outputisi)
	
				marshal, err := json.Marshal(output)
				if err != nil {
					fmt.Println(err)
				}
	
				res.Write(marshal)
			}
			
		} else {

			outputisi := HasilPrediksi{
				NamaPasien : Nama,
				PenyakitPrediksi : Penyakit,
				TanggalPrediksi : Tanggal,
				TingkatKemiripan : 0,
				Status : -1,
			}

			output = append(output, outputisi)

			marshal, err := json.Marshal(output)
			if err != nil {
				fmt.Println(err)
			}

			res.Write(marshal)
		}

	}
}

func submitDisease(res http.ResponseWriter, req *http.Request) {
	// setupResponse(&res, req)

	// decoder := json.NewDecoder(req.Body)

	// var data Penyakit
	// err := decoder.Decode(&data)
	// if err != nil {
	// 	panic(err)
	// }

	// // NamaPenyakit := data.NamaPenyakit
	// // DNA := data.DNA

	// db := openDatabase()
	// // query := ("INSERT INTO jenispenyakit VALUES(" + "'" + NamaPenyakit + "', '" + DNA + "')")
	// query := ("INSERT INTO jenispenyakit VALUES('test', 'test')")

}

func main() {
	sm.BoyerMoore("a pattern matching algorithm", "rithm")
	sm.BoyerMoore("abacaabadcabacabaabb", "abacab")

	// Server
	http.HandleFunc(getEnv("BASE_PORT")+"/get-detailprediction", getDetailPrediction)
	http.HandleFunc(getEnv("BASE_PORT")+"/get-diseaseprediction", getDiseasePrediction)
	http.HandleFunc(getEnv("BASE_PORT")+"/submitDisease", submitDisease)

	fmt.Println("Starting server at port " + getEnv("BACKEND_PORT"))
	if err := http.ListenAndServe(":"+getEnv("BACKEND_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

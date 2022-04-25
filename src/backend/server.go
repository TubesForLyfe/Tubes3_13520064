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

		s1 := strings.Replace(string_body, "{", "", -1)
		s2 := strings.Replace(s1, "}", "", -1)
		s3 := strings.Replace(s2, `"`, "", -1)
		data := strings.Split(s3, ",")

		output := []HasilPrediksi{}

		FilePath := "../../test/" + strings.Split(data[1], ":")[1]

		buf, err := ioutil.ReadFile(FilePath)

		if err != nil {
			log.Fatalln(err)
		}

		isi := string(buf)

		if sm.Regex(isi) {
			outputisi := HasilPrediksi{
				NamaPasien:       strings.Split(data[0], ":")[1],
				PenyakitPrediksi: strings.Split(data[2], ":")[1],
				TanggalPrediksi:  strings.Split(data[3], ":")[1],
				TingkatKemiripan: 80,
				Status:           1,
			}

			output = append(output, outputisi)

			marshal, err := json.Marshal(output)
			if err != nil {
				fmt.Println(err)
			}

			res.Write(marshal)
		} else {
			outputisi := HasilPrediksi{
				NamaPasien:       strings.Split(data[0], ":")[1],
				PenyakitPrediksi: strings.Split(data[2], ":")[1],
				TanggalPrediksi:  strings.Split(data[3], ":")[1],
				TingkatKemiripan: 0,
				Status:           -1,
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

	var check bool = sm.Regex("AGTC")
	if check {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}
	// Server
	http.HandleFunc(getEnv("BASE_PORT")+"/get-detailprediction", getDetailPrediction)
	http.HandleFunc(getEnv("BASE_PORT")+"/get-diseaseprediction", getDiseasePrediction)
	http.HandleFunc(getEnv("BASE_PORT")+"/submitDisease", submitDisease)

	fmt.Println("Starting server at port " + getEnv("BACKEND_PORT"))
	if err := http.ListenAndServe(":"+getEnv("BACKEND_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

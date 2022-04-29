# Tugas Besar 3 IF2211 Strategi Algoritma 2021/2022
##	Kelompok 16 - GuysNamanyaMauApa?
+ Ziyad Dhia Rafi 13520064
+ Vincent Christian Siregar 13520136
+ Willy Wilsen 13520160

### Tentang Proyek ini
Ini adalah sebuah proyek yang bertujuan untuk menuntaskan tugas mata kuliah Strategi Algoritma. Repository ini berisi source code windows form app untuk melakukan pencarian file pada suatu directory

Ini adalah proyek yang bertujuan untuk menuntaskan tugas mata kuliah Strategi Algoritma. Repository ini berisi source code backend dan frontend untuk pembuatan aplikasi berbasis web DNA pattern matching dengan menggunakan algoritma KMP, Booyer Moore dan Regex.

### Backend
Backend bertujuan untuk menjalankan server yang terhubung dengan database, serta algoritma-algoritma untuk pattern matching. Backend ditulis dalam bahasa GO.

### Frontend
Frontend bertujuan untuk memberikan antarmuka yang dapat digunakan pengguna untuk menggunakan aplikasi berbasis web. Ditulis dalam javascript menggunakan framework frontend React js.

### Fitur-fitur
1. Menambahkan penyakit baru ke dalam database
2. Melakukan prediksi penyakit dengan input DNA pasien, dan memberikan hasil dengan persen kecocokan tertentu
3. Melakukan pencarian terhadap tes (prediksi penyakit) yang telah dilakukan sebelumnya

### Cara menjalankan Program & Requirement
#### Melalui Website
##### Cara menjalankan aplikasi
1. Buka https://tesdnanamanyamauapa.netlify.app/ pada web browser.
#### Local
##### Requirements
1. Install Go Lang pada https://go.dev/
2. Install NodeJs pada https://nodejs.org/
##### Cara menjalankan aplikasi
1. Buka directory frontend pada terminal,
```
cd src/frontend/
```
2. Install modul
```
npm install
```
3. Jalankan frontend
```
npm start
```
4. Aplikasi akan terbuka secara otomatis, atau buka pada web browser, localhost:3000
5. Buka directory backend pada terminal, 
```
cd ../backend/
```
6. Jalankan server.go dengan command:
```
go run server.go
```
7. Aplikasi DNA pattern matching sudah dapat digunakan sepenuhnya.




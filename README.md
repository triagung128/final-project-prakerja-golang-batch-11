# Final Project Pelatihan Kartu Prakerja (Alterra Academy - Golang Batch 11)

## Task
Membuat sebuah aplikasi REST API dengan Golang dengan ketentuan berikut :
1. Carilah ide dengan tema bebas.
2. Buatlah API untuk tema tersebut.
3. API dapat mengeluarkan data.
4. Setidaknya memiliki 2 endpoint dari API yang kalian buat.
5. Gunakanlah error handling, mvc, dan materi lain untuk pembuatannya.
6. Gunakan postman untuk menguji API yang kalian buat.
7. Menggunakan version control GIT untuk management tugas.
8. Batas pengerjaan adalah 3 hari setelah pelatihan berakhir.
9. Dikerjakan secara mandiri.
10. Lakukan deployment.

## Features
- [x] Menampilkan data kursus.
- [x] Login.
- [x] Register.
- [x] User melakukan enrolment pada kursus yang dipilih.
- [x] User dapat melihat daftar list enrolment.
- [x] User yang sudah melakukan enrolment dapat melakukan review.
- [x] User dapat melihat daftar list review kursus.

## Endpoint
- Mendapatkan data kursus pelatihan.
```
GET /courses
```
- Mendaftarkan user baru.
```
POST /register
```
- Login untuk user.
```
POST /login
```
- Menambahkan enrollment.
```
POST /enrollments
```
- Mendapatkan data enrollment.
```
GET /enrollments
```
- Menambahkan review pada kursus pelatihan.
```
POST /reviews
```
- Mendapatkan data review kursus pelatihan.
```
GET /reviews
```

## Built with
* [Golang](https://go.dev/)
* [Echo](https://echo.labstack.com/)
* [Gorm](https://gorm.io/index.html)
* [MySQL](https://www.mysql.com/)
* [Bcrypt (Hashing)](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
* [Godotenv](https://github.com/joho/godotenv)
* [Validator](https://github.com/go-playground/validator)
* [JWT Token](https://jwt.io/)
* [Scalingo](https://scalingo.com/)
* [DB4Free](https://www.db4free.net/)
* MVC Pattern

## Documentation
Postman : https://documenter.getpostman.com/view/14265076/2s9YJaYPbZ

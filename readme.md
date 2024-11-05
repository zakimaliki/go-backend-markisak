# MyApp

MyApp adalah aplikasi sederhana yang dibangun menggunakan Go dan Gofiber. Aplikasi ini menyediakan command line interface untuk menjalankan migrasi dan seeder database.

## Prerequisites

Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/) (versi terbaru)
- [Gofiber](https://gofiber.io/docs) (dapat diinstal melalui `go get`)

## Instalasi

1. Clone repositori ini:  
   ```bash
   git clone https://github.com/zakimaliki/go-backend-markisak
   cd go-backend-markisak
   ```

2. Install dependensi yang diperlukan:  
   ```bash
   go mod tidy
   ```

3. Buat file `.env` di direktori root proyek Anda dan tambahkan konfigurasi yang diperlukan untuk koneksi database

## Menjalankan Aplikasi

Untuk menjalankan aplikasi, Anda dapat menggunakan perintah berikut:

- Untuk menjalankan server:  
  ```bash
  go run main.go
  ```

- Untuk menjalankan migrasi:  
  ```bash
  go run main.go migrate
  ```

- Untuk menjalankan seeder:  
  ```bash
  go run main.go seed
  ```

## Struktur Proyek

- `main.go`: File utama yang berisi logika aplikasi dan command line interface.
- `src/configs`: Konfigurasi aplikasi, termasuk pengaturan database.
- `src/helpers`: Fungsi-fungsi utilitas, termasuk migrasi dan seeder.
- `src/routes`: Definisi rute untuk aplikasi Gofiber.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan buat pull request atau buka issue untuk diskusi.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).

## About

Proyek ini dapat diakses di [GitHub Repository](https://github.com/zakimaliki/go-backend-markisak).
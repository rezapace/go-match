# Mini-Project-PongPedia

Mini Project Alterra Academy

## Deskripsi

PongPedia adalah aplikasi web untuk mengelola turnamen tenis meja. Aplikasi ini menggunakan framework Echo untuk membangun server web dan GORM untuk interaksi dengan database. Fitur-fitur yang disediakan meliputi autentikasi pengguna, manajemen turnamen, pertandingan, dan profil pengguna, serta menyediakan rute yang berbeda untuk fungsi admin dan pengguna.

## Fitur

- **Autentikasi Pengguna**: Registrasi dan login pengguna.
- **Manajemen Turnamen**: Membuat, memperbarui, dan menghapus turnamen.
- **Manajemen Pertandingan**: Membuat dan memperbarui pertandingan.
- **Dashboard Admin**: Menampilkan statistik pengguna, pemain, turnamen, dan pertandingan.
- **Manajemen Pengguna**: Melihat dan mengelola data pengguna.

## Struktur Proyek

- **main.go**: Titik masuk utama aplikasi yang menginisialisasi database dan mengatur rute.
- **repository**: Berisi logika untuk interaksi dengan database.
- **controllers**: Berisi endpoint untuk admin dan pengguna.
- **usecase**: Berisi logika bisnis untuk berbagai fungsi.
- **routes**: Mendefinisikan endpoint API.

## Instalasi

1. Clone repositori ini:
    ```sh
    git clone https://github.com/username/Mini-Project-PongPedia.git
    ```
2. Masuk ke direktori proyek:
    ```sh
    cd Mini-Project-PongPedia
    ```
3. Instal dependensi:
    ```sh
    go mod tidy
    ```
4. Konfigurasi database di `config/config.go`.

## Menjalankan Aplikasi

1. Jalankan aplikasi:
    ```sh
    go run main.go
    ```
2. Akses aplikasi di `http://localhost:8080`.

## Rute API

- **Autentikasi**
  - `POST /register`: Registrasi pengguna baru.
  - `POST /login`: Login pengguna.

- **Admin**
  - `GET /admin/dashboard`: Mendapatkan dashboard admin.
  - `GET /admin/users`: Mendapatkan semua pengguna.
  - `POST /admin/tournaments`: Membuat turnamen baru.
  - `PUT /admin/tournaments/:id`: Memperbarui turnamen.
  - `POST /admin/matches`: Membuat pertandingan baru.
  - `PUT /admin/matches/:id`: Memperbarui pertandingan.

## Kontribusi

1. Fork repositori ini.
2. Buat branch fitur baru (`git checkout -b fitur-baru`).
3. Commit perubahan (`git commit -m 'Menambahkan fitur baru'`).
4. Push ke branch (`git push origin fitur-baru`).
5. Buat Pull Request.

## Lisensi

Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.

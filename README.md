# RESTful News & Topics Management
Layanan service API untuk mengelola data topik dan berita.

### Layanan API Server
-    CRUD Topik
-    CRUD Berita
-    Pencarian daftar berita berdasarkan topik
-    Pencarian daftar berita berdasarkan status draft, deleted, atau publish.
-    Menampilkan berita berdasarkan slug (identitas set endpoint berita)

### Inisialisasi API Server
Database yang digunakan `PostgreSQL`, cache `Redis`, ORM menggunakan pustaka `gorm` dan web framework menggunakan ```echo```. Sebelum memulai copy atau rename file ```.env_example``` menajdi ```.env``` dan sesuaikan dengan konfigurasi yang digunakan. Stuktur database sudah otomatis ter-migrasi ke database yang sudah diset di file ```.env``` .

### Persiapan Menjalankan API Server
-    Database sudah disiapkan dalam container docker, pastikan anda sudah meng-install docker sebelum menjalankan perintah ini.
     ```console
     docker-compose up -d
     ```

-    Menjalankan service API
     ```console
     go run .
     ```

     Service API secara default berjalan di: ```http://localhost:8000/v1```

### Melakukan pengujian API server
Untuk dokumentasi singkat penggunaan API server yang berkaitan dengan nama end point, parameter dan respon data dapat dilihat pada ```postman_collection.json``` .

### Kontak
Widya Ade Bagus - https://www.linkedin.com/in/widya-ade-bagus-3a660716b/

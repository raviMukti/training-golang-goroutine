## NOTE

1. Buat project menjadi go module dengan perintah `go mod init training-golang-goroutine`
2. Menjalankan function secara async adalah dengan menambahkan perintah go di depan function contoh `go RunHelloWorld()`
3. Tapi hati - hati ketika aplikasi program kita sudah selesai, tapi go routine kita belum selesai, bisa jadi function async itu tidak pernah akan menghasilkan sesuatu
4. Channel merupakan tempat komunikasi secara synchronous yang bisa dilakukan oleh GoRoutine (Pengirim-Penerima)
5. Channel hanya bisa menampung 1 data (satu jenis), channel harus di close jika tidak digunakan, jika tidak bisa menyebabkan memory leak
6. Tipe data Channel adalah `chan` untuk membuat channel kita bisa menggunakan perintah `make()` dan kita harus menentukan tipe data inputan `channel := make(chan string)`
7. Deadlock akan terjadi ketika tidak ada yg mengirim dan menerima data
8. Kita bisa menandai channel tersebut sebagai parameter in dan out menggunakan `chan<-` untuk in dan `<-chan` untuk out
9. Buffered Channel merupakan buffer yang bisa digunakan untuk menampung data antrian di channel, agar tidak terjadi proses yang menjadi lambat ketika mengirim atau menerima data di channel
10. Untuk membuat Buffered Channel menggunakan perintah `channel := make(chan string, 3)` 3 sebagai ukuran buffer, `cap(channel)` untuk melihat panjang buffer dan `len(channel)` untuk melihat jumlah data di channel
11. Range Channel, kadang - kadang ada sebuah case dimana sebuah channel dikirim data secara terus menerus, dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data, untuk case ini maka yang bisa kita lakukan adalah menggunakan perulangan range ketika menerima data dari channel
12. Ketika kita ingin mendapatkan data dari semua channel, kita bisa melakukannya dengan Select Channel, dengan hal ini kita bisa memilih data tercepat dari beberapa channel, jika datang secara bersamaan di beberapa channel, maka akan dipilih secara random
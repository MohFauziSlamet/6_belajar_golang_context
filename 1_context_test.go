package belajargolangcontext

import (
	"context"
	"fmt"
	"testing"
)

/*
Pengenalan Context
● Context merupakan sebuah data yang membawa value, sinyal cancel,
sinyal timeout dan sinyal deadline.
● Context biasanya dibuat per request (misal setiap ada request
masuk ke server web melalui http request).
● Context digunakan untuk mempermudah kita meneruskan value,
dan sinyal antar proses.
*/

/*
Kenapa Context Perlu Dipelajari?
● Context di Golang biasa digunakan untuk mengirim data request
atau sinyal ke proses lain.
● Dengan menggunakan context, ketika kita ingin membatalkan
semua proses, kita cukup mengirim sinyal ke context, maka secara
otomatis semua proses akan dibatalkan.
● Hampir semua bagian di Golang memanfaatkan context,
seperti database, http server, http client, dan lain-lain.
● Bahkan di Google sendiri, ketika menggunakan Golang,
context wajib digunakan dan selalu dikirim ke setiap function
yang dikirim.
*/

/*
Package Context
● Context direpresentasikan di dalam sebuah interface Context.
● interface Context terdapat dalam package context.
● https://golang.org/pkg/context/
*/

/*
Membuat Context
● Karena Context adalah sebuah interface, untuk membuat context
kita butuh sebuah struct yang sesuai dengan kontrak interface
Context.
● Namun kita tidak perlu membuatnya secara manual.
● Di Golang package context terdapat function yang bisa kita
gunakan untuk membuat Context.
*/

/*
Function Membuat Context
Function
context.Background() :  Membuat context kosong. Tidak pernah
dibatalkan, tidak pernah timeout, dan tidak memiliki value
apapun. Biasanya digunakan di main function atau dalam test,
atau dalam awal proses request terjadi.

context.TODO() :  Membuat context kosong seperti
Background(), namun biasanya menggunakan ini ketika belum jelas
context apa yang ingin digunakan.
*/

func TestContext(t *testing.T) {
	backgroud := context.Background()
	todo := context.TODO()

	fmt.Printf("backgroud: %v\n", backgroud)
	fmt.Printf("todo: %v\n", todo)
}

/*
Parent dan Child Context
● Context menganut konsep parent dan child
● Artinya, saat kita membuat context, kita bisa membuat child
context dari context yang sudah ada.
● Parent context bisa memiliki banyak child, namun child hanya
bisa memiliki satu parent context.
● Konsep ini mirip dengan pewarisan di pemrograman
berorientasi object.
*/

/*
Hubungan Antara Parent dan Child Context
● Parent dan Child context akan selalu terhubung.
● Saat nanti kita melakukan misal pembatalan context A, maka
semua child dan sub child dari context A akan ikut dibatalkan.
● Namun jika misal kita membatalkan context B, hanya context B
dan semua child dan sub child nya yang dibatalkan, parent
context B tidak akan ikut dibatalkan.
● Begitu juga nanti saat kita menyisipkan data ke dalam context A,
semua child dan sub child nya bisa mendapatkan data tersebut.
● Namun jika kita menyisipkan data di context B, hanya context B
dan semua child dan sub child nya yang mendapat data, parent
context B tidak akan mendapat data.
*/

/*
Immutable
● Context merupakan object yang Immutable, artinya setelah 
Context dibuat, dia tidak bisa diubah lagi.
● Ketika kita menambahkan value ke dalam context, atau 
menambahkan pengaturan timeout dan yang lainnya, secara 
otomatis akan membentuk child context baru, bukan merubah 
context tersebut.
*/

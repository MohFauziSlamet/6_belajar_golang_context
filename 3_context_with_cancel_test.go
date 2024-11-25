package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
Context With Cancel
● Selain menambahkan value ke context, kita juga bisa menambahkan
sinyal cancel ke context.
● Kapan sinyal cancel diperlukan dalam context?.
● Biasanya ketika kita butuh menjalankan proses lain, dan kita ingin
bisa memberi sinyal cancel ke proses tersebut.
● Biasanya proses ini berupa goroutine yang berbeda, sehingga dengan
mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim
sinyal cancel ke context nya.
● Namun ingat, goroutine yang menggunakan context, tetap harus melakukan
pengecekan terhadap context nya, jika tidak, tidak ada gunanya.
● Untuk membuat context dengan cancel signal, kita bisa menggunakan
function context.WithCancel(parent).
*/

func CreateCounterGoroutineLeak() chan int {
	destinationChannel := make(chan int)

	/// make goroutine
	go func() {
		/// close chanel , when all process done
		defer close(destinationChannel)

		/// looping terus menerus.
		counter := 1
		for {
			destinationChannel <- counter
			counter++
		}
	}()

	return destinationChannel
}
func TestContextGoroutineLeak(t *testing.T) {
	/// membuat goroutine leak
	/// start
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	destination := CreateCounterGoroutineLeak()

	for n := range destination {
		fmt.Printf("Counter : %v\n", n)

		/// kita hentikan looping pada createCounter
		if n == 10 {
			break
		}
	}

	fmt.Printf("end runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 3
	/// end
	/*
		/ Apa itu Goroutine Leak?
		Goroutine leak terjadi ketika goroutine (unit eksekusi ringan dalam Go, mirip dengan
		thread) dibiarkan berjalan terus-menerus atau menunggu tanpa pernah dihentikan atau
		diambil hasilnya. Ini menyebabkan penggunaan sumber daya (seperti memori) yang tidak
		efisien dan akhirnya dapat menguras sumber daya sistem.

		/ Penyebab Goroutine Leak
		1. Goroutine yang tidak pernah berakhir:
		Goroutine yang ditunggu tanpa batas waktu atau tidak memiliki kondisi penghentian
		yang jelas.
		2. Channel yang tidak pernah ditutup:
		Goroutine menunggu pada channel yang tidak pernah menerima data atau tidak pernah
		ditutup, sehingga goroutine tetap dalam keadaan menunggu.
		3. Tidak memeriksa atau mengabaikan error:
		Tidak memeriksa error dari operasi yang bisa memblokir goroutine (seperti operasi
		jaringan atau I/O).
		Lupa mengonsumsi channel:
		Membuat goroutine untuk mengirim data ke channel tetapi tidak ada yang menerima data
		tersebut.

		/ Akibat Goroutine Leak
		1. Konsumsi Memori Berlebih:
		Goroutine yang bocor tetap memakan memori, dan semakin banyak goroutine yang bocor,
		semakin besar konsumsi memori.
		2. Penurunan Performa:
		Goroutine yang bocor dapat membuat aplikasi menjadi lambat karena banyaknya goroutine
		yang aktif.
		Kehabisan Sumber Daya Sistem:
		Jika cukup banyak goroutine yang bocor, aplikasi dapat kehabisan file descriptors,
		network connections, atau memori, yang akhirnya dapat menyebabkan aplikasi crash.

		/ Cara Mencegah Goroutine Leak
		Pastikan goroutine memiliki cara untuk keluar:
		Gunakan context dengan timeout atau deadline.
		Pastikan goroutine keluar jika tidak diperlukan lagi.
		Tutup channel:
		Selalu tutup channel setelah tidak lagi digunakan untuk memberi tahu goroutine
		penerima bahwa tidak ada lagi data yang akan datang.
		Konsumsi channel dengan benar:
		Pastikan ada penerima untuk setiap channel yang dikirim data oleh goroutine.
		Gunakan WaitGroup:
		Untuk memastikan goroutine selesai sebelum fungsi keluar.
	*/
}

func CreateCounterWithContext(ctx context.Context) chan int {
	destinationChannel := make(chan int)

	/// make goroutine
	go func() {
		/// close chanel , when all process done
		defer close(destinationChannel)

		/// looping terus menerus.
		counter := 1
		for {

			/// untuk menghentikan looping ini ,
			/// kita akan check dari context.
			/// jika context sudah done() , kita hentikan go func ini ,
			/// bukan menghentikan loopingnya. ini dilakukan
			/// agar tidak terjadi goroutine leak.

			select {
			case <-ctx.Done():
				return // stop this go func

			default: // send data to chanel
				destinationChannel <- counter
				counter++
			}

		}
	}()

	return destinationChannel
}

func awaitWithSecond(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}

func TestContextWithCancel(t *testing.T) {
	/// start
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	/// create parent context dan cancel context
	parentCtx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(parentCtx)

	destination := CreateCounterWithContext(cancelCtx)
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	for n := range destination {
		fmt.Printf("Counter : %v\n", n)

		/// kita hentikan looping pada createCounter
		if n == 10 {
			break
		}
	}

	cancelFunc()

	awaitWithSecond(1)

	fmt.Printf("end runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 3
	/// end

}

/// ==================================================================

/*
/ Context With Timeout
● Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa
menambahkan sinyal cancel ke context secara otomatis dengan menggunakan pengaturan
timeout.
● Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel
secara manual, cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati.
● Penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query
ke database atau http api, namun ingin menentukan batas maksimal timeout nya.
● Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout,
kita bisa menggunakan function context.WithTimeout(parent, duration).

*/

func CreateCounterContextWithTimeOut(ctx context.Context) chan int {
	destinationChannel := make(chan int)

	/// make goroutine
	go func() {
		/// close chanel , when all process done
		defer close(destinationChannel)

		/// looping terus menerus.
		counter := 1
		for {

			/// untuk menghentikan looping ini ,
			/// kita akan check dari context.
			/// jika context sudah done() , kita hentikan go func ini ,
			/// bukan menghentikan loopingnya. ini dilakukan
			/// agar tidak terjadi goroutine leak.

			select {
			case <-ctx.Done():
				return // stop this go func

			default: // send data to chanel
				destinationChannel <- counter
				counter++
				awaitWithSecond(1) // simulate slow process
			}

		}
	}()

	return destinationChannel
}

func TestContextWithTimeOut(t *testing.T) {
	/// start
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	/// create parent context dan cancel context
	parentCtx := context.Background()
	cancelCtx, cancelFunc := context.WithTimeout(parentCtx, 5*time.Second)

	/// stop the process if , time is timeout.
	defer cancelFunc()

	destination := CreateCounterContextWithTimeOut(cancelCtx)
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	for n := range destination {
		fmt.Printf("Counter : %v\n", n)

	}

	fmt.Printf("end runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 3
	/// end

}

/// ==================================================================

/*
/ Context With Deadline
● Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa
menggunakan deadline.
● Pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu
dari sekarang, kalo deadline ditentukan kapan waktu timeout nya, misal jam 12 siang
hari ini.
● Untuk membuat context dengan cancel signal secara otomatis menggunakan deadline,
kita bisa menggunakan function context.WithDeadline(parent, time).
*/

func TestContextWithDeadLine(t *testing.T) {
	/// start
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	/// create parent context dan cancel context
	parentCtx := context.Background()
	cancelCtx, cancelFunc := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))

	/// stop the process if , time is timeout.
	defer cancelFunc()

	destination := CreateCounterContextWithTimeOut(cancelCtx)
	fmt.Printf("start runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 2

	for n := range destination {
		fmt.Printf("Counter : %v\n", n)

	}

	fmt.Printf("end runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // result = 3
	/// end

}



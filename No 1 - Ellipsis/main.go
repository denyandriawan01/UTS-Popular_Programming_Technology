/* ===========================================================
 ____  ____  _  _  ___  ____  ____  ____  ____    __    _  _ 
(  _ \( ___)( \( )/ __)( ___)(  _ \(_  _)(_  _)  /__\  ( \( )
 )___/ )__)  )  (( (_-. )__)  )   /  )(   _)(_  /(__)\  )  ( 
(__)  (____)(_)\_)\___/(____)(_)\_) (__) (____)(__)(__)(_)\_)

   =========================================================== */

/* Ellipsis pada Golang adalah fitur yang memungkinkan sebuah fungsi menerima argumen sebanyak apapun tanpa batasan. Di Golang, ellipsis ditandai dengan tiga titik (...), dan ditempatkan sebelum tipe data parameter terakhir pada definisi fungsi.

Dengan menggunakan ellipsis, kita dapat memanggil fungsi dengan sejumlah argumen yang berbeda-beda, mulai dari satu sampai banyak. Kemudian, semua argumen ini akan dikumpulkan dalam bentuk slice di dalam fungsi tersebut. */


/* ===========================================================
  ___  __   __ _  ____  __   _  _ 
 / __)/  \ (  ( \(_  _)/  \ / )( \
( (__(  O )/    /  )( (  O )) __ (
 \___)\__/ \_)__) (__) \__/ \_)(_/

   =========================================================== */

package main

import "fmt"

func hitungTotalHarga(items ...int) float64 {
	total := 0
	for _, item := range items {
		total += item
	}
	return float64(total)
}

func main() {
	var daftarHarga []int
	var daftarBanyak []int
	var x int

	fmt.Println("--------------------------------------------------   ")
	fmt.Println("   ____  __    __    ____  ____  ___  ____  ___      ")
	fmt.Println("  ( ___)(  )  (  )  (_  _)(  _ \\/ __)(_  _)/ __)    ")
	fmt.Println("   )__)  )(__  )(__  _)(_  )___/\\__ \\ _)(_ \\__ \\ ")
	fmt.Println("  (____)(____)(____)(____)(__)  (___/(____)(___/     ")
	fmt.Println(" ")
	fmt.Println("--------------------------------------------------   ")
	fmt.Printf("[!] Masukkan jumlah barang: ")
	fmt.Scanln(&x)
	fmt.Println("--------------------------------------------------   ")

	for i := 0; i < x; i++ {
		var harga, jumlah int
		fmt.Printf(" | [>] Masukkan harga barang ke-%d   : ", i+1)
		fmt.Scanln(&harga)
		fmt.Printf(" | [>] Masukkan jumlah barang ke-%d  : ", i+1)
		fmt.Scanln(&jumlah)

		daftarHarga  = append(daftarHarga, harga)
		daftarBanyak = append(daftarBanyak, jumlah) 
	}

	var items []int
	for i := 0; i < x; i++ {
		items = append(items, daftarHarga[i]*daftarBanyak[i])
	}

	totalHarga := hitungTotalHarga(items...)

	fmt.Println("--------------------------------------------------   ")
	fmt.Printf("[!] Total harga: Rp%.2f\n", totalHarga)
	fmt.Println("--------------------------------------------------   ")
}
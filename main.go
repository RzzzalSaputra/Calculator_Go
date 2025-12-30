package main

import "fmt"

func main() {
	// Variabel untuk menyimpan pilihan menu pengguna
	var pilihanMenu int
	// Variabel untuk mengontrol loop program (lanjut atau berhenti)
	var lanjutLoop string
	lanjutLoop = "y"

	// Loop utama program, berjalan selama input bukan "n"
	for lanjutLoop != "n" {
		fmt.Println("Pilihan Menu Kalkulator")
		fmt.Println("1. Pertambahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("")
		fmt.Print("Masukkan Pilihan Menu: ")

		fmt.Scan(&pilihanMenu) // Membaca input menu dari pengguna

		// Memilih operasi berdasarkan input menu
		switch pilihanMenu {
		case 1:
			prosesInputDanHitung(add)
		case 2:
			prosesInputDanHitung(sub)
		case 3:
			prosesInputDanHitung(multi)
		case 4:
			prosesInputDanHitung(div)
		default:
			fmt.Println("Menu Pilihan Tidak ada")
		}

		fmt.Print("Apakah Ingin Lanjut atau Tidak (Ketik (y) jika lanjut (n) jika tidak): ")
		fmt.Scan(&lanjutLoop) // Membaca input untuk melanjutkan atau berhenti

	}
}

// Fungsi helper untuk mengambil input angka dan menjalankan fungsi operasi
func prosesInputDanHitung(f func(int, int)) {
	var angka1, angka2 int

	fmt.Print("Masukkan Angka 1: ")
	fmt.Scan(&angka1)

	fmt.Print("Masukkan Angka 2: ")
	fmt.Scan(&angka2)

	f(angka1, angka2)
}

// Fungsi penjumlahan
func add(angka1, angka2 int) {
	fmt.Println("Hasil Pertambahan ", angka1, " + ", angka2, "= ", angka1+angka2)
}

// Fungsi pengurangan
func sub(angka1, angka2 int) {
	fmt.Println("Hasil Pengurangan ", angka1, " - ", angka2, "= ", angka1-angka2)
}

// Fungsi perkalian
func multi(angka1, angka2 int) {
	fmt.Println("Hasil Perkalian ", angka1, " x ", angka2, "= ", angka1*angka2)
}

// Fungsi pembagian
func div(angka1, angka2 int) {
	fmt.Println("Hasil Pembagian ", angka1, " : ", angka2, "= ", angka1/angka2)
}

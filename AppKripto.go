package main

import (
	"fmt"
)

const NMAX int = 100

type aset struct {
	nama, kode             string
	harga, nominal, jumlah float64
}

type kripto struct {
	nama, kode string
	harga      float64
}

type transaksi struct {
	tipe, nama, kode     string
	harga, total, jumlah float64
}

type arrKripto [NMAX]kripto
type arrAset [NMAX]aset
type arrTransaksi [NMAX]transaksi

func main() {
	var pilihanMenu int = 0
	menuUtama(&pilihanMenu)
}

// Bagian Nakulla

func menuUtama(pilihanMenu *int) {
	var jumlahAsetUser, jumlahTransaksiUser int
	var asetUser arrAset
	var riwayatUser arrTransaksi
	var saldoUser float64
	var data arrKripto

	isiDataKripto(&data)

	for *pilihanMenu != 5 {
		fmt.Println()
		fmt.Println("    ===== Selamat Datang =====")
		fmt.Println()
		fmt.Println("Aplikasi Simulasi Perdagangan Kripto")
		fmt.Println("------------------------------------")
		fmt.Println("1. Transaksi Kripto")
		fmt.Println("2. Riwayat Transaksi")
		fmt.Println("3. Pencarian Aset Kripto")
		fmt.Println("4. Cek Portofolio")
		fmt.Println("5. Keluar")
		fmt.Println("------------------------------------")
		fmt.Print("Pilihan Anda (1/2/3/4/5)? ")
		fmt.Scan(pilihanMenu)

		if *pilihanMenu == 1 {
			transaksiKripto(data, &asetUser, &riwayatUser, &saldoUser, &jumlahAsetUser, &jumlahTransaksiUser)
		} else if *pilihanMenu == 2 {
			if jumlahTransaksiUser == 0 {
				fmt.Println("Anda belum melakukan transaksi, silakan transaksi terlebih dahulu.")
			} else {
				tampilkanRiwayat(&riwayatUser, jumlahTransaksiUser)
			}
		} else if *pilihanMenu == 3 {
			if jumlahAsetUser == 0 {
				fmt.Println("Anda belum memiliki aset, silakan beli terlebih dahulu.")
			} else {
				mencariAset(data, &asetUser, &jumlahAsetUser)
			}
		} else if *pilihanMenu == 4 {
			if jumlahAsetUser == 0 {
				fmt.Println("Anda belum memiliki aset, silakan beli dan jual terlebih dahulu.")
			} else {
				menampilkanPorto(&asetUser, &jumlahAsetUser)
			}
		} else {
			fmt.Println("Tolong masukkan angka yang sesuai!")
		}
	}
	fmt.Println("------------------------------------")
	fmt.Println("Kamu sudah keluar dari aplikasi ini")
	fmt.Println()
	fmt.Println("     ===== Terima Kasih =====")
	fmt.Println()
}

// Bagian Nakulla

func transaksiKripto(data arrKripto, asetUser *arrAset, riwayatUser *arrTransaksi, saldoUser *float64, jumlahAsetUser, jumlahTransaksiUser *int) {
	var PilihanMenuTransaksi int = 0

	for PilihanMenuTransaksi != 5 {
		fmt.Println()
		fmt.Println("---------------------------------")
		fmt.Println("   Apa yang ingin anda lakukan?")
		fmt.Println("---------------------------------")
		fmt.Println("1. Beli Aset Kripto")
		fmt.Println("2. Jual Aset Kripto")
		fmt.Println("3. Cek Saldo")
		fmt.Println("4. Isi Saldo")
		fmt.Println("5. Kembali")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan Anda (1/2/3/4/5)? ")
		fmt.Scan(&PilihanMenuTransaksi)

		if PilihanMenuTransaksi == 1 {
			if *saldoUser == 0 {
				fmt.Println("Saldo Anda masih kosong, silakan isi saldo terlebih dahulu.")
			} else {
				beliAsetKripto(data, asetUser, riwayatUser, saldoUser, jumlahAsetUser, jumlahTransaksiUser)
			}
		} else if PilihanMenuTransaksi == 2 {
			if *jumlahAsetUser == 0 {
				fmt.Println("Anda belum memiliki aset, silakan beli terlebih dahulu.")
			} else {
				jualAsetKripto(data, asetUser, riwayatUser, saldoUser, jumlahAsetUser, jumlahTransaksiUser)
			}
		} else if PilihanMenuTransaksi == 3 {
			fmt.Printf("Anda memiliki saldo sebanyak: Rp.%.0f\n", *saldoUser)
		} else if PilihanMenuTransaksi == 4 {
			var tambah float64
			fmt.Print("Masukkan jumlah saldo yang ingin ditambahkan: Rp.")
			fmt.Scan(&tambah)

			if tambah > 0 {
				*saldoUser = *saldoUser + tambah
				fmt.Printf("Saldo berhasil ditambahkan! jumlah saldo anda : Rp.%.0f\n", *saldoUser)
			} else {
				fmt.Println("Jumlah tidak valid!")
			}
		} else {
			fmt.Println("Tolong masukkan angka yang sesuai!")
		}
	}
}

// Bagian Nakulla

func isiDataKripto(data *arrKripto) {
	*data = [NMAX]kripto{
		{"Wrapped_Bitcoin", "WBTC", 1737999996},
		{"Bitcoin", "BTC", 1719700000},
		{"Yearn_Finance", "YFI", 94001331},
		{"Pax_Gold", "PAXG", 53939733},
		{"Tether_Gold", "XAUT", 52990080},
		{"Ethereum", "ETH", 42556000},
		{"Maker", "MKR", 29566504},
		{"BNB", "BNB", 10808112},
		{"Bitcoin_Cash", "BCH", 6650000},
		{"Aave", "AAVE", 3750267},
	}
}

// Bagian Nakulla

func menampilkanDataArray(data arrKripto, n int) {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("|        Silahkan Pilih Aset Kripto Yang Ingin Dibeli       |")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-10s | %-15s |\n", "No", "Nama", "Kode", "Harga (IDR)")
	fmt.Println("-------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-3d | %-20s | %-10s | Rp.%-12.0f |\n", i+1, data[i].nama, data[i].kode, data[i].harga)
	}
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()
}

// Bagian Nakulla

func beliAsetKripto(data arrKripto, asetUser *arrAset, riwayatUser *arrTransaksi, saldoUser *float64, jumlahAsetUser, jumlahTransaksiUser *int) {
	var kode string
	var idx, idn int
	var jumlah, nominal float64

	menampilkanDataArray(data, 10)

	fmt.Println("PETUNJUK !")
	fmt.Println("Jika ingin membatalkan cukup masukkan '-'")
	fmt.Print("Masukkan kode aset yang ingin dipilih: ")
	fmt.Scan(&kode)

	if kode == "-" {
		return
	}

	idx = seqsearchidx(data, 10, kode)
	if kode != "-" {
		for idx == -1 {
			fmt.Println()
			fmt.Print("Aset tidak ditemukan, tolong masukkan kode dengan benar!")
			fmt.Println()
			fmt.Println("PETUNJUK !")
			fmt.Println("Jika ingin membatalkan cukup masukkan '-'")
			fmt.Print("Masukkan kode aset yang ingin dipilih: ")
			fmt.Scan(&kode)

			if kode == "-" {
				return
			}

			idx = seqsearchidx(data, 10, kode)
		}
	}

	fmt.Print("Masukkan nominal rupiah yang ingin dibelanjakan: Rp.")
	fmt.Scan(&nominal)

	if *saldoUser < nominal {
		fmt.Println("Mohon maaf saldo tidak cukup! Silahkan cek saldo anda")
		return
	}

	jumlah = nominal / float64(data[idx].harga)

	idn = -1
	for i := 0; i < *jumlahAsetUser; i++ {
		if asetUser[i].kode == data[idx].kode {
			asetUser[i].jumlah += jumlah
			idn = i
		}
	}

	if idn == -1 {
		asetUser[*jumlahAsetUser].nama = data[idx].nama
		asetUser[*jumlahAsetUser].kode = data[idx].kode
		asetUser[*jumlahAsetUser].jumlah = jumlah
		asetUser[*jumlahAsetUser].harga = data[idx].harga
		asetUser[*jumlahAsetUser].nominal = nominal
		*jumlahAsetUser++
	}

	*saldoUser = *saldoUser - nominal

	riwayatUser[*jumlahTransaksiUser].nama = data[idx].nama
	riwayatUser[*jumlahTransaksiUser].kode = data[idx].kode
	riwayatUser[*jumlahTransaksiUser].harga = data[idx].harga
	riwayatUser[*jumlahTransaksiUser].jumlah = jumlah
	riwayatUser[*jumlahTransaksiUser].tipe = "Beli"
	riwayatUser[*jumlahTransaksiUser].total = nominal
	*jumlahTransaksiUser++

	fmt.Println()
	fmt.Printf("Transaksi berhasil! Anda berhasil membeli %f %s\n", jumlah, data[idx].nama)
	fmt.Printf("Jumlah saldo anda sekarang Rp.%.0f\n", *saldoUser)
}

// Bagian Nakulla

func jualAsetKripto(data arrKripto, asetUser *arrAset, riwayatUser *arrTransaksi, saldoUser *float64, jumlahAsetUser, jumlahTransaksiUser *int) {
	var nama string
	var idx int
	var jumlahJual, harga, income float64

	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("|        Silahkan Pilih Aset Kripto Yang Ingin DiJual       |")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-10s | %-15s |\n", "No", "Nama", "Kode", "Harga (IDR)")
	fmt.Println("-------------------------------------------------------------")
	for i := 0; i < *jumlahAsetUser; i++ {
		fmt.Printf("| %-3d | %-20s | %-10s | %-15f |\n", i+1, asetUser[i].nama, asetUser[i].kode, asetUser[i].jumlah)
	}
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()

	fmt.Println("PETUNJUK !")
	fmt.Println("Jika ingin membatalkan cukup masukkan '-'")
	fmt.Print("Masukkan nama aset yang ingin dipilih: ")
	fmt.Scan(&nama)

	if nama == "-" {
		return
	}

	idx = binsearchidx(*asetUser, *jumlahAsetUser, nama)
	if nama != "-" {
		for idx == -1 && nama != "-" {
			fmt.Println()
			fmt.Print("Anda tidak memiliki aset tersebut!")
			fmt.Println()
			fmt.Println("PETUNJUK !")
			fmt.Println("Jika ingin membatalkan cukup masukkan '-'")
			fmt.Print("Masukkan kode aset yang ingin dipilih: ")
			fmt.Scan(&nama)

			if nama == "-" {
				return
			}

			idx = binsearchidx(*asetUser, *jumlahAsetUser, nama)
		}
	}

	fmt.Print("Masukkan jumlah aset yang ingin dijual: ")
	fmt.Scan(&jumlahJual)

	const epsilon = 1e-6

	var selisih float64 = asetUser[idx].jumlah - jumlahJual

	if selisih < -epsilon {
		fmt.Println()
		fmt.Println("Mohon maaf aset yang anda miliki tidak cukup!")
		return
	}

	asetUser[idx].jumlah -= jumlahJual

	harga = 0
	for i := 0; i < 10; i++ {
		if data[i].kode == asetUser[idx].kode {
			harga = data[i].harga
		}
	}

	income = jumlahJual * harga
	*saldoUser += income

	riwayatUser[*jumlahTransaksiUser].nama = data[idx].nama
	riwayatUser[*jumlahTransaksiUser].kode = data[idx].kode
	riwayatUser[*jumlahTransaksiUser].harga = data[idx].harga
	riwayatUser[*jumlahTransaksiUser].jumlah = jumlahJual
	riwayatUser[*jumlahTransaksiUser].tipe = "Jual"
	riwayatUser[*jumlahTransaksiUser].total = income
	*jumlahTransaksiUser++

	if asetUser[idx].jumlah < epsilon {
		for i := idx; i < *jumlahAsetUser-1; i++ {
			asetUser[i] = asetUser[i+1]
		}
		*jumlahAsetUser--
	}

	fmt.Println()
	fmt.Printf("Transaksi berhasil! Anda berhasil menjual %f %s\n", jumlahJual, data[idx].nama)
	fmt.Printf("Rp.%.0f masuk ke saldo anda! jumlah saldo anda sekarang Rp.%.0f\n", income, *saldoUser)
}

// Bagian Rasya

func tampilkanRiwayat(riwayatUser *arrTransaksi, jumlah int) {
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Println("|                            Riwayat Transaksi Anda                             |")
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-6s | %-20s | %-6s | %-12s | %-15s |\n", "No", "Tipe", "Nama", "Kode", "Jumlah", "Total")
	fmt.Println("---------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("| %-3d | %-6s | %-20s | %-6s | %-12f | Rp.%-12.0f |\n", i+1, riwayatUser[i].tipe, riwayatUser[i].nama, riwayatUser[i].kode, riwayatUser[i].jumlah, riwayatUser[i].total)
	}
	fmt.Println("---------------------------------------------------------------------------------")
}

// Bagian Rasya

func mencariAset(data arrKripto, asetUser *arrAset, jumlahAsetUser *int) {
	var kode string
	var idn int

	menampilkanDataArray(data, 10)

	fmt.Println("PETUNJUK !")
	fmt.Println("Jika ingin membatalkan cukup masukkan '-'")
	fmt.Print("Masukkan nama aset yang ingin dipilih: ")
	fmt.Scan(&kode)

	if kode == "-" {
		return
	}

	idn = binsearchidx(*asetUser, *jumlahAsetUser, kode)

	if kode != "-" {
		for idn == -1 {
			fmt.Println()
			fmt.Print("Anda belum memiliki aset tersebut!")
			fmt.Println()
			fmt.Println("PETUNJUK !")
			fmt.Println("Jika ingin membatalkan cukup masukkan '-'")
			fmt.Print("Masukkan nama aset yang ingin dipilih: ")
			fmt.Scan(&kode)

			if kode == "-" {
				return
			}

			idn = binsearchidx(*asetUser, *jumlahAsetUser, kode)
		}
	}

	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("|                   Berikut Informasi Aset Yang Anda Miliki                        |")
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s | %-15s | %-26s |\n", "Nama", "Kode", "Jumlah", "Nominal")
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Printf("| %-20s | %-10s | %-15f | Rp.%-23.0f |\n", asetUser[idn].nama, asetUser[idn].kode, asetUser[idn].jumlah, asetUser[idn].nominal)
	fmt.Println("------------------------------------------------------------------------------------")
}

// Bagian Nakulla

func menampilkanPorto(asetUser *arrAset, jumlahAsetUser *int) {
	var pilihanMenuPorto int
	var totalNilai float64

	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Println("|                        Berikut Informasi Aset Yang Anda Miliki                                  |")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-6s | %-12s | %-17s | %-22s |\n", "No", "Nama", "Kode", "Jumlah", "Harga", "Nominal")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	for i := 0; i < *jumlahAsetUser; i++ {
		fmt.Printf("| %-3d | %-20s | %-6s | %-12f | Rp.%-14.0f | Rp.%-19.0f |\n", i+1, asetUser[i].nama, asetUser[i].kode, asetUser[i].jumlah, asetUser[i].harga, asetUser[i].nominal)
		totalNilai += asetUser[i].nominal
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Printf("%-30s Total nilai portofolio anda: Rp.%.0f\n", " ", totalNilai)
	fmt.Println("---------------------------------------------------------------------------------------------------")

	for pilihanMenuPorto != 3 {
		fmt.Println()
		fmt.Println("-----------------------------------------------------")
		fmt.Println("             Apa yang ingin anda lakukan?")
		fmt.Println("-----------------------------------------------------")
		fmt.Println("1. Terurut Menurut Jumlah Terkecil")
		fmt.Println("2. Terurut Menurut Jumlah Terbesar")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------------------------------")
		fmt.Print("Pilihan Anda (1/2/3)? ")
		fmt.Scan(&pilihanMenuPorto)

		if pilihanMenuPorto == 1 {
			selectionSortUp(asetUser, *jumlahAsetUser)
			fmt.Println()
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Println("|                        Berikut Informasi Aset Yang Anda Miliki                                  |")
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Printf("| %-3s | %-20s | %-6s | %-12s | %-13s | %-25s |\n", "No", "Nama", "Kode", "Jumlah", "Harga", "Nominal")
			fmt.Println("---------------------------------------------------------------------------------------------------")
			for i := 0; i < *jumlahAsetUser; i++ {
				fmt.Printf("| %-3d | %-20s | %-6s | %-12f | Rp.%-14.0f | Rp.%-19.0f |\n", i+1, asetUser[i].nama, asetUser[i].kode, asetUser[i].jumlah, asetUser[i].harga, asetUser[i].nominal)
				totalNilai += asetUser[i].nominal
			}
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Printf("%-30s Total nilai portofolio anda: Rp.%.0f\n", " ", totalNilai)
			fmt.Println("---------------------------------------------------------------------------------------------------")
		} else if pilihanMenuPorto == 2 {
			insertionSortDown(asetUser, *jumlahAsetUser)
			fmt.Println()
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Println("|                        Berikut Informasi Aset Yang Anda Miliki                                  |")
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Printf("| %-3s | %-20s | %-6s | %-12s | %-13s | %-25s |\n", "No", "Nama", "Kode", "Jumlah", "Harga", "Nominal")
			fmt.Println("---------------------------------------------------------------------------------------------------")
			for i := 0; i < *jumlahAsetUser; i++ {
				fmt.Printf("| %-3d | %-20s | %-6s | %-12f | Rp.%-14.0f | Rp.%-19.0f |\n", i+1, asetUser[i].nama, asetUser[i].kode, asetUser[i].jumlah, asetUser[i].harga, asetUser[i].nominal)
				totalNilai += asetUser[i].nominal
				totalNilai += asetUser[i].nominal
			}
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Printf("%-30s Total nilai portofolio anda: Rp.%.0f\n", " ", totalNilai)
			fmt.Println("---------------------------------------------------------------------------------------------------")
		} else {
			fmt.Println("Tolong masukkan angka yang sesuai!")
		}
	}
}

// Bagian Rasya

func seqsearchidx(data arrKripto, n int, x string) int {
	for i := 0; i < n; i++ {
		if data[i].kode == x {
			return i
		}
	}
	return -1
}

// Bagian Rasya

func binsearchidx(asetUser arrAset, n int, x string) int {
	selectionSortUpByNama(&asetUser, n)
	var left, right, mid int
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if asetUser[mid].nama == x {
			return mid
		} else if x < asetUser[mid].nama {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// Bagian Rasya

func selectionSortUpByNama(asetUser *arrAset, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if asetUser[j].nama < asetUser[minIdx].nama {
				minIdx = j
			}
		}
		asetUser[i], asetUser[minIdx] = asetUser[minIdx], asetUser[i]
	}
}

// Bagian Rasya

func selectionSortUp(asetUser *arrAset, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if asetUser[j].jumlah < asetUser[maxIdx].jumlah {
				maxIdx = j
			}
		}
		asetUser[i], asetUser[maxIdx] = asetUser[maxIdx], asetUser[i]
	}
}

// Bagian Rasya

func insertionSortDown(asetUser *arrAset, n int) {
	var pass, i int
	var temp aset
	for pass = 1; pass < n; pass++ {
		temp = asetUser[pass]
		i = pass
		for i > 0 && temp.jumlah > asetUser[i-1].jumlah {
			asetUser[i] = asetUser[i-1]
			i--
		}
		asetUser[i] = temp
	}
}

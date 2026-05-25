package main

import (
	"fmt"
	"strings"
)

type Film struct {
	Judul  string
	Genre  string
	Tahun  int
	Rating float64
}

func main() {

	var films []Film = []Film{
		{"Interstellar", "Sci-Fi", 2014, 9.0},
		{"Avengers", "Action", 2012, 8.5},
		{"Parasite", "Thriller", 2019, 8.8},
		{"Inception", "Sci-Fi", 2010, 8.7},
		{"Joker", "Drama", 2019, 8.4},
	}

	var pilihan int

	for {

		fmt.Println("\n===== CINE REVIEW =====")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Edit Film")
		fmt.Println("3. Hapus Film")
		fmt.Println("4. Cari Film")
		fmt.Println("5. Cari Judul Film")
		fmt.Println("6. Urutkan Rating Film")
		fmt.Println("7. Urutkan Tahun Film")
		fmt.Println("8. Tampilkan Semua Film")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			tambahFilm(&films)

		} else if pilihan == 2 {

			editFilm(&films)

		} else if pilihan == 3 {

			hapusFilm(&films)

		} else if pilihan == 4 {

			var keyword string

			fmt.Print("Masukkan judul / genre : ")
			fmt.Scan(&keyword)

			cariFilm(films, keyword)

		} else if pilihan == 5 {

			var judul string

			urutJudul(&films)

			fmt.Print("Masukkan judul film : ")
			fmt.Scan(&judul)

			cariJudul(films, judul)

		} else if pilihan == 6 {

			urutRating(&films)

			fmt.Println("\nFilm berhasil diurutkan berdasarkan rating")
			tampilFilm(films)

		} else if pilihan == 7 {

			urutTahun(&films)

			fmt.Println("\nFilm berhasil diurutkan berdasarkan tahun")
			tampilFilm(films)

		} else if pilihan == 8 {

			tampilFilm(films)

		} else if pilihan == 0 {

			fmt.Println("Program selesai")
			break

		} else {

			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func tampilFilm(films []Film) {

	var i int

	for i = 0; i < len(films); i++ {

		fmt.Println("----------------------")
		fmt.Println("No      :", i+1)
		fmt.Println("Judul   :", films[i].Judul)
		fmt.Println("Genre   :", films[i].Genre)
		fmt.Println("Tahun   :", films[i].Tahun)
		fmt.Println("Rating  :", films[i].Rating)
	}
}

// tambah data film (@jebb_24)
func tambahFilm(films *[]Film) {

	var filmBaru Film

	fmt.Print("Masukkan Judul  : ")
	fmt.Scan(&filmBaru.Judul)

	fmt.Print("Masukkan Genre  : ")
	fmt.Scan(&filmBaru.Genre)

	fmt.Print("Masukkan Tahun  : ")
	fmt.Scan(&filmBaru.Tahun)

	fmt.Print("Masukkan Rating : ")
	fmt.Scan(&filmBaru.Rating)

	*films = append(*films, filmBaru)

	fmt.Println("Film berhasil ditambahkan")
}

// edit data film (eel)
func editFilm(films *[]Film) {

	var nomor int

	tampilFilm(*films)

	fmt.Print("Pilih nomor film yang ingin diedit : ")
	fmt.Scan(&nomor)

	nomor--

	if nomor >= 0 && nomor < len(*films) {

		fmt.Print("Judul Baru  : ")
		fmt.Scan(&(*films)[nomor].Judul)

		fmt.Print("Genre Baru  : ")
		fmt.Scan(&(*films)[nomor].Genre)

		fmt.Print("Tahun Baru  : ")
		fmt.Scan(&(*films)[nomor].Tahun)

		fmt.Print("Rating Baru : ")
		fmt.Scan(&(*films)[nomor].Rating)

		fmt.Println("Data film berhasil diubah")

	} else {

		fmt.Println("Nomor film tidak valid")
	}
}

// hapus data film (@jebb_24)
func hapusFilm(films *[]Film) {

	var nomor int

	tampilFilm(*films)

	fmt.Print("Pilih nomor film yang ingin dihapus : ")
	fmt.Scan(&nomor)

	nomor--

	if nomor >= 0 && nomor < len(*films) {

		*films = append((*films)[:nomor], (*films)[nomor+1:]...)

		fmt.Println("Film berhasil dihapus")

	} else {

		fmt.Println("Nomor film tidak valid")
	}
}

func cariFilm(films []Film, keyword string) {

	var i int
	var found bool = false

	fmt.Println("\nHasil Pencarian :")

	for i = 0; i < len(films); i++ {

		if strings.EqualFold(films[i].Judul, keyword) ||
			strings.EqualFold(films[i].Genre, keyword) {

			fmt.Println("----------------------")
			fmt.Println("Judul  :", films[i].Judul)
			fmt.Println("Genre  :", films[i].Genre)
			fmt.Println("Tahun  :", films[i].Tahun)
			fmt.Println("Rating :", films[i].Rating)

			found = true
		}
	}

	if found == false {
		fmt.Println("Film tidak ditemukan")
	}
}

func cariJudul(films []Film, judul string) {

	var kiri int = 0
	var kanan int = len(films) - 1
	var found bool = false

	for kiri <= kanan {

		var tengah int = (kiri + kanan) / 2

		if strings.EqualFold(films[tengah].Judul, judul) {

			fmt.Println("\nFilm ditemukan :")
			fmt.Println("Judul  :", films[tengah].Judul)
			fmt.Println("Genre  :", films[tengah].Genre)
			fmt.Println("Tahun  :", films[tengah].Tahun)
			fmt.Println("Rating :", films[tengah].Rating)

			found = true
			break

		} else if strings.ToLower(judul) >
			strings.ToLower(films[tengah].Judul) {

			kiri = tengah + 1

		} else {

			kanan = tengah - 1
		}
	}

	if found == false {
		fmt.Println("Film tidak ditemukan")
	}
}

func urutRating(films *[]Film) {

	var i int
	var j int
	var max int
	var temp Film

	for i = 0; i < len(*films)-1; i++ {

		max = i

		for j = i + 1; j < len(*films); j++ {

			if (*films)[j].Rating > (*films)[max].Rating {
				max = j
			}
		}

		temp = (*films)[i]
		(*films)[i] = (*films)[max]
		(*films)[max] = temp
	}
}

func urutTahun(films *[]Film) {

	var i int
	var j int
	var key Film

	for i = 1; i < len(*films); i++ {

		key = (*films)[i]
		j = i - 1

		for j >= 0 && (*films)[j].Tahun > key.Tahun {

			(*films)[j+1] = (*films)[j]
			j--
		}

		(*films)[j+1] = key
	}
}

func urutJudul(films *[]Film) {

	var i int
	var j int
	var key Film

	for i = 1; i < len(*films); i++ {

		key = (*films)[i]
		j = i - 1

		for j >= 0 &&
			strings.ToLower((*films)[j].Judul) >
				strings.ToLower(key.Judul) {

			(*films)[j+1] = (*films)[j]
			j--
		}

		(*films)[j+1] = key
	}
}
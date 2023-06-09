package main

func main() {

	// // simple array
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// // pengisian elemen array yang melebihi alokasi awal
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"
	// names[4] = "ez" // baris kode ini menghasilkan error

	// // Inisialisasi nilai awal array
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// // Inisialisasi nilai array dengan gaya vertikal
	// var fruits [4]string

	// // cara horizontal
	// fruits = [4]string{"apple", "grape", "banana", "melon"}

	// // cara vertikal
	// fruits = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }

	// // Array Multidimensi
	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// // Perulangan elemen array menggunakan keyword for
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// // Perulangan elemen array menggunakan keyword for - range
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// // Penggunaan variabel underscore _ Dalam for - range
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// }

	// alokasi elemen array menggunakan keyword make
	// var fruits = make([]string, 2)
	// fruits[0] = "apple"
	// fruits[1] = "manggo"

	// fmt.Println(fruits) // [apple manggo]
}

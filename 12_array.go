/*	Array
	Array adalah tipe data yang berisikan kumpulan data dengan tipe data yang sama.
	Saat membuat Array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut.
	Daya tampung Array tidak bisa bertambah setelah Array dibuat.
	isi array (index) dimulai dari 0
*/
package main

import "fmt"

func main() {
	// deklarasi array
	var name [2]string

	//assigment array ke 0
	name[0] = "R"
	//assigment array ke 1
	name[1] = "Danang"

	fmt.Println(name[0])
	fmt.Println(name[1])

	// assigment on declaration
	var DeretBilangan = [4]int{
		1,
		3,
		5,
		9,
	}
	fmt.Println(DeretBilangan)

	// bila panjang array belum diketahui kita dapat menggunakan [...]
	var unknownLength = [...]int8{
		2,
		4,
		6,
		8,
	}
	fmt.Println(unknownLength)
}

/* VARIABLE

- Variable iala tempat untuk menyimpan data.
- Variable digunakan agar kita dapat mengakses data yang sama dimanapun kita mau.
- Untuk membuat variable, kita dapat menggunakan kata  kunci var, lalu diikuti dengan nama variable dan tipe datanya.
- dalam pendeklarasiannya tidak boleh dimulai dengan anggka, tidak boleh menggunakan spasi, dan tidak boleh menggunakan reserved word seperti var/for/func.
ex
var duapuluh string			(correct)
var dua_puluh string		(correct)
var 2puluh string 			(wrong)
var func string				(wrong)
var dua puluh string		(wrong)

dalam programming, biasanya terdapat standarisasi penulisan variable, seperti camel case, snake case, dan lain-lain.
duaPuluhSatu		(Camel Case)
dua_puluh_satu		(Snake Case)
DuaPuluhSatu 		(Pascal Case)
dua-puluh-satu 		(Kebab Case)
DUA_PULUH_SATU 		(Upper Case Snake Case)

pada golang jika satu kata biasa menggunakan huruf kecil saja, sedangkan lebih dari satu kata mengunakan pascal case.

*/

package main

import "fmt"

func main() {
	// Deklarasi variable menggunakan var, kemudian nama variable (name) dan tipe data (string)
	var name string
	// assigment atau mengisi variable dengan data
	name = "Ridho danang"
	fmt.Println(name)

	// variabel dapat dideklarasikan bersamaan dengan assigment tanpa harus memberikan tipe data
	var namadua = "danang"
	fmt.Println(namadua)

	// penggunaannya dapat disingkat menggunakan := , bila disingkat kita tidak perlu menulis var dan tipe datanya
	namaku := "rdanang"
	fmt.Println(namaku)

	// defaultnya menggunakan tipe data int seperti alias jadi bila sistem menggunakan 64 bit maka akan menjadi int 64
	var umur = 12
	fmt.Println(umur)

	// kita dapat merubah defaultnya dengan mengetikan tipedatanya
	var umurdua uint8 = 32
	fmt.Println(umurdua)

	// declare multi variable
	var (
		firstName string
		lastName  = "Danang"
	)
	firstName = "r"
	fmt.Println(firstName + lastName)
}

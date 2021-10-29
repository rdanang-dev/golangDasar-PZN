/* Type declaration

- Type Declaration ialah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada (seperti alias).
- Type Declaration biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, Dengan tujuan agar lebih mudah dimengerti.

*/
package main

import "fmt"

func main() {
	type noKTP string
	//noKTP menjadi sebuah tipe data baru yang tipe datanya string
	var noKTPGUE noKTP = "123123123123123123455"
	fmt.Println(noKTPGUE)
}

/* Const

- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah diberi nilainya
- Cara pembuatannya samam dengan variable namun kata kuncinya const.
- Saat pemmbuatan constant, kita wajib langsung menginisialisasikan datanya.

*/
package main

import "fmt"

func main() {
	const fname string = "r"
	const lname = "danang"
	const (
		ndepan    string = "r"
		nbelakang        = "danang"
	)
	fmt.Println(fname + lname)
}

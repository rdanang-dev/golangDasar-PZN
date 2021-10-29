/*	Comparatioin Operator

	Operator Perbandingan
	- Operasi perbandingan adalah operasi untuk memandingkan dua buah data.
	- Operasi perbandingan adalah operasi yang menghasilkan boolean benar (true) atau salah (false).

	Operator							Keterangan
	>									Lebih dari
	<									Kurang dari
	>=									Lebih dari Sama Dengan
	<=									Kurang dari Sama Dengan
	==									Sama dengan
	!=									Tidak Sama Dengan

*/
package main

import "fmt"

func main() {
	namadepan := "r"
	namatengah := "danang"
	isSame := namadepan == namatengah
	fmt.Println(isSame)

	nilai1 := 10
	nilai2 := 11
	isBigger := nilai2 > nilai1
	fmt.Println(isBigger)
}

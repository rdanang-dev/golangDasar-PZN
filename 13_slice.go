/*	SLICE
	- Tipe data Slice adalah potongan dari data Array.
	- Slice mirip dengan Array yang memedakan adalah ukuran Slice bisa berubah.
	- Slice dan Array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di Array

	Tipe Data Slice memiliki 3 data, yaitu Pointer, Length dan Capacity
	- Pointer adalah petunjuk data pertama di Array pada Slice.
	- Length adalah panjang dari Slice.
	- Capacity ialah kapasitas dari slice, dimana length tidak boleh lebih dari capacity.

	Membuat Slice
	Basic Syntax						Keterangan
	array[foo:bar]						Membuat Slice dari Array dimulai dari index ke-foo sampai index sebelum bar.
	array[foo:]							Membuat Slice dari Array dimulai dari index ke-foo sampai akhir array.
	array[:bar]							Membuat Slice dari Array dimulai dari index ke-0 sampai index sebbelum bar.
	array[:]							Membuat slice dari array dimulai dari index ke-0 sampai index akhi Array.
*/
package main

import "fmt"

func main() {
	var month = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	fmt.Println(month)

}

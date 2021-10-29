/* Data Types Conversion

Konversi tipe data
untuk merubah cukup menuliskan function dengan nama tipe data yang diinginkan,
kemudian masukan variable yang akan diubah, jika niliai yang ditampung tidak
cukup maka nilai akan diulang dari awal nilai tersebut

contoh
var nilai int8 = 200
var convertnilai int32 = int32(nilai)

*/
package main

import "fmt"

func main() {
	var nilai32 uint32 = 258
	var nilai64 uint64 = uint64(nilai32)
	/*
		nilai8 akan terisi 1 karena batas dari uint8 ialah 255,
		setelah 255 kembali ke 0, karena nilai awal unit8 ialah 0,
		kemudian diteruskan dengan sisanya menjadi 2, karena valuenya 258
	*/
	var nilai8 uint8 = uint8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
}

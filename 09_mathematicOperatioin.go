/*	Mathematical Operation

	Mathematical Operation
	Operator				Definisi
	*						Perkalian
	/						Pembagian
	%						Modulus (Sisa hasil bagi)
	-						Pengurangan
	+						Penjumlahan

	Augmented Assigments
	Operasi Matematik					Augmented Assigments
	a = a * 10							a *= 10
	a = a / 10							a /= 10
	a = a % 10							a %= 10
	a = a - 10							a -= 10
	a = a + 10							a += 10
	maksudnya ialah nilai dari variable a ialah nilai a sebelumnya, kemudian di lakukan pengolahan dengan
	menggunakan operasi matematiknya dengan 10

	Unary Operator
	Operator				Definisi
	++						a = a+1
	--						a = a-1
	-						Negatif
	+						Positif (Bawaan)
	!						Kebalikan


*/
package main

import "fmt"

func main() {
	//alur proses
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	//proses lebih efisien
	result := 10 + 10
	fmt.Println(result)

	//deklarasi nilai awal
	nilai := 10
	fmt.Println(nilai)

	nilai += 2
	fmt.Println(nilai)
}

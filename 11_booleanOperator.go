/* 	Boolean Operator

Operasi Boolean
Operator				Keterangan
&&						Dan
||						Atau
!						Kebalikan

Tabel Nilai Bool Operator && (dan)
Nilai1				Operator 			Niliai2				Hasil
true				&&					true				true
true				&&					false				false
false				&&					true				false
false				&&					false				false
Mudahnya, kedua buah nilai harus terpenuhi atau bernilai benar (true) agar hasil true, selain itu false,
jadi bila salah satu kondisi tidak terpenuhi maka hasilnya false

Tabel Nilai Bool Operator || (atau)
Nilai1				Operator 			Niliai2				Hasil
true				&&					true				true
true				&&					false				true
false				&&					true				true
false				&&					false				false
Mudahnya, bila salah satu kondisi terpenuhi maka akan menghasilkan true, selain itu false.


Tabel Nilai Bool Operator ! (negasi)
Nilai				Operator 			Hasil
true				!					false
false				!					true
hanya membalikan nilai

*/
package main

import "fmt"

func main() {
	uts := 80
	uas := 80

	var LulusUjian = uts >= 80 && uas >= 80
	fmt.Println(LulusUjian)
}

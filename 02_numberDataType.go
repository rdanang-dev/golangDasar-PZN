/* Number

ada dua jenis tipe data number, yaitu:
- integer
- floating point


penulisan tipe data pada golang case sensitve
- int8 (correct)
- Int8 (wrong)
- INT8 (wrong)

Iteger min and max range
- uint8		unsigned 8-bit integer	(0 to 255)
- uint16	unsigned 16-bit integer	(0 to 65535)
- uint32	unsigned 32-bit integer	(0 to 4294967295)
- uint64	unsigned 64-bit integer	(0 to 18446744073709551615)
- int8		signed 8-bit integer	(-128 to 127)
- int16		signed 16-bit integer	(-32768 to 32767)
- int32		signed 32-bit integer	(-2147483648 to 2147483647)
- int64		signed 64-bit integer	(-9223372036854775808 to 9223372036854775807)

Floating point min and max range
- float32     IEEE-754 32-bit floating-point numbers (1.18*10^38 to 3.4x10^38)
- float64     IEEE-754 64-bit floating-point numbers (2.23*10^308 to 1.80x10^308)
untuk complex jarang digunakan, biasa digunakan untuk program matematik atau statistik
- complex64   complex numbers with float32 real and imaginary parts
- complex128  complex numbers with float64 real and imaginary parts

Alias
Tipe Data					Alias
- byte						uint8
- rune						int32
- int 						mengikuti sistem operasi misal 32 bit maka int32
- uint 						-----------------------  misal 64 bit maka uint64

*/
package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima =", 3.5)
}

package main

import "fmt"

var book map[string]float64
var listaTelefonica map[string]uint32

func main() {
	//	[ 0 1 2 3 4 ]
	// "Maçã" -> 3
	// 3 -> 0.67
	// "Leite" -> 0
	// 0 -> 1.49
	// [ 1.49 0.79 2.49 0.67 1.49 ]
	// "Hmm, abacate?" -> "Abacate" -> 4 -> 1.49!
	book = make(map[string]float64)
	book["maçã"] = 0.67
	book["leite"] = 1.49
	book["abacate"] = 1.49
	fmt.Println(book)
	fmt.Println(book["maçã"])

	listaTelefonica = make(map[string]uint32)
	listaTelefonica["jenny"] = 8675309
	listaTelefonica["emergency"] = 911
	fmt.Println(listaTelefonica["jenny"])
}

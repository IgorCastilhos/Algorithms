package main

import "fmt"

// fatorialTailRec é uma função auxiliar que implementa a recursão de cauda.
// Ela recebe dois argumentos: n (o número para o qual o fatorial é calculado)
// e acc (o acumulador que carrega o resultado intermediário).
func fatorialTailRec(n int, acc int) int {
	if n == 0 {
		return acc
	}
	return fatorialTailRec(n-1, n*acc)
}

// Fatorial é a função pública que inicia a recursão de cauda.
// Ela chama fatorialTailRec com o acumulador inicializado em 1.
func Fatorial(n int) int {
	return fatorialTailRec(n, 1)
}

func main() {
	fmt.Println("Fatorial de 5:", Fatorial(5))
}

package main

import "fmt"

func sauda(name string) {
	fmt.Println(fmt.Sprintf("Olá, %s!", name))
	sauda2(name)
	fmt.Println("preparando para dizer tchau...")
	bye()
}

func sauda2(name string) {
	fmt.Println(fmt.Sprintf("Como vai, %s?", name))
}

func bye() {
	fmt.Printf("ok, tchau!")
}

func main() {
	sauda("Igor")
}

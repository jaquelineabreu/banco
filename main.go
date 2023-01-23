package main

import "fmt"

type ContaCorrente struct{
	titular string 
	numeroAgencia int64
	numeroConta int64
	saldo float64
}

func main(){
	fmt.Println(ContaCorrente{})	

}
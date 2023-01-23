package main

import "fmt"

type ContaCorrente struct{
	titular string 
	numeroAgencia int64
	numeroConta int64
	saldo float64
}

func main(){
	contaDaJaqueline := ContaCorrente{"Jaqueline", 589, 123456, 123.5}

	contaDaIvete := ContaCorrente{"Ivete",895,65321, 200}

	contaDaJorge := ContaCorrente{titular:"Jorge", saldo:200}


	fmt.Println(contaDaJaqueline)	
	fmt.Println(contaDaIvete)	
	fmt.Println(contaDaJorge)

}
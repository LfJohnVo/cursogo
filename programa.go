package main

import "fmt"

func main(){
	fmt.Println("**************MI PROGRAMA CON GO**************")

	edad := 18

	if edad >= 18 && edad <= 99 {
	fmt.Println("Es mayor de edad")
	}else {
		fmt.Println("Es menor de edad o demasiado mayor")
	}
}

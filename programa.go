package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("**************MI PROGRAMA CON GO**************")

	/*fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en go")

	edad, _ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad != 25 && edad != 99 {
	fmt.Println("Es mayor de edad")
	}else if edad == 25{
		fmt.Println("Tienes 25 años")
	}else if edad == 99{
		fmt.Println("Pronto moriras")
	}else{
		fmt.Println("eres menor de edad")
	}*/

	/*
	numero,_ := strconv.Atoi(os.Args[1])

	if numero % 2 == 0 {
		fmt.Println("El numero es par")
	}else{
		fmt.Println("el numero es impar")
	}
	 */

	//peliculas := []string{"Pelicula 1", "El club de la lucha" , "A todo gas", "Gran torino"}

	/*for i := 0; i < len(peliculas); i++ {
		if i % 2 == 0 {
			fmt.Println("La pelicula ->"+ peliculas[i] +" es par" ,i )
		}else{
			fmt.Println("La pelicula ->"+ peliculas[i] + " es impar" , i)
		}
		//fmt.Println(i , " es el numero actual");
	}*/

	/*for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}*/

	momento := time.Now()
	hoy := momento.Weekday();

	switch hoy {
	case 0:
		fmt.Println("domingo")
	case 1:
		fmt.Println("lunes")
	case 2:
		fmt.Println("martes")
	case 3:
		fmt.Println("miercoles")
	case 4:
		fmt.Println("jueves")
	case 5:
		fmt.Println("viernes")
	case 6:
		fmt.Println("sabado")
	default:
		fmt.Println("Es otro dia")
	}


}//Termina func

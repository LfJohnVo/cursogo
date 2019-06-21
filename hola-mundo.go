package main

import (
	"fmt"
	"time"
)

type Gorra struct {
	marca string
	color string
	precio float32
	plana bool
}


func main() {
	/*var gorra_negra =  Gorra {
		"nike",
		"Negro",
		25.50,
		false}
*/
	//fmt.Println(gorra_negra)

	/*
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Victor Robles"
	var apellidos string = " WEB"
	apellidos = " Lopez"
	pais := " España"
	fmt.Println(nombre + apellidos + pais)
	fmt.Println(suma)
	fmt.Println(resta)
*/
/*
	var numero1 float32 = 10;
	var numero2 float32 = 6;
	fmt.Println("calculadora 1: ");
	calculadora(numero1,numero2);
	var numero3 float32 = 44;
	var numero4 float32 = 7;
	fmt.Println("calculadora 2: ");
	calculadora(numero3,numero4);
*/
/*
	fmt.Println(gorras(45, "$"));
*/
	//pantalon("rojo", "largo" , "sin bolsillos" , "nike");

	var peliculas [3][2] string;
	peliculas[0][0] = "La verdad duele"
	peliculas[0][1] = "Mientras duermes"
	peliculas[1][0] = "Ciudadano ejemplar"
	peliculas[1][1] = "El señor de los anillos"
	peliculas[2][0] = "Gran torino"
	peliculas[2][1] = "A todo gas"



	//	peliculas := [3] string{"La verdad duele" , "Ciudadano ejemplar", "Batman"}



	fmt.Println(len(peliculas));

	time.Sleep(time.Second * 1);

}

func pantalon(caracteristicas ... string ){
	for _, caracteristica := range caracteristicas{
		fmt.Println(caracteristica)
	}
}

func gorras(pedido float32, moneda string)(string,string, float32) {
	precio := func() float32 {
		return pedido * 7
	}

	return "El precio de gorras pedidas es: ",moneda , precio()
}

func devolverTexto() (string , string) {
	dato1 := "victor"
	dato2 := "robles"

	return dato1,dato2
}

func holaMundo (){
	fmt.Println("Hola mundo");
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32;
	if (op == "+") {
			resultado = n1 + n2
	}
	if (op == "-") {
		resultado = n1 - n2
	}
	if (op == "*") {
		resultado = n1 * n2
	}
	if (op == "/") {
		resultado = n1 / n2
	}

	return resultado
}

func calculadora(numero1 float32, numero2 float32){
	//Suma
	fmt.Print("La suma es: ");
	fmt.Println(operacion(numero1, numero2, "+"));

	//Resta
	fmt.Print("La resta: ");
	fmt.Println(operacion(numero1, numero2, "-"));

	//multiplicacion
	fmt.Print("La multiplicacion es: ");
	fmt.Println(operacion(numero1, numero2, "*"));

	//Division
	fmt.Print("La division es: ");
	fmt.Println(operacion(numero1, numero2, "/"));
}
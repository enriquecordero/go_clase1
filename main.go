package main
import "fmt" // paquete para print 

const MiConstante = "Valor de mi constante"     
// variables y costantes 
func main() {
	fmt.Print("Hola Mundo desde Go")

	//declaracion por inferencia 
	var nombre string = "Enrique"
	fmt.Println(nombre)
	//declaracion rapida 

	nombre1 :="Juan"
	fmt.Println(nombre1)

	fmt.Printf("El valor de mi MiConstante es: %s \n",MiConstante)
}




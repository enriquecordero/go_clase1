package main

import (
	"flag"
	"fmt"
) // paquete para print


func main() {
		nombre := flag.String("nombre","Juan Del Pueblo","El nombre de la persona ")
		edad := flag.Int("edad",100 , "La edad de la persona")
		flag.Parse()
		fmt.Println("Tu nombre es:", *nombre)
		fmt.Println("Tu edad es:", *edad)
}




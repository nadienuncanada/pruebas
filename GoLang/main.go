package main

import "fmt"
type estruk struct{
	nom string
	cant int
}
var array [10] int
func main(){
	var a= "Texto auto definido?"
	fmt.Println("Tipo de dato ",a)
	for i:=1;i<10;i++{
		array[i]=i
	}
	var b=estruk {nom:"pepe",cant:15}
	fmt.Println(b)
	fmt.Println(array)
}

package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Print("This one is without linebreak!")
// }

func main(){
	Greeting()
	advancedGreeting("Arpan")
	advancedGreeting("V")
}

func Greeting(){
	fmt.Println("Hey There!")
}

func advancedGreeting(name string){
	fmt.Printf("Hello, %s\n", name)
}
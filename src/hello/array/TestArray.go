package main

import "fmt"

func main() {
	var langs [3] string
	langs[0]="GO"
	langs[1]="Javascript"
	langs[2]="Java"
	fmt.Println(langs)
	var langs2 [] string
	langs2 = append(langs2,"ciccio");

	fmt.Print(langs2)


	 langs3 := []string{"GO","JavaScript","Java"}

	fmt.Println(langs3)

	for i:= range langs{
		fmt.Println(langs[i])
	}


	for _,element:= range langs{
		fmt.Println(element)
	}
	}



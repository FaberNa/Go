package main

import ("fmt"
        )

func main(){
	var  i=0
	//var exitOnFive=true
	for {
		if i==5{
			break
		}
		fmt.Println(i);
		fmt.Println(main.testing())
		i++
	}

}

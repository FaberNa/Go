package main

import (
	"time"
	"fmt"
	"errors"
	"os"
)

func main(){

	hourOfDay:=time.Now().Hour()
	message,error :=getGreeting(hourOfDay)
	if(error!=nil){
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println(message)
}


func getGreeting(Hour int) (string,error){
	var message string
	if(Hour<6) {
		err := errors.New("too early")
		return message,err
	}
if Hour <12 {
	message= "Good Morning"
}else if Hour <18 {
	message= "Good afternoon"
}else{
	message= "Good evening"
}
return message,nil

}
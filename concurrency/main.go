package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	var wg sync.WaitGroup
	wg.Add(1) // Adds a cointer to the wait group
	// Anynomous immediately invoked function
	go func(){
		count("sheep")
		wg.Done() // Decrements the counter in wg.add( for the waitgroups) by 1
     }()
	wg.Wait() // this will wait untill the counter is 0 before exiting rather than exiiting all at once
}

func count( thing string){
	for i:=1;true; i++{
		fmt.Println(i,thing)
		time.Sleep(time.Millisecond * 500)
}

}
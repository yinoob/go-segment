package main

import "fmt"

func main() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	ch <- struct{}{}
	//The order is as follows:variable increment < channel send < channel receive < variable read

}
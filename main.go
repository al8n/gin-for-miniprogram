package main

import (
	"./apis"
	"log"
)

func main(){
	log.Println("Main log...")
	log.Fatal(apis.RunAPI("127.0.0.1:6340"))
}

package main

import (
	"fmt"
	"log"
	api "rest-api-demo/apis"
)

func main(){
	fmt.Println("welcome to system!")
	apiServer := api.NewAPIServer(":8080")
	err := apiServer.Run()
	if (err != nil){
		log.Fatalf("error when run server: %s", err)
	}
}
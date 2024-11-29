package main

import (
	"fmt"
	api "rest-api-demo/apis"
	"rest-api-demo/config"
	"rest-api-demo/db"
)

func main(){
	fmt.Println("welcome to system!")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
										config.Envs.Host,
										config.Envs.Port,
										config.Envs.User,
										config.Envs.Password,
										config.Envs.DBName,
									)
	//connect to database
	dbConn, err := db.NewPostgreSQL(connectionString)

	if err != nil {
		fmt.Println(err)
	}

	if (dbConn != nil){
		fmt.Println("connected to database")
	}
	apiServer := api.NewAPIServer(":8080", dbConn)
	errDb := apiServer.Run()
	if (errDb != nil){
		fmt.Printf("error when run server: %s", err)
	}


}
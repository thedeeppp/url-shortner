package main

import (
	database "url-shortner/db"
	router "url-shortner/routes"

	"github.com/joho/godotenv"
)

func init(){
		err := godotenv.Load(".env")
		if err != nil{
			panic(err)
		}
	database.ConnectDb()

}

func main() {
	router.ClientRoutes()
}
package main

import (
	db "aed/database"
	"aed/routers"
)
func main()  {
	defer db.SqlDB.Close()

	router := routers.InitRouter()

	router.Run(":8080")

}

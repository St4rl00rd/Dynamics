package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/st4rl00rd/dynamics/movements"
)

func main() {
	// dbConnection()
	httpPort := "3000"

	router := httprouter.New()
	router.GET("/", movements.Index)
	router.GET("/movements/:id", movements.Show)

	fmt.Println("\n\nFull speed ahead Mr. Boatswain, full speed ahead! \nFull speed ahead it is, Sergeant. \nCut the cable, drop the cable! \nAye! Sir, Aye Captain, captain!\nSystem's Up!!! And Listening on port: " + httpPort)

	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}

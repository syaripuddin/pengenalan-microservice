// package main

// import (
// 	"github.com/gorilla/mux"
// )

// func main(){
// router := mux.NewRouter()

// router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

// fmt.Handle("Menu service listen port : 8000")
// log.Panic(htpp.ListenAndService)
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"https://github.com/syaripuddin/pengenalan-microservice/menu-service/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service listen on port :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}

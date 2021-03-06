// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/wskurniawan/digitalent-microservice/menu-service/handler"
// )

// func main() {
// 	router := mux.NewRouter()

// 	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

// 	fmt.Println("Menu service listen on port :8000")
// 	log.Panic(http.ListenAndServe(":8000", router))
// }

//test

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syaripuddin/menu-service/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service listen on port :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}

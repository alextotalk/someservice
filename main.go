package main

import (
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc/someservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello

	sm := http.NewServeMux()
	sm.Handle("/", hh(l))

	/* x := "IT`s working"
	fmt.Printf("%v \n", x)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		})

	http.HandleFunc("/goodbay", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbay World")
	})
	*/
	err := http.ListenAndServe("127.0.0.1:9090", sm)
	if err != nil {
		log.Fatal("Something is bad", err)
	}
}

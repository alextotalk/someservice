package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {

		//func of http package
		http.Error(rw, "Ooops", http.StatusBadRequest)

		/* res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Oooops"))
		return */
	}

	//log.Printf("Data %s \n", d)
	fmt.Fprintf(rw, "Hellow %s", d)

}

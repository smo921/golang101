package golang101

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type foobar struct {
	Req    []string
	hidden string
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := strings.TrimLeft(r.URL.Path, "/json")
	resp := foobar{strings.Split(reqPath, "/"), "you can't see me"}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(b))
}

func txtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path[1:])
}

// RunServer is a more complex example of an http server
func RunServer() {
	http.HandleFunc("/json/", jsonHandler)
	http.HandleFunc("/", txtHandler)
	http.ListenAndServe(":8080", nil)
}

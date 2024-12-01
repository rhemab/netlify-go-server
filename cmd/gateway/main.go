package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/carlmjohnson/gateway"
)

type apiResponse struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := ""
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./static")))
	}
	http.HandleFunc("/api", apiRoute)

	log.Fatal(listener(portStr, nil))
}

func apiRoute(w http.ResponseWriter, r *http.Request) {
	apiRouteRes := &apiResponse{
		Url:    r.Host + "/api",
		Method: r.Method,
	}
	jsonRes, err := json.Marshal(apiRouteRes)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonRes)
}

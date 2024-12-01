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
	http.HandleFunc("/js", jsPage)

	log.Fatal(listener(portStr, nil))
}

func apiRoute(w http.ResponseWriter, r *http.Request) {
	apiRouteRes := &apiResponse{
		Url:    "/api",
		Method: r.Method,
	}
	jsonRes, err := json.Marshal(apiRouteRes)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonRes)
}

func jsPage(w http.ResponseWriter, r *http.Request) {
	htmlPage := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Web App on Netlify</title>
    <style>
  	body {background: grey;}
    </style>
</head>
<body>
    <h1>Go Web App on Netlify</h1>
    <h4>Super Speedy Website</h4>
    <h4>%s</h4>
</body>
</html>
`, r.Host+"/js")
	w.Write([]byte(htmlPage))
	// var htmlTemplate = template.Must(template.New("").Parse(htmlPage))
	// htmlTemplate.Execute(w, r.Host+"/js")
}

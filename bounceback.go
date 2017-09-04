package main

import (
	"os"
	"log"
	"net/http"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	_, set := os.LookupEnv("BOUNCEBACK_URLS")

	if ! set {
		log.Println("fn=main at=no-bounceback-urls")
	}

	http.ListenAndServe(":" + port, handler())
}


func handler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request){
		resp.Write([]byte("testing 1 2 3"))
	}
}

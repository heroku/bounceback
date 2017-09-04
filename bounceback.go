package main

import (
	"os"
	"log"
	"net/http"
	"strings"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}



	http.ListenAndServe(":" + port, handler())
}


func handler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request){
		for k, _ := range pgbouncerUrls() {
			resp.Write([]byte(k))
		}
	}
}


func pgbouncerUrls() map[string]string {
	urls := map[string]string{}
	envs := os.Environ()
	for _, env := range envs {
		kv := strings.SplitN(env, "=", 1)
		if strings.HasSuffix(kv[0], "_PGBOUNCER"){
			urls[kv[0]] = kv[1]
		}
	}
}

func bouncebackUrls() []string {

	urlz, set := os.LookupEnv("BOUNCEBACK_URLS")

	if ! set {
		log.Println("fn=main at=no-bounceback-urls")
	}

	return strings.Split(urlz,",")
}

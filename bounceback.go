package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	go notify()

	http.ListenAndServe(":"+port, handler())
}

func handler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		for k, _ := range pgbouncerUrls() {
			resp.Write([]byte(k))
		}
	}
}

func pgbouncerUrls() map[string]string {
	urls := map[string]string{}
	envs := os.Environ()
	for _, env := range envs {
		kv := strings.Split(env, "=")
		if strings.HasSuffix(kv[0], "_PGBOUNCER") {
			urls[kv[0]] = kv[1]
		}
	}
	return urls
}

func bouncebackUrls() []string {

	urlz, set := os.LookupEnv("BOUNCEBACK_URLS")

	if !set {
		log.Println("fn=main at=no-bounceback-urls")
		return []string{}
	}

	return strings.Split(urlz, ",")
}

func notify() {
	pgUrls := pgbouncerUrls()
	j, err := json.Marshal(pgUrls)
	if err != nil {
		log.Printf("func=notify at=json-err error=%q", err)
		return
	}

	for _, url := range bouncebackUrls() {
		req, err := http.NewRequest("POST", url, bytes.NewReader(j))
		if err != nil {
			log.Printf("func=notify at=req-err error=%q", err)
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("func=notify at=http-err error=%q", err)
			continue
		}

		if resp.StatusCode > 299 {
			log.Printf("func=notify at=http-code-err code=%d", resp.StatusCode)
		}
	}
}

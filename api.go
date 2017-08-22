package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ua, err := ParseUA(r.UserAgent())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())

			return
		}

		client := &http.Client{}
		req, err := http.NewRequest(ua.Method, ua.URL, bytes.NewReader([]byte(ua.Body)))
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, err.Error())

			return
		}

		if ua.Encoding != "" {
			req.Header.Set("content-type", ua.Encoding)
		}

		for k, v := range ua.Headers {
			req.Header.Set(k, v)
		}

		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, err.Error())

			return
		}

		for k, v := range resp.Header {
			w.Header().Set(k, v[0])
		}

		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		fmt.Fprintf(w, buf.String())
	})

	log.Panic(http.ListenAndServe(":666", nil))
}

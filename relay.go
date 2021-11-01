package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func relay(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	reqip := strings.Split(u.Host, ":")[0]
	wsurl := ip2wsurl(reqip)
	// fmt.Println(wsurl)

	var body io.Reader
	if r.Method == "POST" {
		body = r.Body
	}

	req, _ := http.NewRequest(r.Method, wsurl, body)
	a, err := client.Do(req)
	if err != nil {
		w.WriteHeader(502)
		fmt.Println(502, err)
		return
	}

	for k, _ := range a.Header {
		w.Header().Set(k, a.Header.Get(k))
	}

	// if a.StatusCode != 200 {
	// 	log.Println(a.StatusCode, reqip, ip2dc(reqip), wsurl)
	// }

	w.WriteHeader(a.StatusCode)

	io.Copy(w, a.Body)
}

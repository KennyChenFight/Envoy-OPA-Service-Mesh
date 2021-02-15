package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from service b")
}

func callingServiceA(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	req, err := http.NewRequest("GET", fmt.Sprintf("http://service-a:8080/service-a-hello?name=%s", r.Form.Get("name")), nil)
	if err != nil {
		fmt.Fprintln(w, "prepare calling service A fail:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(w, "start to calling service A fail:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(w, "marshal service A resp body fail:", err)
		return
	}
	fmt.Fprintln(w, "calling service A result: "+string(body))
}


func main() {
	http.HandleFunc("/service-b-hello", hello)
	http.HandleFunc("/calling-service-a", callingServiceA)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

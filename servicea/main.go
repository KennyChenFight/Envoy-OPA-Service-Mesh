package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func callingServiceB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	req, err := http.NewRequest("GET", fmt.Sprintf("http://service-b:8081/service-b-hello?name=%s", r.Form.Get("name")), nil)
	if err != nil {
		fmt.Fprintln(w, "prepare calling service B fail:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(w, "start to calling service B fail:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(w, "marshal service B resp body fail:", err)
		return
	}
	fmt.Fprintln(w, "calling service B result: "+string(body))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from service a")
}

func main() {
	http.HandleFunc("/service-a-hello", hello)
	http.HandleFunc("/calling-service-b", callingServiceB)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
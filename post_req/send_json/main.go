package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

    url := "http://localhost:8080/get_json"
    fmt.Println("URL:>", url)

    //var jsonStr = []byte(`{"name":"jinyoung.", "age":20}`)
    //println("jsonStr", jsonStr)
    jsonStr := &Person{Name: "jin", Age: 20}
    
    println("jsonStr", jsonStr)
    
    jsonBytes, _ := json.Marshal(jsonStr)
    println("jsonBytes", string(jsonBytes))
    
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}
package main
 
import (
    "io/ioutil"
    "net/http"
    "net/url"
)
 
func main() {
    // 간단한 http.PostForm 예제
    /*
    resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
    if err != nil {
        panic(err)
    }
 
    defer resp.Body.Close()
 
    // Response 체크.
    respBody, err := ioutil.ReadAll(resp.Body)
    if err == nil {
        str := string(respBody)
        println(str)
    }
    */
    var message []string 
    var nick []string 
    message = []string{"sending message"}
    nick = []string{"miumiu"}
    
    resp, err := http.PostForm("http://localhost:8080/form_post", url.Values{
    
		"message": message ,
		"nick":   nick,
    })
    if err != nil {
        panic(err)
    }
 
    defer resp.Body.Close()
 
    // Response 체크.
    respBody, err := ioutil.ReadAll(resp.Body)
    if err == nil {
        str := string(respBody)
        println(str)
    }    
}
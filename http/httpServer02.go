/*
Handling 404 Request in webserver
*/
package main

import (
  "log"
  "net/http"
)

func helloGopher(w http.ResponseWriter, r *http.Request){
  GreetGopher := "Hello Gopher"
  if r.URL.Path != "/" {
      http.NotFound(w,r)
      return
  }
  w.Write([]byte(GreetGopher))
}

func helloUser(w http.ResponseWriter, r *http.Request){
  GreetUser := "Hello User"
  w.Write([]byte(GreetUser))
}
func main(){
  // handleFunc
  http.HandleFunc("/", helloGopher)
  http.HandleFunc("/user", helloUser)

  // start server
  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal(err)
  }
}

/*
output
curl -is http://localhost:8000/
HTTP/1.1 200 OK
Date: Mon, 29 Apr 2019 21:26:53 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello Gopher

curl -is http://localhost:8000/user
HTTP/1.1 200 OK
Date: Mon, 29 Apr 2019 21:27:09 GMT
Content-Length: 10
Content-Type: text/plain; charset=utf-8

Hello User

curl -is http://localhost:8000/aaa
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 29 Apr 2019 21:27:48 GMT
Content-Length: 19

404 page not found

curl -is http://localhost:8000/user/aaa
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Mon, 29 Apr 2019 21:28:00 GMT
Content-Length: 19

404 page not found
*/

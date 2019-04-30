/*
Http webserver in GO
Golang offers creating of web server to load web contents to created web server.
net/http standard library package is helpful to create a web server.

Documentation: https://godoc.org/net/http
*/
package main

import (
//  "fmt"
  "net/http"
  "log"
)
// writing handler function
// func Handle(pattern string, handler Handler)
func helloGopher(w http.ResponseWriter, r *http.Request){
  message := "Hello Gopher"
  w.Write([]byte(message))
}

func helloUser(w http.ResponseWriter, r *http.Request){
  userName := "Hello Milind Dhoke"
  w.Write([]byte(userName))
}
func main(){
  // function handleFunc , doc: https://godoc.org/net/http#example-Handle
  // display hellow gopher on webserver
  // func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
  http.HandleFunc("/", helloGopher) // function which has contents to load

  http.HandleFunc("/user", helloUser)
  err := http.ListenAndServe(":8000", nil)
  if err != nil{
    log.Fatal(err)
  }

}
/*
output on http://localhost:8000/
Hello Gopher

mdhoke＠docker ➤ curl -I localhost:8000
HTTP/1.1 200 OK
Date: Mon, 29 Apr 2019 18:58:16 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8


Program starts executing from main function.
For creating webserver 2 methods have been used.
1. HandleFunc:
  This function is kind of providing a platform to handler function. It registers the function to map to URL address.
  First argument is the path to be loaded. From above example, when url is pointing to / , helloGopher function will be loaded, when path is /user then helloUser
  function will be loaded

＠docker ➤ curl -is http://localhost:8000/    // check path here
HTTP/1.1 200 OK
Date: Mon, 29 Apr 2019 19:07:37 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello Gopher

mdhoke＠docker ➤ curl -is http://localhost:8000/user   // check path here
HTTP/1.1 200 OK
Date: Mon, 29 Apr 2019 19:07:40 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

Hello Milind Dhokemdhoke

  This function take handle function as an argument for routing purpose. This can be called as routing mesh. What is to route and where can be done using HandleFunc
  method.

When routing is done, this should be published to webserver. To server the conter to the webserver, http.ListenAndServe method is used.
ListenAndServe:
  First argument to this function is port mapping. Port value is assigned where webserver should display the content. Second argument is handler function.Return value
  is error.

  func ListenAndServe(addr string, handler Handler) error
  ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
  Accepted connections are configured to enable TCP keep-alives. The handler is typically nil, in which case the DefaultServeMux is used.
  ListenAndServe always returns a non-nil error.

Handler Function:
  When route is mappeed to handler function. Handler funtion is responsible for what to request and what to response on webserver.
  Handler functinos has access to request and response. Response is written using Write method. Before writting response, collect all data for request and then
  write the response at the end.

Drawback:
  If you run below command:
  mdhoke＠docker ➤ curl -is http://localhost:8000/ssa // this path is not made available
  HTTP/1.1 200 OK
  Date: Mon, 29 Apr 2019 19:17:30 GMT
  Content-Length: 12
  Content-Type: text/plain; charset=utf-8

  Hello Gopher

  lets see the explaination of this Drawback in next program httpServer02.go
*/

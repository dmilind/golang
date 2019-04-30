/*
Custom http Request in Go
Custom http request is needed when we need more control on http request. We use custom http client for this.
http package's Get request sends default client request. If user wants to send some custom request like, setting header, connection time out etc.this can be done
using custom client.
In background ,

*/
package main

import(
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)
func main(){
  client := &http.Client{} // Client is the struct used for http method.
//  fmt.Println(client) // this will give nil data in the stuct.
  // lets make a GET method
  request, err := http.NewRequest("GET", "https://ifconfig.co", nil)  // request is declared but not used.
  if err != nil{
    log.Fatal(err)
  }
  // lets initialize the request
  resp, err := client.Do(request) // Do method is used to send the request and handle the response.
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(resp)
  defer resp.Body.Close()
  /*
  &{200 OK 200 HTTP/2.0 2 0 map[Cf-Ray:[4cf56899ce3428ac-SJC] Content-Length:[13] Content-Type:[text/plain; charset=utf-8] Date:[Tue, 30 Apr 2019 00:30:42 GMT] Expect-Ct:[max-a
  ge=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"] Server:[cloudflare] Set-Cookie:[__cfduid=d20612a15a0b241be74f4823741cff2d01556584242; expi
  res=Wed, 29-Apr-20 00:30:42 GMT; path=/; domain=.ifconfig.co; HttpOnly] Via:[1.1 vegur]] {0xc00026c140} 13 [] false false map[] 0xc000106000 0xc0000c0370}
  */

  // reading body from the response
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(string(body))
  /*
  &{200 OK 200 HTTP/2.0 2 0 map[Cf-Ray:[4cf5711d0b526ca0-SJC] Content-Length:[13] Content-Type:[text/plain; charset=utf-8] Date:[Tue, 30 Apr 2019 00:36:31 GMT] Expect-Ct:[max-a
  ge=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"] Server:[cloudflare] Set-Cookie:[__cfduid=d4e3057d5f8c4fc24420f1503f7c301d21556584590; expi
  res=Wed, 29-Apr-20 00:36:30 GMT; path=/; domain=.ifconfig.co; HttpOnly] Via:[1.1 vegur]] {0xc000218000} 13 [] false false map[] 0xc0000ec000 0xc0000b0370}
  74.217.73.96   ==> this is the response from webserver over http
  */
}

/*
POST Request in GO
Second request method is POST where user want to send some data to the server. This data could be json , file or image.
Lets see how post can be implemented in golang

*/
package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strings"
)

func main(){
  // suppose we want send a json data to an application running on webserver.
  // for testing we are using an http response and request application hosted on https://httpbin.org/
  // out goal is to send json data and getch the response if that one is sent or not
  // data to send is like {"User": "Gopher"} ,
  data := strings.NewReader(`{"User": "Gopher"}`) // reading json in data variable,
//  fmt.Println(data)
  // &{{"User": "Gopher"} 0 -1}
  // using post method to push the data
  resp, err := http.Post("https://httpbin.org/post", "application/json",data)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(resp) // this will return the header of POST method 200 OK means data has been pushed
  /*
  &{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[*] Connection:[keep-alive] Content-Type:[application/json] Date:[Mon, 29 Ap
  r 2019 23:51:14 GMT] Referrer-Policy:[no-referrer-when-downgrade] Server:[nginx] X-Content-Type-Options:[nosniff] X-Frame-Options:[DENY] X-Xss-Protection:[1; mode=block]] 0xc
  00020e140 -1 [] false true map[] 0xc000108000 0xc0000c22c0}
*/
 defer resp.Body.Close()
 // lets read the content of the body
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil{
   log.Fatal(err)
 }
// fmt.Println(body) // again body is in hexa format
 /*
 &{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[*] Connection:[keep-alive] Content-Type:[application/json] Date:[Mon, 29 Ap
r 2019 23:54:27 GMT] Referrer-Policy:[no-referrer-when-downgrade] Server:[nginx] X-Content-Type-Options:[nosniff] X-Frame-Options:[DENY] X-Xss-Protection:[1; mode=block]] 0xc
00000e1e0 -1 [] false true map[] 0xc000108000 0xc000306000}
[123 10 32 32 34 97 114 103 115 34 58 32 123 125 44 32 10 32 32 34 100 97 116 97 34 58 32 34 123 92 34 85 115 101 114 92 34 58 32 92 34 71 111 112 104 101 114 92 34 125 34 44
 32 10 32 32 34 102 105 108 101 115 34 58 32 123 125 44 32 10 32 32 34 102 111 114 109 34 58 32 123 125 44 32 10 32 32 34 104 101 97 100 101 114 115 34 58 32 123 10 32 32 32
32 34 65 99 99 101 112 116 45 69 110 99 111 100 105 110 103 34 58 32 34 103 122 105 112 34 44 32 10 32 32 32 32 34 67 111 110 116 101 110 116 45 76 101 110 103 116 104 34 58
32 34 49 56 34 44 32 10 32 32 32 32 34 67 111 110 116 101 110 116 45 84 121 112 101 34 58 32 34 97 112 112 108 105 99 97 116 105 111 110 47 106 115 111 110 34 44 32 10 32 32
32 32 34 72 111 115 116 34 58 32 34 104 116 116 112 98 105 110 46 111 114 103 34 44 32 10 32 32 32 32 34 85 115 101 114 45 65 103 101 110 116 34 58 32 34 71 111 45 104 116 11
6 112 45 99 108 105 101 110 116 47 49 46 49 34 10 32 32 125 44 32 10 32 32 34 106 115 111 110 34 58 32 123 10 32 32 32 32 34 85 115 101 114 34 58 32 34 71 111 112 104 101 114
 34 10 32 32 125 44 32 10 32 32 34 111 114 105 103 105 110 34 58 32 34 55 52 46 50 49 55 46 55 51 46 57 54 44 32 55 52 46 50 49 55 46 55 51 46 57 54 34 44 32 10 32 32 34 117
114 108 34 58 32 34 104 116 116 112 115 58 47 47 104 116 116 112 98 105 110 46 111 114 103 47 112 111 115 116 34 10 125 10]
*/
 // lets make it in human readable format
 fmt.Println(string(body))
 /*
 &{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[*] Connection:[keep-alive] Content-Type:[application/json] Date:[Mon, 29 Ap
r 2019 23:55:22 GMT] Referrer-Policy:[no-referrer-when-downgrade] Server:[nginx] X-Content-Type-Options:[nosniff] X-Frame-Options:[DENY] X-Xss-Protection:[1; mode=block]] 0xc
00016e080 -1 [] false true map[] 0xc0000fc000 0xc0004b00b0}
{
  "args": {},
  "data": "{\"User\": \"Gopher\"}",  // see here 
  "files": {},
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "18",
    "Content-Type": "application/json",
    "User-Agent": "Go-http-client/1.1"
   },
   "json": {
     "User": "Gopher"
   },
   "origin": "74.217.73.96, 74.217.73.96",
   "url": "https://httpbin.org/post"
 }
    "Host": "httpbin.org",

*/
}

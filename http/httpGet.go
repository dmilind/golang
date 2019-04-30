/*
GET method in GO
http (hypertext transfer protocol) is a network protocol which is used to collect and send data over to internet. This data might be document, json file, image etc.
Curl command is mostly used in general to perform https actions.
In go lang http actions can be performed using net/http standard library  package.
*/
package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "reflect"
)

func main(){
  // goal is to hit website www.orboo.in and get the response
  // if curl is used to do so below would be the response
  /*
  < HTTP/2 200
< link: <https://fonts.googleapis.com>; rel=preconnect; crossorigin,<https://fonts.gstatic.com>; rel=preconnect; crossorigin,<https://img1.wsimg.com>; rel=preconnect; crossorigin
< cache-control: max-age=30
< content-security-policy: frame-ancestors 'self'
< content-type: text/html;charset=utf-8
< vary: Accept-Encoding
< content-encoding: raw
< server: DPS/1.6.0
< x-siteid: 2000
< set-cookie: dps_site_id=2000; path=/; secure
< etag: 1dacc2ef33c3b64776502fef8ae3683f
< date: Mon, 29 Apr 2019 23:12:49 GMT
<
{ [10792 bytes data]

we want same response using golang
  */
// step 1. send request to http client , function used to so is func Get(url string) (resp *Response, err error)
// this function get to url and fetch te response in resp variable, else err. Retunrs values are 2
  req, err := http.Get("https://orboo.in/")
  // check if there is any error while fetching the response
  if err != nil {
    log.Fatal(err)
  }
  // if no error then lets print the response which is stored in resp var.
  fmt.Println(req)
  // output is in slice data structure
  /*
  &{200 OK 200 HTTP/2.0 2 0 map[Cache-Control:[max-age=30] Content-Security-Policy:[frame-ancestors 'self'] Content-Type:[text/html;charset=utf-8] Date:[Mon, 29 Apr 2019 23:18:
  37 GMT] Etag:[1dacc2ef33c3b64776502fef8ae3683f] Link:[<https://fonts.googleapis.com>; rel=preconnect; crossorigin,<https://fonts.gstatic.com>; rel=preconnect; crossorigin,<ht
  tps://img1.wsimg.com>; rel=preconnect; crossorigin] Server:[DPS/1.6.0] Set-Cookie:[dps_site_id=2000; path=/; secure] Vary:[Accept-Encoding] X-Siteid:[2000]] 0xc0000a50e0 -1 [
  ] false true map[] 0xc000108000 0xc0000c4370}
  */
  // at this point response body method is loaded in the RAM, to get efficient program alway close this method once Get method is done performing.
  // defer keyword is used to close the body once surrounded functions is executed succcessfully.
  defer req.Body.Close()

 // till this point , we actually got a header body. But what is the content of the website . To do so we need to read all content from the body
 body, err := ioutil.ReadAll(req.Body)
 if err != nil {
   log.Fatal(err)
 }
 // now all content is stored in body variable , lets try to printit
 //fmt.Println(body)
 // what has been printed ? something weird right ? lets check the type of variable body
 fmt.Println(reflect.TypeOf(body)) // slice of uint , means all data is encoded in hexa number. Now to find the actual data we need to convert this into string data type.
 fmt.Println(string(body))

 /* Output :
 &{200 OK 200 HTTP/2.0 2 0 map[Cache-Control:[max-age=30] Content-Security-Policy:[frame-ancestors 'self'] Content-Type:[text/html;charset=utf-8] Date:[Mon, 29 Apr 2019 23:36:
04 GMT] Etag:[1dacc2ef33c3b64776502fef8ae3683f] Link:[<https://fonts.googleapis.com>; rel=preconnect; crossorigin,<https://fonts.gstatic.com>; rel=preconnect; crossorigin,<ht
tps://img1.wsimg.com>; rel=preconnect; crossorigin] Server:[DPS/1.6.0] Set-Cookie:[dps_site_id=2000; path=/; secure] Vary:[Accept-Encoding] X-Siteid:[2000]] 0xc0000825d0 -1 [
] false true map[] 0xc000112000 0xc0003020b0}

[]uint8

<!DOCTYPE html><html lang="en-IN"><head><meta charSet="utf-8"/><meta http-equiv="X-UA-Compatible" content="IE=edge"/><meta name="viewport" content="width=device-width, initia
l-scale=1"/><title>Orboo</title><meta name="author" content="Orboo"/><meta name="generator" content="Starfield Technologies; Go Daddy Website Builder 8.0.0000"/><meta propert
y="og:url" content="https://orboo.in/"/>
<meta property="og:site_name" content="Orboo"/>
<meta property="og:title" content="We are coming soon !!!"/>
<meta property="og:type" content="website"/>
<meta property="og:locale" content="en_IN"/>
<meta name="twitter:card" content="summary"/>
<meta name="twitter:title" content="Orboo"/>

<meta name="twitter:description" content="We are coming soon !!!"/><script type="text/javascript" src="https://img1.wsimg.com/poly/v2/polyfill.min.js?unknown=polyfill&amp;fla
Git
GitHub
Initialize a new project directory with a Git repository
Create repository
....
...
...


 if you see we fetch the content of URL in html format.
 */
}

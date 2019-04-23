/*
Error Handling in Go
*/
package main

import (
  "fmt"
  "io/ioutil"
  "log"
)

func main(){
  // code to open a file anything.txt
  _, err := ioutil.ReadFile("/tmp/anything.txt") // this file is not available
  if err != nil {
    fmt.Println(err)    // printing error message
  }
  // output of this code
  // open /tmp/anything.txt: no such file or directory ==> this is error returned
  // lets throw custom error message
  _, err = ioutil.ReadFile("/tmp/blank.txt")
  if err != nil {
    log.Fatal("File not present in this host")
  }
  // output
  // 2019/04/23 00:33:56 File not present in this host
  // exit status 1
}

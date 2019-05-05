/*

*/
package main

import(
  "fmt"
  "time"
)

func main(){
  ticker := time.Tick(2 * time.Second)
  for t := range ticker{
    fmt.Println("Current time now", t)
  }
}
/*
output
Current time now 2019-04-30 00:11:05.663946 -0700 PDT m=+2.002372509
Current time now 2019-04-30 00:11:07.666069 -0700 PDT m=+4.004435031
Current time now 2019-04-30 00:11:09.666494 -0700 PDT m=+6.004799952
Current time now 2019-04-30 00:11:11.667226 -0700 PDT m=+8.005472277
Current time now 2019-04-30 00:11:13.667445 -0700 PDT m=+10.005631560
^Csignal: interrupt
*/

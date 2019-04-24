/*
Goroutines in Go
This is really cool and interesting feature provided by golang. I dont think much explaination is needed for the concept of Goroutines.
Before understanding the Goroutines , lets understand the concept of sequential and concurrent execution.

Sequential Execution:
  Some programming language are executed line by line. Once first instruction is executed then second instruction will be taken for execution. This is called
  as sequential execution. For example bash scripting comes under sequential Execution pattern. This approach is time consuming. Depending on the earlier
  instrunction's execution time , program takes time to get the final result.
  Now a days where everything is running fast, this is little ineffective. So what can be done to resolve this issue.

Concurrent Execution:
  To increase the performance of programming, concurrent execution concept has been placed. Here tasks or instructions can be executed parallely or concurrently.
  So irrespective of the execution time of any instruction, other instruction might get response , so there is no halting time under concurrent execution.

Example:
  Consider an example: We have 3 URLs, we want to calculate the time frame in which response is accepted by the program. if we use sequential execution then
  logic would be
  -> hit first URL and get the response  => this will take 1 sec to get response
  -> hit second URL and get the response => this instruction will be loaded after 1 sec => it takes 3 sec to get response
  -> hit third URL and get the response  => this instruction will be loaded after 3 sec => it take 2 sec to get response
  -> print a message -> total time to complete the program execution =>  6 ~ 7 sec
  So we have seen a halted execution with delay in execution.

  Same scenario in concurrent program
  -> hit first URL and get the response -> started execution at the same time
  -> hit second URL and get the response -> started execution at the same time
  -> hit third URL and get the response -> started execution at the same time -
  -> print a message -> started execution at the same time

  This is how concurrency is helpful to mitigate the time factory

Goroutines:
  concurrency can be implemented in golang using Goroutines. this implementation can be done just by using go keyword in front of instruction where we want concurrency.
  Fairly simple
*/
package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
)
func GetResponseTime(link string) {
  currentTime := time.Now()
  resp, err := http.Get(link)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  elapsedTime := time.Since(currentTime).Seconds()
  fmt.Printf("Time take by %s is %v\n", link, elapsedTime)
}
func main(){
  // get URL's
  // create a slice of url
  url := make([]string, 3 )
  url[0] = "https://www.amazon.in"
  url[1] = "https://www.amazon.de"
  url[2] = "https://www.amazon.com"

  for _, u := range url {
// without using goroutine
//    GetResponseTime(u)
// using goroutine
     go GetResponseTime(u)
  }
  time.Sleep(time.Second * 5) // this is a
}
/*
output without using goroutine
Time take by https://www.amazon.in is 0.430080191  ==> this has taken much time
Time take by https://www.amazon.de is 0.237607168  ==> less time
Time take by https://www.amazon.com is 0.1443611  ==> lesser time


// now lets try the same program using goroutine : by simply adding go keyword in front of function call GetResponseTime

Time take by https://www.amazon.com is 0.321830278  ==> check the order of the respose. 
Time take by https://www.amazon.de is 0.414331879
Time take by https://www.amazon.in is 0.422152276

*/

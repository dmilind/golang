/*
Reading a file in a golang
*/
package main

import (
  "fmt"
  "log"
  "io/ioutil"
)

func Reader(filePath string) string {
  // filePath is provided by the calling code.
  // using ReadFile method from ioutil package  func ReadFile(filename string) ([]byte, error)
  content, err := ioutil.ReadFile(filePath)
  if err != nil{
    log.Fatal(err)
  }
  // at this point content container a slice of byte , lets print it out to check
  fmt.Println(content)
  /*
  output:
  [84 104 105 115 32 105 115 32 97 32 115 97 109 112 108 101 32 102 105 108 101 46 32 84 104 105 115 32 102 105 108 101 32 119 105 108 108 32 98 101 32 114 101 97 100 32 98 121
  32 114 101 97 100 70 105 108 101 32 112 114 111 103 114 97 109 46 10 73 110 32 112 114 111 103 114 97 109 32 119 101 32 119 105 108 108 32 116 114 121 32 116 111 32 111 112
  101 110 32 116 104 105 115 32 102 105 108 101 44 32 97 108 108 32 99 111 110 116 101 110 116 115 32 111 102 32 116 104 105 115 32 102 105 108 101 32 119 105 108 108 32 98 101
  32 115 116 111 114 101 100 32 105 110 32 97 32 115 105 110 103 108 101 32 118 97 114 105 97 98 108 101 46 10 67 111 110 116 101 110 116 32 119 105 108 108 32 98 101 32 115 1
  16 111 114 101 100 32 105 110 32 116 104 101 32 102 111 114 109 32 111 102 32 115 108 105 99 101 32 111 102 32 98 121 116 101 46 32 84 111 32 109 97 107 101 32 105 116 32 104
  117 109 97 109 32 114 101 97 100 97 98 108 101 44 32 119 101 32 115 104 111 117 108 100 32 99 111 110 118 101 114 116 32 105 116 32 105 110 116 111 32 115 116 114 105 110 10
  3 46 10]
  */
  fileContent := string(content)
  return fileContent
}

func main(){
  // call Reader function by passing filepath
  fromFile := Reader("sampleFile")
  fmt.Println("File content is as below")
  fmt.Println(fromFile)
}
/*
Final output
File content is as below
This is a sample file. This file will be read by readFile program.
In program we will try to open this file, all contents of this file will be stored in a single variable.
Content will be stored in the form of slice of byte. To make it humam readable, we should convert it into string.
*/

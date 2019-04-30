/*
Write a file in go
*/

package main

import (
  "fmt"
  "log"
  "io/ioutil"
)

func Writer(filePath string, content string) string {
  // call WriteFile method from ioutil package ==> func WriteFile(filename string, data []byte, perm os.FileMode) error
  // since data should be in the type of []byte , if we pass data directly then compiler will error out bu we can show case this pattern.
  // err := ioutil.WriteFile(filePath, content, 0644)
  // if err != nil{
  //   log.Fatal(err)
  // }
  // output of above block of code
  // ./writeFile.go:16:26: cannot use content (type string) as type []byte in argument to ioutil.WriteFile
  // lets convert this string data into slice of byte. for that we can use make function
  byteContent := []byte(content)
//  fmt.Println(byteContent)
/*
output:
84 104 105 115 32 105 115 32 97 32 102 105 108 101 32 119 104 105 99 104 32 105 115 32 119 114 105 116 116 101 110 32 117 115 105 110 103 32 97 32 119 114 105 116 101 114 32
 102 117 110 99 116 105 111 110 32 105 110 32 103 111 108 97 110 103 10 84 104 105 115 32 100 97 116 97 32 105 115 32 98 101 105 110 103 32 112 97 115 115 101 100 32 116 111
32 116 104 101 32 102 117 110 99 116 105 111 110 32 119 104 101 114 101 32 115 108 105 99 101 32 111 102 32 98 121 116 101 32 119 105 108 108 10 98 101 32 103 101 110 101 114
 97 116 101 100 32 97 110 100 32 116 104 97 116 32 115 108 105 99 101 32 105 115 32 98 101 105 110 103 32 112 97 115 115 101 100 32 116 111 32 98 101 32 119 114 105 116 116 1
01 110 46]
*/
  // now we can use this slice of byte.
  err := ioutil.WriteFile(filePath,byteContent,0644)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println("File has been created in current location. Name of the file is",filePath)
  return "success"
}

func main(){
  // Make data available to be passed for writing a file
  data := "This is a file which is written using a writer function in golang\nThis data is being passed to the function where slice of byte will\nbe generated and that slice is being passed to be written.\n"

  // call function to pass this data
  status := Writer("writtenFile", data)
  if status != "success" {
    log.Fatal("function failed to create file")
    }
}
/*
Final Output:
File has been created in current location. Name of the file is writtenFile
*/

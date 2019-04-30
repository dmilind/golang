/*
List the content of the directory
if files is present then display the name of the file and the permission for that file
*/
package main

import (
  "fmt"
  "log"
  "io/ioutil"
)

func listDir(dirPath string) {
  // method used to list directory ==> func ReadDir(dirname string) ([]os.FileInfo, error)
  // return value would be a slice of filesinfo and error
  files, err := ioutil.ReadDir(dirPath)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(files)
  // output to above command
  // [0xc000070820 0xc0000a00d0 0xc000070750 0xc0000a0000 0xc000070680] ==> memory addess has been passed , we need to use pointer to get content from this memmory
  // since this is a slice so lets use for loop
  for _, file := range files {
    fmt.Println(file.Name(), file.Mode())
  }
}

func main(){
  // pass filepath to the function listDir
  listDir(".")
}
/*
Output:
listDir.go -rw-r--r--
readFile.go -rw-r--r--
sampleFile -rw-r--r--
writeFile.go -rw-r--r--
writtenFile -rw-r--r--
*/

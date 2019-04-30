/*
Copy one file to other file in go
Here ioutil package is not helpful so, we will use os package to perform this action
*/
package main

import(
  "os"
  "log"
  "io"
)

func CopyFile(){
  //step 1: open a parent file.
  sourceFile, err := os.Open("./sampleFile")
  if err != nil{
    log.Fatal(err)
  }
//  fmt.Println(sourceFile) // output ==> &{0xc00007c060}
  defer sourceFile.Close()

  targetFile, err := os.OpenFile("./sampleFile_copy", os.O_RDWR|os.O_CREATE, 0666) // func OpenFile(name string, flag int, perm FileMode) (*File, error)
  if err != nil{
    log.Fatal(err)
  }
  defer targetFile.Close()

  // copy sourceFile to targetFile
  _, err = io.Copy(sourceFile, targetFile)
  if err != nil{
    log.Fatal(err)
  }
}


func main(){
  // call CopyFile function
  CopyFile()
}

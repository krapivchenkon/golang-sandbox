package main

import (
    "fmt"
    "os"
)


func main(){
    file := createFile("test-file-create.txt")
    defer closeFile(file)
    writeToFile(file,"test-string")
}


func createFile(path string) *os.File{  
    fmt.Printf("createing a file %v\n", path)

    file,err := os.Create(path)
    if err != nil{
        panic(err)
    }

    return file
}


func writeToFile(file *os.File, str string) {
    fmt.Println("writing to file")
    fmt.Fprintln(file,str)
}


func closeFile(file *os.File) {
    fmt.Println("closing file")
    file.Close()
}
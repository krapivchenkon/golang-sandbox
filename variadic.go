package main

import (
    "fmt"
)


func main() {
    fmt.Println("Hello, world")

    printList("qweq","dhfw","rgqerg")


    list:=[]string{"bt2t24tb","42b2trdtu","aeyh45bq45g"}


    printList(list...)

}


func printList(d ...string){
    for _,name:=range d{
        fmt.Println(name)
    }
}

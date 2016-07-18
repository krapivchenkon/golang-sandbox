package main

func main() {
	var x string //defaults to "" (zero value)

	if x == "" {
		x = "default"
	}
}

// Strings cannot be nil
// FAILS

// func main() {
//     var x string = nil //error

//     if x == nil { //error
//         x = "default"
//     }
// }

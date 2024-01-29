package main
import ("fmt")

// parameter consist of parameter and type
func myMessage(fname string) {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}
package main
import ("fmt")

// parameter consist of parameter and type
func myMessage(fname string, secondvariable int) {
  fmt.Println("I just got executed, " + fname + secondvariable)
}

func main() {
  myMessage("john", 32) // call the function
}
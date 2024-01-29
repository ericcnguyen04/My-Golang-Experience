package main
import ("fmt")

// parameter consist of parameter and type
func myMessage(fname string, secondvariable int) {
  fmt.Println("I just got executed,", fname, secondvariable)
	// next time when console.logging; dont put plus +
}

func main() {
  myMessage("john", 32) // call the function
}
// int - whole numbers
// float32 - decimals
// string
// bool

// // two ways to declare var

// // 1.
// var variablename type = value

// // 2.
// secondvariable := value

package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}


// const
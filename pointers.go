package main
// Pointers - To get the memory address of a variable ("&"var)
// Function to change the value of a variable using its pointer.

import ("fmt")

func main() {
  age := 32
  // Using another function to increment using pointers
  inc_age(&age)
  fmt.Println("Incremented age is ", age)
  // You can view the memory address by using ""&var"
  fmt.Println("memory address for age = ", &age)
  // You can also do the below -
  name := "Manu"
  // To create a pointer variable for name -
  name_ptr := &name
  fmt.Println("Memory address for name = ", name_ptr)
  // To get the value of variable from the memory address, use "*"
  fmt.Println("Value of Name is ", *name_ptr)
}
// To modify the original variables in its address using "*".
func inc_age(x *int) {
  *x++
}
// This returns 32 + 1 = 33

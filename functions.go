package main

import (
  "fmt"
  "strings"
)

func main() {
  // For a simple function
  a := 5
  fmt.Println(square(a))

  // For multiple return values
  fmt.Println(sq_cube(a))

  // For String operations
  name := "manu manjunath"
  loc := "bangalore india"
  fmt.Println(string_oper(name, loc))

  // For varadic functions
  myScore := best_score(98, 90, 88, 78, 82, 92)
  fmt.Println(myScore)
}

// Simple function
func square(x int) int {
  // specify the type of parameter and return values
  return x * x
}

// To return multiple values -
func sq_cube(x int) (int, int) {
  sq := x * x
  cube := x * x * x
  return sq, cube
  }

	// To ignore certain return values
	_, qb := sqCube(a)
	fmt.Println(qb)

// String operations
func string_oper(name, loc string) (newName, newLoc string) {
  // Convert name to title case
  newName = strings.Title(name)
  // Convert location to upper case
  newLoc = strings.ToUpper(loc)
  return newName, newLoc
}

// Variadic function - when you don't know the number of parameters
func best_score(marks ... int) int {
  // use "..." in case of unknown number of parameters
  best := marks[0]
  for _, i := range marks {
    if i > best {
      best = i
    }
  }
  return best
}

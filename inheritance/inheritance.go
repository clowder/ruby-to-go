package main

import "fmt"

type Bird struct {}

func (bird Bird) speak() string {
  return "Squark"
}

func (bird Bird) name() string {
  return "Generic Bird Type"
}

type Parrot struct {
  Bird
}

func (bird Parrot) name() string {
  return "Parrot"
}

var bird Parrot

func main() {
  fmt.Printf("The %v says %v.\n", bird.name(), bird.speak())
}

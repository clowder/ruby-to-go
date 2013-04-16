package main

import "fmt"

type Bird struct {}

func (bird Bird) name() string {
  return "Generic bird type"
}

type TalkingBird struct {}

func (bird TalkingBird) speak() string {
  return "Hello polly"
}

type Parrot struct {
  Bird
  TalkingBird
}

func (bird Parrot) name() string {
  return "Parrot"
}

var parrot Parrot

func main() {
  fmt.Printf("The %v says %v.\n", parrot.name(), parrot.speak())
}

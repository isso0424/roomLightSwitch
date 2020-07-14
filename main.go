package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
  commandPtr, err := getCommand()
  if err != nil {
    log.Fatalln("Error: ", err)
  }

  switch (*commandPtr) {
  case "on":
    fmt.Println("on")
  case "off":
    fmt.Println("off")
  }
}

func getCommand() (*string, error) {
  flag.Parse()
  flags := flag.Args()

  if len(flags) == 0 {
    return nil, &getCommandError{}
  }

  command := flags[0]

  if command != "on" && command != "off" {
    return nil, &getCommandError{}
  }

  return &command, nil
}

type getCommandError struct {
}

func(ptr *getCommandError) Error() string {
  return "Invalid command"
}

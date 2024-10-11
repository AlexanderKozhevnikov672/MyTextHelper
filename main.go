package main

import (
	"connector"
	"os"
)

func main() {
	connector.Run(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
)

func main() {
	birthdays := readFile("reminder.conf")

	fmt.Printf("%#v\n", birthdays[len(birthdays)-1])
}

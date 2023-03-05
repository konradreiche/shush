package a

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello world") // want `extraneous fmt.Println statement`

	msg := "Hello world"
	msg = strings.ToUpper(strings.TrimSpace(msg))
	fmt.Println(msg) // want `extraneous fmt.Println statement`

	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteString(" ")
	buf.WriteString("world")

	fmt.Println(buf.String()) // want `extraneous fmt.Println statement`

	printHelloWorldTime()
}

func printHelloWorldTime() {
	timestamp := "2022-03-05T08:09:12-08:00"
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		fmt.Println(err) // want `extraneous fmt.Println statement`
		return
	}
	fmt.Println(t.In(time.Local).Add(1 * time.Hour)) // want `extraneous fmt.Println statement`
}

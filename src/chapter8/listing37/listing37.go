/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-19 15:01
 */
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello "))

	_, _ = fmt.Fprintf(&b, "World!")

	_, _ = b.WriteTo(os.Stdout)
}

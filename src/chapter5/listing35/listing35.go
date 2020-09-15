/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 22:12
 */
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("hello"))

	_, _ = fmt.Fprintf(&b, "world!")

	_, _ = io.Copy(os.Stdout, &b)
}

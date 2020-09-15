/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 22:24
 */
package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	//duration(42).pretty()
	s := duration(42)
	s.pretty()
}

/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 23:16
 */
package main

import (
	"chapter5/listing74/entities"
	"fmt"
)

func main() {
	a := entities.Admin{Rights: 10}

	a.Name = "Bill"
	a.Email = "bill@email.com"
	fmt.Printf("User: %v\n", a)
}

/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 23:11
 */
package main

import (
	"chapter5/listing71/entities"
	"fmt"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		Email: "bill@email.com",
	}
	fmt.Printf("User: %v\n", u)
}

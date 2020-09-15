/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 17:43
 */
package main

import (
	"database/sql"

	_ "chapter3/dbdriver/postgres"
)

// main is the entry point for the application.
func main() {
	sql.Open("postgres", "mydb")
}

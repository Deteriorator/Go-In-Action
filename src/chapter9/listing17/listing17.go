/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-19 16:23
 */
package main

import (
	"chapter9/listing17/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}

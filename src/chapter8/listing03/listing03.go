/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-19 13:50
 */
package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")

	log.Fatalln("fatal message")

	log.Panicln("panic message")
}

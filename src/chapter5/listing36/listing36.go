/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 22:16
 */
package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.email, u.email)
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotifition(&u)
}

func sendNotifition(n notifier) {
	n.notify()
}

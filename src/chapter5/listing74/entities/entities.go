/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 23:16
 */
package entities

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Rights int
}

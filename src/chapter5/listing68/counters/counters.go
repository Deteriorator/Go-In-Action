/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 22:58
 */
package counters

type alertCounter int

//type AlertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}

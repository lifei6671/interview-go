package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	letter,number := make(chan bool),make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		i := 0
		for{
			select {
			case <-letter:
				if i >= strings.Count(str,"")-1 {
					wait.Done()
					return
				}

				fmt.Print(str[i:i+1])
				i++
				if i >= strings.Count(str,"") {
					i = 0
				}
				fmt.Print(str[i:i+1])
				i++
				number <- true
				break
			default:
				break
			}

		}
	}(&wait)
	number<-true
	wait.Wait()
}



package Debug

import "fmt"

func Error(msg string) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println(e)
			}
		}()
		panic(msg)
	}()
}

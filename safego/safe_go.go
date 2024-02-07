package safego

import "fmt"

func Go(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go f()
}

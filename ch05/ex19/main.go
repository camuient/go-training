package main

func f(i int) (result int) {
	defer func() {
		result = recover().(int)
	}()
	panic(i)
}

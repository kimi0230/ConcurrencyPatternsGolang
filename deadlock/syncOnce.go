package main

import "sync"

func main() {
	// <1>調用的Do 直到<2>調用Do並退出才會繼續
	var onceA, onceB sync.Once
	var initB func()
	initA := func() {
		onceB.Do(initB)
	}
	initB = func() { // <1>
		onceA.Do(initA)
	}
	onceA.Do(initA) // <2>
}

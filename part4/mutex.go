package main

import (
	"fmt"
	"runtime"
	"sync"
)
var counter = 0

func main(){
	var wg sync.WaitGroup
	var m = sync.RWMutex{}
	
	
	runtime.GOMAXPROCS(100)
	for i:=0;i<10; i++{
		wg.Add(2)
		m.RLock()
		go sayHello(&m,&wg)
		m.Lock()
		go increment(&m,&wg)
	}
	
	wg.Wait()
}

func sayHello(m *sync.RWMutex,wg *sync.WaitGroup){
	
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment(m *sync.RWMutex,wg *sync.WaitGroup){
	
	counter++
	m.Unlock()
	wg.Done()
}

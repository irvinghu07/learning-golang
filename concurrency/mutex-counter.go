package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct{
	m sync.Mutex
	locker map[string]int
}


func (s *safeCounter) increment(key string){
	s.m.Lock()
	s.locker[key]++
	s.m.Unlock()
}

func (s *safeCounter) fetchValue(key string) int{
	defer s.m.Unlock()
	s.m.Lock()
	return s.locker[key]
}

func main(){
	token := "Tony Stark"
	s := &safeCounter{locker : make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go s.increment(token)
	}

	time.Sleep(time.Second)
	fmt.Println(s.fetchValue(token))
}
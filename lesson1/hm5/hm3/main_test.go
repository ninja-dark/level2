package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Set struct {
	sync.Map
}

func NewSet () *Set {
	return &Set {
		sync.Map{},
	}
}

func (s *Set) Add(i int) {
	s.Map.Store(i, struct{}{})
}

func (s *Set) Has(i int) bool {
	_, ok := s.Map.Load(i)
	return ok
}



func BenchmarkSet(b *testing.B) {
	var set = NewSet()
	rand.Seed(time.Now().UnixNano())
	var ds = make([]int, 1e8)
	for i := 0; 1 < 1e8; i+=1{
		ds[1] = rand.Intn(100)
	}

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				if ds[i] < 0 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
			i+=1
		})
	})

}



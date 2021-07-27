package main

import (
	"math/rand"
	"time"

	ruthSkiplist "github.com/Ruth-Seven/skiplist"
)

type Item struct {
	key   int
	value []byte
}

func (item Item) ExtentedKey() float64 {
	return float64(item.key)
}

func ruthInserts(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   n - i,
			value: testByteString,
		})
	}
}

func ruthWorstInserts(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)
	defer timeTrack(time.Now(), n-1)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   i,
			value: testByteString,
		})
	}
}

func ruthRandomInserts(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)

	rList := rand.Perm(n)

	defer timeTrack(time.Now(), n)
	for _, e := range rList {
		list.Insert(Item{
			key:   e,
			value: testByteString,
		})
	}
}

func ruthAvgSearch(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   i,
			value: testByteString,
		})
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(float64(i))
	}
}

func ruthSearchEnd(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   i,
			value: testByteString,
		})
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(float64(n))
	}
}

func ruthDelete(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   i,
			value: testByteString,
		})
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Delete(float64(i))
	}
}

func ruthWorstDelete(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   i,
			value: testByteString,
		})
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Delete(float64(n - i))
	}
}

func ruthRandomDelete(n int) {
	list := ruthSkiplist.NewSkipList(ruthSkiplist.RecommendedEps)

	for i := 0; i < n; i++ {
		list.Insert(Item{
			key:   i,
			value: testByteString,
		})
	}

	rList := rand.Perm(n)

	defer timeTrack(time.Now(), n)

	for _, e := range rList {
		_ = list.Delete(float64(e))
	}
}

var ruthFunctions = []func(int){ruthInserts, ruthWorstInserts, ruthRandomInserts,
	ruthAvgSearch, ruthSearchEnd, ruthDelete, ruthWorstDelete, ruthRandomDelete}

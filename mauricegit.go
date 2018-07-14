package main

import (
	"fmt"
	mtSkiplist "github.com/MauriceGit/skiplist"
	"math/rand"
	"time"
)

type element uint64

func (e element) ExtractKey() float64 {
	return float64(e)
}
func (e element) String() string {
	return fmt.Sprintf("%03d", e)
}

func mtInserts(n int) {
	list := mtSkiplist.New()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}
}

func mtWorstInserts(n int) {
	list := mtSkiplist.New()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(element(i))
	}
}

func mtRandomInserts(n int) {
	list := mtSkiplist.New()

	rList := rand.Perm(n)

	defer timeTrack(time.Now(), n)
	for _, e := range rList {
		list.Insert(element(e))
	}
}

func mtAvgSearch(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(element(i))
	}
}

func mtSearchEnd(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(element(n))
	}
}

func mtDelete(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Delete(element(i))
	}
}

func mtWorstDelete(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Delete(element(n - i))
	}
}

func mtRandomDelete(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	rList := rand.Perm(n)

	defer timeTrack(time.Now(), n)

	for _, e := range rList {
		list.Delete(element(e))
	}
}

var mtFunctions = []func(int){mtInserts, mtWorstInserts, mtRandomInserts,
	mtAvgSearch, mtSearchEnd, mtDelete, mtWorstDelete, mtRandomDelete}

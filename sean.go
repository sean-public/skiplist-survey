package main

import (
	"math/rand"
	"time"

	seaSkiplist "github.com/sean-public/fast-skiplist"
)

func seanInserts(n int) {
	list := seaSkiplist.New()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}
}

func seanWorstInserts(n int) {
	list := seaSkiplist.New()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}
}

func seanRandomInserts(n int) {
	list := seaSkiplist.New()

	rList := rand.Perm(n)

	defer timeTrack(time.Now(), n)
	for _, e := range rList {
		list.Set(float64(e), testByteString)
	}
}

func seanAvgSearch(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(float64(i))
	}
}

func seanSearchEnd(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(float64(n - 1))
	}
}

func seanDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(float64(i))
	}
}

func seanWorstDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(float64(n - i))
	}
}

func seanRandomDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}

	rList := rand.Perm(n)

	defer timeTrack(time.Now(), n)

	for _, e := range rList {
		_ = list.Remove(float64(e))
	}
}

var seanFunctions = []func(int){seanInserts, seanWorstInserts, seanRandomInserts,
	seanAvgSearch, seanSearchEnd, seanDelete, seanWorstDelete, seanRandomDelete}

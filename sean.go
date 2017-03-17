package main

import (
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

func seanAvgSearch(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(float64(i))
	}
}

func seanSearchEnd(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(float64(n))
	}
}

func seanDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(float64(i))
	}
}

func seanWorstDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(float64(n - i))
	}
}

var seanFunctions = []func(int){seanInserts, seanWorstInserts,
	seanAvgSearch, seanSearchEnd, seanDelete, seanWorstDelete}

package main

import (
	"time"

	rysSkiplist "github.com/ryszard/goskiplist/skiplist"
)

func ryszardInserts(n int) {
	list := rysSkiplist.NewIntMap()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}
}

func ryszardWorstInserts(n int) {
	list := rysSkiplist.NewIntMap()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}
}

func ryszardAvgSearch(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Get(i)
	}
}

func ryszardSearchEnd(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Get(n)
	}
}

func ryszardDelete(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(i)
	}
}

func ryszardWorstDelete(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(n - i)
	}
}

var ryszardFunctions = []func(int){ryszardInserts, ryszardWorstInserts,
	ryszardAvgSearch, ryszardSearchEnd, ryszardDelete, ryszardWorstDelete}

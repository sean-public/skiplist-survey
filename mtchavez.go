package main

import (
	"time"

	mtcSkiplist "github.com/mtchavez/skiplist"
)

func mtchavezInserts(n int) {
	list := mtcSkiplist.NewList()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}
}

func mtchavezWorstInserts(n int) {
	list := mtcSkiplist.NewList()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}
}

func mtchavezAvgSearch(n int) {
	list := mtcSkiplist.NewList()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Search(i)
	}
}

func mtchavezSearchEnd(n int) {
	list := mtcSkiplist.NewList()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Search(n)
	}
}

func mtchavezDelete(n int) {
	list := mtcSkiplist.NewList()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Delete(i)
	}
}

func mtchavezWorstDelete(n int) {
	list := mtcSkiplist.NewList()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Delete(n - i)
	}
}

var mtchavezFunctions = []func(int){mtchavezInserts, mtchavezWorstInserts,
	mtchavezAvgSearch, mtchavezSearchEnd, mtchavezDelete, mtchavezWorstDelete}

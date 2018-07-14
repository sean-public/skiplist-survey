package main

import (
	mtcSkiplist "github.com/mtchavez/skiplist"
	"math/rand"
	"time"
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

func mtchavezRandomInserts(n int) {
	list := mtcSkiplist.NewList()
	rList := rand.Perm(n)
	defer timeTrack(time.Now(), n)

	for _, e := range rList {
		list.Insert(e, testByteString)
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

func mtchavezRandomDelete(n int) {
	list := mtcSkiplist.NewList()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	rList := rand.Perm(n)
	defer timeTrack(time.Now(), n)

	for _, e := range rList {
		_ = list.Delete(e)
	}
}

var mtchavezFunctions = []func(int){mtchavezInserts, mtchavezWorstInserts, mtchavezRandomInserts,
	mtchavezAvgSearch, mtchavezSearchEnd, mtchavezDelete, mtchavezWorstDelete, mtchavezRandomDelete}

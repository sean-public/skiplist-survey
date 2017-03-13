package main

import (
	"time"

	colSkiplist "github.com/golang-collections/go-datastructures/slice/skip"
)

type mockEntry uint64

func (me mockEntry) Compare(other colSkiplist.Entry) int {
	otherU := other.(mockEntry)
	if me == otherU {
		return 0
	}

	if me > otherU {
		return 1
	}

	return -1
}

func newMockEntry(key uint64) mockEntry {
	return mockEntry(key)
}

func colInserts(n int) {
	list := colSkiplist.New(uint(0))
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(newMockEntry(uint64(n - i)))
	}
}

func colWorstInserts(n int) {
	list := colSkiplist.New(uint(0))
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(newMockEntry(uint64(i)))
	}
}

func colAvgSearch(n int) {
	list := colSkiplist.New(uint(0))

	for i := 0; i < n; i++ {
		list.Insert(newMockEntry(uint64(i)))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.GetWithPosition(newMockEntry(uint64(i)))
	}
}

func colSearchEnd(n int) {
	list := colSkiplist.New(uint(0))

	for i := 0; i < n; i++ {
		list.Insert(newMockEntry(uint64(i)))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.GetWithPosition(newMockEntry(uint64(n)))
	}
}

func colDelete(n int) {
	list := colSkiplist.New(uint(0))

	for i := 0; i < n; i++ {
		list.Insert(newMockEntry(uint64(i)))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Delete(newMockEntry(uint64(i)))
	}
}

func colWorstDelete(n int) {
	list := colSkiplist.New(uint(0))

	for i := 0; i < n; i++ {
		list.Insert(newMockEntry(uint64(i)))
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Delete(newMockEntry(uint64(n - i)))
	}
}

var colFunctions = []func(int){colInserts, colWorstInserts,
	colAvgSearch, colSearchEnd, colDelete, colWorstDelete}

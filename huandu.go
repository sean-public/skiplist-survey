package main

import (
	"time"

	huaSkiplist "github.com/huandu/skiplist"
)

func huanduInserts(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}
}

func huanduWorstInserts(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}
}

func huanduAvgSearch(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(i)
	}
}

func huanduSearchEnd(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(n)
	}
}

func huanduDelete(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(i)
	}
}

func huanduWorstDelete(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(n - i)
	}
}

var huanduFunctions = []func(int){huanduInserts, huanduWorstInserts,
	huanduAvgSearch, huanduSearchEnd, huanduDelete, huanduWorstDelete}

package main

import (
	"time"

	zheSkiplist "github.com/zhenjl/skiplist"
)

func zhenjlInserts(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}
}

func zhenjlWorstInserts(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}
}

func zhenjlAvgSearch(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Select(i)
	}
}

func zhenjlSearchEnd(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Select(n)
	}
}

func zhenjlDelete(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(i)
	}
}

func zhenjlWorstDelete(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(n - i)
	}
}

var zhenjlFunctions = []func(int){zhenjlInserts, zhenjlWorstInserts,
	zhenjlAvgSearch, zhenjlSearchEnd, zhenjlDelete, zhenjlWorstDelete}

package concurrency

import (
	"fmt"
	"reflect"
	"testing"
)

func webChecker(url string) bool {
	if url == "www.baidu.com" {
		return true
	}
	return false
}

//go test -race  || check race condition

func BenchmarkCheckWebsiteUrls(t *testing.B) {
	var urls = []string{
		"www.baidu.com", "www.google.com",
	}
	var want = map[string]bool{
		"www.baidu.com":  true,
		"www.google.com": false,
	}

	//重置函数执行开始时间
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		got := CheckWebsiteUrls(webChecker, urls)
		if !reflect.DeepEqual(want, got) {
			t.Fatal(fmt.Sprintf("want %v go %v", want, got))
		}
	}
}

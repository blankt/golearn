package concurrency

type WebChecker func(string) bool

type CheckerResult struct {
	u    string
	pass bool
}

var resultChan = make(chan CheckerResult)

func CheckWebsiteUrls(webChecker WebChecker, urls []string) map[string]bool {
	var result = map[string]bool{}

	for _, url := range urls {
		go func(u string) {
			resultChan <- CheckerResult{u, webChecker(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		result[r.u] = r.pass
	}
	return result
}

package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			//results[url] = wc(url)
			resultChannel <- result{url, wc(url)}
		}()
	}
	for _, _ = range urls {
		rslt := <-resultChannel
		results[rslt.string] = rslt.bool
	}

	return results
}

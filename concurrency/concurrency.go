package concurrency

// WebsiteChecker heathcheck
type WebsiteChecker func(string) bool

// CheckWebsites heathcheck
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

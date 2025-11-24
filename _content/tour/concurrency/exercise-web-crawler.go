//go:build OMIT

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch возвращает тело URL и
	// срез URL, найденных на этой странице.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl использует fetcher для рекурсивного обхода
// страниц, начиная с url, до максимальной глубины.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Получать URL параллельно.
	// TODO: Не получать один и тот же URL дважды.
	// Эта реализация не делает ни того, ни другого:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("найдено: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher — это Fetcher, который возвращает предопределенные результаты.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("не найдено: %s", url)
}

// fetcher — это заполненный fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"Язык программирования Go",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Пакеты",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Пакет fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Пакет os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

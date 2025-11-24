//go:build OMIT

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch возвращает тело URL и
	// срез URL, найденных на этой странице.
	Fetch(url string) (body string, urls []string, err error)
}

// fetched отслеживает URL, которые были (или находятся в процессе) получены.
// Блокировка должна удерживаться при чтении из или записи в хэш-таблицу.
// См. https://golang.org/ref/spec#Struct_types раздел о встроенных типах.
var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("загрузка URL в процессе") // сигнальное значение

// Crawl использует fetcher для рекурсивного обхода
// страниц, начиная с url, до максимальной глубины.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Printf("<- Готово с %v, глубина 0.\n", url)
		return
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("<- Готово с %v, уже получено.\n", url)
		return
	}
	// Мы помечаем url как загружаемый, чтобы избежать повторной загрузки другими одновременно.
	fetched.m[url] = loading
	fetched.Unlock()

	// Мы загружаем его параллельно.
	body, urls, err := fetcher.Fetch(url)

	// И обновляем статус в синхронизированной зоне.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Printf("<- Ошибка на %v: %v\n", url, err)
		return
	}
	fmt.Printf("Найдено: %s %q\n", url, body)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Обход дочернего %v/%v из %v : %v.\n", i, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}
	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Ожидание дочернего %v.\n", url, i, len(urls), u)
		<-done
	}
	fmt.Printf("<- Готово с %v\n", url)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)

	fmt.Println("Статистика получения\n--------------")
	for url, err := range fetched.m {
		if err != nil {
			fmt.Printf("%v не удалось: %v\n", url, err)
		} else {
			fmt.Printf("%v было получено\n", url)
		}
	}
}

// fakeFetcher — это Fetcher, который возвращает предопределенные результаты.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("не найдено: %s", url)
}

// fetcher — это заполненный fakeFetcher.
var fetcher = &fakeFetcher{
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

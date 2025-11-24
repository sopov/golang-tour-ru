//go:build OMIT

// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// Значение успешно отправлено.
	case <-quit:
		return
	}
	walkImpl(t.Right, ch, quit)
}

// Walk обходит дерево t, отправляя все значения
// из дерева в канал ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	walkImpl(t, ch, quit)
	close(ch)
}

// Same определяет, содержат ли деревья
// t1 и t2 одинаковые значения.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, w1, quit)
	go Walk(t2, w2, quit)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("ПРОЙДЕНО")
	} else {
		fmt.Println("НЕ ПРОЙДЕНО")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("ПРОЙДЕНО")
	} else {
		fmt.Println("НЕ ПРОЙДЕНО")
	}
}

//go:build OMIT

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk обходит дерево t, отправляя все значения
// из дерева в канал ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Нужно закрыть канал здесь
	close(ch)
}

// Same определяет, содержат ли деревья
// t1 и t2 одинаковые значения.
// ПРИМЕЧАНИЕ: Реализация утекает горутинами, когда деревья различны.
// См. binarytrees_quit.go для лучшего решения.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

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

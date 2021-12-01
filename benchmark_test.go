package main

import (
	"testing"
)

type node struct {
	child []*node
}

func generateTree(h int) *node {
	if h <= 0 {
		return nil
	}
	n := new(node)
	for i := 0; i < 2; i++ {
		c := generateTree(h - 1)
		if c != nil {
			n.child = append(n.child, c)
		}
	}
	return n
}

func bfs(n *node) {
	queue := []*node{n}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		queue = append(queue, top.child...)
	}
}

func dfs(n *node) {
	stack := []*node{n}
	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		stack = append(stack, top.child...)
	}
}

func Benchmark(b *testing.B) {
	tree := generateTree(10)
	b.Run("queue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bfs(tree)
		}
	})
	b.Run("stack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dfs(tree)
		}
	})
}

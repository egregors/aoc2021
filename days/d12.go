/*
	https://adventofcode.com/2021/day/12
*/

//nolint:revive //meh
package days

import (
	"fmt"
	"strings"
)

type Set map[string]bool

type Graph struct {
	vertices     Set
	adjacentList map[string]Set
}

func (g *Graph) String() string {
	vertices := ""
	for k := range g.vertices {
		vertices += k + " "
	}

	adjacent := ""
	for k, v := range g.adjacentList {
		inner := ""
		for k := range v {
			inner += k + ","
		}
		adjacent += "\t" + k + ": [" + inner + "]\n"
	}

	return fmt.Sprintf(
		"[%s]\n{\n%s}\n",
		vertices,
		adjacent,
	)
}

func (g *Graph) addVertex(v string) {
	if !g.vertices[v] {
		g.vertices[v] = true
		g.adjacentList[v] = make(Set)
	}
}

func (g *Graph) addEdge(v1, v2 string, directed bool) {
	if v1 != v2 {
		g.addVertex(v1)
		g.addVertex(v2)
		g.adjacentList[v1][v2] = true
		if directed {
			g.adjacentList[v2][v1] = true
		}
	}
}

func newGraph(lines map[string][]string) *Graph {
	g := &Graph{
		vertices:     make(Set),
		adjacentList: make(map[string]Set),
	}

	for k, v := range lines {
		for _, x := range v {
			d := true
			if x == "end" {
				d = false
			}
			if k == "start" {
				d = false
			}
			g.addEdge(k, x, d)
		}
	}

	return g
}

type pair struct {
	from, to string
}

func getLines(xs []string) map[string][]string {
	pairs := make([]pair, len(xs))
	for i, x := range xs {
		l := strings.Split(x, "-")
		pairs[i] = pair{l[0], l[1]}
	}

	lines := make(map[string][]string)
	for _, p := range pairs {
		lines[p.from] = append(lines[p.from], p.to)
	}
	return lines
}

func passagePathingP1(xs []string) int {
	lines := getLines(xs)
	g := newGraph(lines)

	count := 0
	for edge := range g.adjacentList["start"] {
		count += visit(g, edge, map[string]bool{"start": true})
	}

	return count
}

func passagePathingP2(xs []string) int {
	lines := getLines(xs)
	g := newGraph(lines)

	paths := make(map[string]bool)

	for n := range g.adjacentList {
		if n == "start" || n == "end" || !isLower(n) {
			continue
		}

		for e := range g.adjacentList["start"] {

			visit2(g, e, n, false, "start", paths)

		}

	}

	return len(paths)
}

func visit(g *Graph, startAt string, seen map[string]bool) (count int) {
	if startAt == "end" {
		return 1
	}

	seen[startAt] = true
	for to := range g.adjacentList[startAt] {
		_, visited := seen[to]
		if visited && isLower(to) {
			continue
		}

		clonedSeen := map[string]bool{}
		for k, v := range seen {
			clonedSeen[k] = v
		}
		count += visit(g, to, clonedSeen)
	}

	return count
}

func visit2(g *Graph, startAt, forTwice string, isVisitedTwice bool, path string, paths map[string]bool) {
	if startAt == "end" {
		path += ",end"
		paths[path] = true
		return
	}

	for to := range g.adjacentList[startAt] {
		secondVisit := isVisitedTwice

		visited := strings.Contains(path, to)
		if visited && isLower(to) {
			if to == forTwice {
				if isVisitedTwice {
					continue
				}
				secondVisit = true
			} else {
				continue
			}
		}

		visit2(g, to, forTwice, secondVisit, fmt.Sprintf("%s,%s", path, startAt), paths)
	}
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

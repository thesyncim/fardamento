package main

import (
	"strings"
)

type Combinacoes []string

func NewCombination() *Combinacoes {
	c := new(Combinacoes)
	return c
}

func (c *Combinacoes) removeindexs(ids []int) {
	w := 0
loop:
	for i, x := range *c {
		for _, id := range ids {
			if id == i {
				continue loop
			}
		}
		(*c)[w] = x
		w++
	}
	*c = (*c)[:w]

}

func (c *Combinacoes) Add(separator string, s ...string) {
	tmpstring := strings.Join(s, separator)
	*c = append(*c, tmpstring)
}

//join elements
func (c *Combinacoes) J(s []string, separator string) string {
	return strings.Join(s, separator)

}

func (c *Combinacoes) Len() int {
	return len(*c)
}

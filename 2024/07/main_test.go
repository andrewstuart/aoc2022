package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	f, err := os.OpenFile("./test.txt", os.O_RDONLY, 0400)
	rq.NoError(err)

	out := aoc(bufio.NewReader(f))

	asrt.Equal(11387, out)
}

func TestGen(t *testing.T) {
	asrt := assert.New(t)
	out := gen(2, []string{"+", "*"})

	asrt.Equal([][]string{
		{"+", "+"},
		{"+", "*"},
		{"*", "+"},
		{"*", "*"},
	}, out)
}

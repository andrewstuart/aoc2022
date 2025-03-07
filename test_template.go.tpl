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

	f, err := os.OpenFile("{{ .Config.TestOutput }}", os.O_RDONLY, 0400)
	rq.NoError(err)

	out := aoc(bufio.NewReader(f))

	asrt.Equal({{ .Expected }}, out)
}

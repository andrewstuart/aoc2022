package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func XTestIt(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(95437, aoc(strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)))
}

func TestFS(t *testing.T) {
	asrt := assert.New(t)
	root := readFS(strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`))
	asrt.Equal(48381165, root.du())
}

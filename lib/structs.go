package lib

import (
	"strings"
)

type Foo struct {
	First string
	Last  string
}

func (f *Foo) Firstifier() string {
	return strings.Split(f.First, "")[0] + "," + strings.Split(f.Last, "")[0]
}

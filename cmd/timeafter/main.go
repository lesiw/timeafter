package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/timeafter"
)

func main() { singlechecker.Main(timeafter.Analyzer) }

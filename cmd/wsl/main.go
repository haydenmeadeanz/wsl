package main

import (
	"github.com/haydenmeadeanz/wsl"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(wsl.NewAnalyzer(nil))
}

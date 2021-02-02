package main

import (
	_ "study-goframe/boot"
	_ "study-goframe/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}

package main

import (
	"github.com/sunjiangjun/go-demo/util"
	"github.com/sunjiangjun/xlog"
)

func main() {

	x := xlog.NewXLogger()
	x.Info("hello world")

	util.MaxInt(1, 2)
}

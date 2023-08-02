package main

import (
	"os"
	"runtime/pprof"

	"example.com/dep"
)

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 0; i < 5; i++ {
		nums := dep.Generate(n)
		dep.Bubble(nums)
		n *= 10
	}
}

/*
go tool pprof cpu.pprof   以交互命令行分析数据
go tool pprof -no_browser -http=0.0.0.0:9999 cpu.pprof   以web-ui模式分析
*/

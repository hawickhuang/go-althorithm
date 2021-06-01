package main

import (
	"github.com/hawickhuang/go-althorithm/mypprof"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func main()  {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	pprof.

	n := 10
	for i := 0; i<5;i++ {
		nums := mypprof.GetRandInt(n)
		_ = mypprof.BubbleSort(nums)
		n *= 10
	}
}

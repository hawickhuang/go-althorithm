package main

import (
	"github.com/hawickhuang/go-althorithm/mypprof"
	"github.com/pkg/profile"
	_ "net/http/pprof"
	"os"
	//"runtime/pprof"
)

func main()  {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	//_ = pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	//n := 10
	//for i := 0; i<5;i++ {
	//	nums := mypprof.GetRandInt(n)
	//	_ = mypprof.BubbleSort(nums)
	//	n *= 10
	//}
	mypprof.Concat2(100)
}

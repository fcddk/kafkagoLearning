package computer_status

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func CpuStatus() string {
	fmt.Printf("cpu status \n")
	return getCpu()
}

func getCpu() (s string) {
	s = "cpu:  "
	cusage, _ := cpu.Percent(time.Second, false)
	v, err := cpu.Info()
	if  err != nil{
		panic(err)
	}

	if len(v) > 1 {
		for _, sup_cpu := range v {
			modeName := sup_cpu.ModelName
			cores := sup_cpu.Cores
			s = "modeName:" + string(modeName)
			s = "cores:" + string(cores)
		}
	}else {
		s = s + v[0].String()
	}
	s = s + "used:" + string(int(cusage[0]))
	return s
}

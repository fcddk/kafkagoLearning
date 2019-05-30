package computer_status

import (
	"github.com/shirou/gopsutil/mem"
)

func MemoryStatus() string {
	 return getMemory()
}

func getMemory() (s string) {
	v, err := mem.VirtualMemory()

	if err != nil {
		panic(err)
	}
	s = "mem:  "  + v.String()
	return s
}

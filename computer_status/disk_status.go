package computer_status

import (
	"github.com/shirou/gopsutil/disk"
)

func DiskStatus() string {
	return getDisk()
}

func getDisk() (s string) {
	s = "disk:  "
	v, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}
	s = s + v.String()
	return s
}

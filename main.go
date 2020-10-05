package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/capnm/sysinfo"
)

func main() {

	progress := 0

	for range time.Tick(time.Millisecond * 200) {
		progress++

		if progress > 5 {
			progress = 0
		}

		si := sysinfo.Get()

		total := si.TotalRam
		free := si.FreeRam
		used := total - free

		usedPercentage := used * 100 / total
		freePercentage := 100 - usedPercentage

		displayMem(int(usedPercentage), int(freePercentage), progress)
	}
}

func displayMem(used, free, progress int) {

	print("\033[H\033[2J")

	// Print Heading
	fmt.Printf("Used Mem: %d%%                 Free Mem: %d%%\n", used, free)

	// Print Bar
	fmt.Printf("+%s+\n", strings.Repeat("-", (used+free)/2))
	fmt.Printf("|%s%s |\n", strings.Repeat("|", used/2), strings.Repeat(" ", free/2))
	fmt.Printf("+%s+\n", strings.Repeat("-", (used+free)/2))

	fmt.Println(strings.Repeat(".", progress))
}

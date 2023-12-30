package format

import "fmt"

func startList() {
	print("  \u25CF ")
}

func startBold() {
	print("\033[1m")
}

func startCode() {
	print("\033[32m") // !FnColor - green
}

func startTick() {
	print("\033[48;5;242m\033[38;5;226m") // !BgColor - grey !FnColor - yellow
}

// func endBold() {
// 	print("\033[0m")
// }

// func endCode() {
// 	print("\033[0m")
// }

// func endTick() {
// 	print("\033[0m")
// }

func end() {
	print("\033[0m")
}

func resetStyle() {
	fmt.Print("\033[0m\n")
}

func HandleFormatStart(format string) {
	switch format {
	case "list1":
		startList()
	case "list2":
		startList()
	case "bold":
		startBold()
	case "code":
		startCode()
	case "tick":
		startTick()
	}
}

func HandleFormatEnd(format string) {
	end()
}

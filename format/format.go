package format

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Format struct {
	name  string
	start string
	end   string
}

type FormatSelector struct {
	formats      []*Format
	currStr      string
	activeFormat *Format
}

func (fs *FormatSelector) addFormat(name string, start string, end string) {
	fs.formats = append(fs.formats, &Format{name, start, end})
}

func NewFormatSelector() *FormatSelector {
	fs := &FormatSelector{
		formats: []*Format{},
	}

	fs.addFormat("list1", "* ", "")
	fs.addFormat("list2", "- ", "")
	fs.addFormat("h2", "## ", "")
	fs.addFormat("code", "```", "```")
	fs.addFormat("bold", "**", "**")
	fs.addFormat("tick", "`", "`")

	fs.currStr = ""

	return fs
}

func (fs *FormatSelector) isFormatStart() (bool, string) {
	for _, format := range fs.formats {
		if strings.HasPrefix(fs.currStr, format.start) {
			if !(length(format.end) == 0) {
				fs.activeFormat = format
			}

			HandleFormatStart(format.name)
			return true, format.start
		}
	}
	return false, ""
}

func (fs *FormatSelector) isFormatEnd() (bool, string) {
	if strings.HasPrefix(fs.currStr, fs.activeFormat.end) {
		endFormat := fs.activeFormat.end
		HandleFormatEnd(fs.activeFormat.name)
		fs.activeFormat = nil
		return true, endFormat
	}

	return false, ""
}

func (fs *FormatSelector) PrintFormat(ch rune) {
	fs.currStr += string(ch)

	if length(fs.currStr) < 3 {
		return
	}

	if fs.activeFormat != nil {
		if isMatch, trimStr := fs.isFormatEnd(); isMatch {
			fs.currStr = fs.currStr[length(trimStr):]
		}
	} else {
		if isMatch, trimStr := fs.isFormatStart(); isMatch {
			fs.currStr = fs.currStr[length(trimStr):]
		}
	}

	if length(fs.currStr) >= 3 {
		fmt.Print(fs.currStr[:1])
		fs.currStr = fs.currStr[1:]
	}
}

func (fs *FormatSelector) Flush() {
	print(fs.currStr)
}

func (fs *FormatSelector) Close() {
	fs = nil
	resetStyle()
}

func length(s string) int {
	return utf8.RuneCountInString(s)
}

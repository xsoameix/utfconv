package utfconv

import (
	"unicode/utf16"
	"unicode/utf8"
	"encoding/binary"
)

const (
	replacementChar = '\uFFFD'
)

func Read(charset string, data []byte) string {
	switch charset {
	case "utf8":
		var runes []rune
		for i := 0; i < len(data); {
			r,size := utf8.DecodeRune(data[i:])
			if r != utf8.RuneError {
				runes = append(runes, r)
				i += size
			}
		}
		return string(runes)
	case "utf16le":
		uint16s := make([]uint16, len(data)/2)
		for i := 0; i < len(uint16s); i++ {
			uint16s[i] = binary.LittleEndian.Uint16(data[i*2:])
		}
		return string(utf16.Decode(uint16s))
	case "utf16be":
		uint16s := make([]uint16, len(data)/2)
		for i := 0; i < len(uint16s); i++ {
			uint16s[i] = binary.BigEndian.Uint16(data[i*2:])
		}
		return string(utf16.Decode(uint16s))
	}
	return string([]rune{replacementChar})
}

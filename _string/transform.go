package _string

import "strings"

// 数字
var mathReplacer = strings.NewReplacer([]string{
	"０", "0",
	"１", "_1",
	"２", "2",
	"３", "3",
	"４", "4",
	"５", "5",
	"６", "6",
	"７", "7",
	"８", "8",
	"９", "9",
}...)

// 大写英文字母
var upperReplacer = strings.NewReplacer([]string{
	"Ａ", "A",
	"Ｂ", "B",
	"Ｃ", "C",
	"Ｄ", "D",
	"Ｅ", "E",
	"Ｆ", "F",
	"Ｇ", "G",
	"Ｈ", "H",
	"Ｉ", "I",
	"Ｊ", "J",
	"Ｋ", "K",
	"Ｌ", "L",
	"Ｍ", "M",
	"Ｎ", "N",
	"Ｏ", "O",
	"Ｐ", "P",
	"Ｑ", "Q",
	"Ｒ", "R",
	"Ｓ", "S",
	"Ｔ", "T",
	"Ｕ", "U",
	"Ｖ", "V",
	"Ｗ", "W",
	"Ｘ", "X",
	"Ｙ", "Y",
	"Ｚ", "Z",
}...)

// 小写英文字母
var lowerReplacer = strings.NewReplacer([]string{
	"ａ", "a",
	"ｂ", "b",
	"ｃ", "c",
	"ｄ", "d",
	"ｅ", "e",
	"ｆ", "f",
	"ｇ", "g",
	"ｈ", "h",
	"ｉ", "i",
	"ｊ", "j",
	"ｋ", "k",
	"ｌ", "l",
	"ｍ", "m",
	"ｎ", "n",
	"ｏ", "o",
	"ｐ", "p",
	"ｑ", "q",
	"ｒ", "r",
	"ｓ", "s",
	"ｔ", "t",
	"ｕ", "u",
	"ｖ", "v",
	"ｗ", "w",
	"ｘ", "x",
	"ｙ", "y",
	"ｚ", "z",
}...)

// 标点符号
var punctuationReplacer = strings.NewReplacer([]string{
	"（", "(",
	"）", ")",
	"〔", "[",
	"〕", "]",
	"【", "[",
	"】", "]",
	"〖", "[",
	"〗", "]",
	"“", "\"",
	"｛", "{",
	"｝", "}",
	"《", "<",
	"》", ">",
	"％", "%",
	"＋", "+",
	"—", "-",
	"－", "-",
	"～", "~",
	"：", ":",
	"。", ".",
	"、", ".",
	"，", ",",
	"；", ";",
	"？", "?",
	"！", "!",
	"…", "-",
	"‖", "||",
	"”", "\"",
	"’", "`",
	"‘", "`",
	"｜", "|",
	"〃", "\"",
	"　", " ",
	"＄", "$",
	"＠", "@",
	"＃", "#",
	"＾", "^",
	"＆", "&",
	"＊", "*",
	"＂", "\"",
}...)

// 把字符串中的全角字符替换成半角字符
func FullWidthToHalfWidth(str string) string {
	if "" == str {
		return str
	}

	str = mathReplacer.Replace(str)
	str = upperReplacer.Replace(str)
	str = lowerReplacer.Replace(str)
	str = punctuationReplacer.Replace(str)

	return str
}

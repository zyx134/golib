package _string

import (
	"bytes"
	"strings"
	"unicode"
)

/*这是对strings包的封装*/

/**
字符串比较
*/
//	用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -_1 ，反之返回 _1 。不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
func Compare(a, b string) int {
	//a := "gopher"
	//b := "hello world"
	//fmt.Println(strings.Compare(a, b))
	//fmt.Println(strings.Compare(a, a))
	//fmt.Println(strings.Compare(b, a))
	return strings.Compare(a, b)
}

//	EqualFold 函数，计算 s 与 t 忽略字母大小写后是否相等。
func EqualFold(s, t string) bool {
	//fmt.Println(strings.EqualFold("GO", "go"))
	//fmt.Println(strings.EqualFold("壹", "一"))
	return strings.EqualFold(s, t)
}

/**
是否存在某个字符或子串
*/

// 子串 substr 在 s 中，返回 true
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// chars 中任何一个 Unicode 代码点在 s 中，返回 true
func ContainsAny(s, chars string) bool {
	//fmt.Println(strings.ContainsAny("team", "i"))
	//fmt.Println(strings.ContainsAny("failure", "u & i"))
	//fmt.Println(strings.ContainsAny("in failure", "s g"))
	//fmt.Println(strings.ContainsAny("foo", ""))
	//fmt.Println(strings.ContainsAny("", ""))
	return strings.ContainsAny(s, chars)
}

// Unicode 代码点 r 在 s 中，返回 true
func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

/*子串出现次数 ( 字符串匹配 )*/
func Count(s, sep string) int {
	//在 Count 的实现中，处理了几种特殊情况，属于字符匹配预处理的一部分。这里要特别说明一下的是当 sep 为空时，Count 的返回值是：utf8.RuneCountInString(s) + _1
	//fmt.Println(strings.Count("cheese", "e"))
	//fmt.Println(len("谷歌中国"))
	//fmt.Println(strings.Count("谷歌中国", ""))
	//Count 是计算子串在字符串中出现的无重叠的次数，比如：
	//fmt.Println(strings.Count("fivevev", "vev"))
	return strings.Count(s, sep)
}

/*字符串分割为[]string*/

//Fields 用一个或多个连续的空格分隔字符串 s，返回子字符串的数组（slice）。如果字符串 s 只包含空格，则返回空列表 ([]string 的长度为 0）。其中，空格的定义是 unicode.IsSpace，之前已经介绍过。
//常见间隔符包括：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
func Fields(s string) []string {
	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	return strings.Fields(s)
}

//FieldsFunc 用这样的 Unicode 代码点 c 进行分隔：满足 f(c) 返回 true。该函数返回[]string。如果字符串 s 中所有的代码点 (unicode code points) 都满足 f(c) 或者 s 是空，则 FieldsFunc 返回空 slice。
func FieldsFunc(s string, f func(rune) bool) []string {
	//fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
	return strings.FieldsFunc(s, f)
}

//这四个函数都是通过 sep 进行分割，返回[]string。如果 sep 为空，相当于分成一个个的 UTF-8 字符，如 Split("abc","")，得到的是[a b c]。
//Split(s, sep) 和 SplitN(s, sep, -_1) 等价；SplitAfter(s, sep) 和 SplitAfterN(s, sep, -_1) 等价。
//Split 会将 s 中的 sep 去掉，而 SplitAfter 会保留 sep。
//带 N 的方法可以通过最后一个参数 n 控制返回的结果中的 slice 中的元素个数，当 n < 0 时，返回所有的子字符串；当 n == 0 时，返回的结果是 nil；当 n > 0 时，表示返回的 slice 中最多只有 n 个元素，其中，最后一个元素不会分割，比如：
func Split(s, sep string) []string {
	//fmt.Printf("%q\n", strings.Split("foo,bar,baz", ",")) //["foo" "bar" "baz"]
	return strings.Split(s, sep)
	//官方栗子
	//fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //["a" "b" "c"]
	//fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) //["" "man " "plan " "canal panama"]
	//fmt.Printf("%q\n", strings.Split(" xyz ", "")) //[" " "x" "y" "z" " "]
	//fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins")) //[""]
}
func SplitAfter(s, sep string) []string {
	//fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ",")) //["foo," "bar," "baz"]
	return strings.SplitAfter(s, sep)
}
func SplitN(s, sep string, n int) []string {
	//fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2)) //["foo" "bar,baz"]
	return strings.SplitN(s, sep, n)
}
func SplitAfterN(s, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}

/**
字符串是否有某个前缀或后缀
如果 prefix 或 suffix 为 "" , 返回值总是 true。
*/

// s 中是否以 prefix 开始
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// s 中是否以 suffix 结尾
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

/**
字符或子串在字符串中出现的位置
*/
// 在 s 中查找 sep 的第一次出现，返回第一次出现的索引
func Index(s, sep string) int {
	return strings.Index(s, sep)
}

// 在 s 中查找字节 c 的第一次出现，返回第一次出现的索引
func IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

// chars 中任何一个 Unicode 代码点在 s 中首次出现的位置
func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// Unicode 代码点 r 在 s 中第一次出现的位置
func IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

// 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
func IndexFunc(s string, f func(rune) bool) int {
	//han := func(c rune) bool {
	//	return unicode.Is(unicode.Han, c) // 汉字
	//}
	//fmt.Println(strings.IndexFunc("Hello, world", han)) //-_1
	//fmt.Println(strings.IndexFunc("Hello, 世界", han)) //7
	return strings.IndexFunc(s, f)
}

// 有三个对应的查找最后一次出现的位置
func LastIndex(s, sep string) int {
	return strings.LastIndex(s, sep)
}
func LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}
func LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}
func LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

/**
字符串 JOIN 操作
*/
func Join(a []string, sep string) string {
	//fmt.Println(Join([]string{"name=xxx", "age=xx"}, "&")) //name=xxx&age=xx
	return strings.Join(a, sep)
}

//自己实现
func JoinByMyself(str []string, sep string) string {
	// 特殊情况应该做处理
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}
	return buffer.String()
}

/**
字符串重复输出几次
*/
//将 s 重复 count 次，如果 count 为负数或返回值长度 len(s)*count 超出 string 上限会导致 panic，
func Repeat(s string, count int) string {
	//fmt.Println("ba" + strings.Repeat("na", 2)) //banana
	return strings.Repeat(s, count)
}

/**
字符替换
*/
//Map 函数，将 s 的每一个字符按照 mapping 的规则做映射替换，如果 mapping 返回值 <0 ，则舍弃该字符。该方法只能对每一个字符做处理，但处理方式很灵活，可以方便的过滤，筛选汉字等。
func Map(mapping func(rune) rune, s string) string {
	//mapping := func(r rune) rune {
	//	switch {
	//	case r >= 'A' && r <= 'Z': // 大写字母转小写
	//		return r + 32
	//	case r >= 'a' && r <= 'z': // 小写字母不处理
	//		return r
	//	case unicode.Is(unicode.Han, r): // 汉字换行
	//		return '\n'
	//	}
	//	return -_1 // 过滤所有非字母、汉字的字符
	//}
	//fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher...")) //hello world hello gopher
	return strings.Map(mapping, s)
}

/**
字符串子串替换
进行字符串替换时，考虑到性能问题，能不用正则尽量别用，应该用这里的函数。
如果我们希望一次替换多个，比如我们希望替换 This is <b>HTML</b> 中的 < 和 > 为 &lt; 和 &gt;，可以调用上面的函数两次。但标准库提供了另外的方法进行这种替换。
*/
// 用 new 替换 s 中的 old，一共替换 n 个。
// 如果 n < 0，则不限制替换次数，即全部替换
func Replace(s, old, new string, n int) string {
	//fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) //oinky oinky oink
	//fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -_1)) //moo moo moo
	return strings.Replace(s, old, new, n)
}

// 该函数内部直接调用了函数 Replace(s, old, new , -_1)
func ReplaceAll(s, old, new string) string {
	//fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) //moo moo moo
	return strings.ReplaceAll(s, old, new)
}

/**
大小写转换
ToLower,ToUpper 用于大小写转换。
ToLowerSpecial,ToUpperSpecial 可以转换特殊字符的大小写
汉字拼音有效
汉字无效
*/
func ToLower(s string) string {
	//fmt.Println(strings.ToLower("HELLO WORLD")) //hello world
	//fmt.Println(strings.ToLower("Ā Á Ǎ À")) //ā á ǎ à
	//fmt.Println(strings.ToLower("Önnek İş")) //önnek iş
	return strings.ToLower(s)
}
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
	//fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "壹")) //壹
	//fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD")) //hello world
	//fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")) //önnek iş
	return strings.ToLowerSpecial(c, s)
}
func ToUpper(s string) string {
	//fmt.Println(strings.ToUpper("hello world")) //HELLO WORLD
	//fmt.Println(strings.ToUpper("ā á ǎ à")) //Ā Á Ǎ À
	//fmt.Println(strings.ToUpper("örnek iş")) //ÖRNEK IŞ
	return strings.ToUpper(s)
}
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
	//fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "一")) //一
	//fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world")) //HELLO WORLD
	//fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş")) //ÖRNEK İŞ
	return strings.ToUpperSpecial(c, s)
}

/**
标题处理
*/
//将 s 每个单词的首字母大写，不处理该单词的后续字符
func Title(s string) string {
	//fmt.Println(strings.Title("hElLo wOrLd")) //HElLo WOrLd
	//fmt.Println(strings.Title("āáǎà ōóǒò êēéěè")) //Āáǎà Ōóǒò Êēéěè
	//fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir")) //Dünyanın Ilk Borsa Yapısı Aizonai Kabul Edilir
	return strings.Title(s)
}

//将 s 的每个字母大写
func ToTitle(s string) string {
	//fmt.Println(strings.ToTitle("hElLo wOrLd")) //HELLO WORLD
	//fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè")) //ĀÁǍÀ ŌÓǑÒ ÊĒÉĚÈ
	//fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir")) //DÜNYANIN ILK BORSA YAPISI AIZONAI KABUL EDILIR
	return strings.ToTitle(s)
}

//将 s 的每个字母大写，并且会将一些特殊字母转换为其对应的特殊大写字母
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
	//fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd")) //HELLO WORLD
	//fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè")) //ĀÁǍÀ ŌÓǑÒ ÊĒÉĚÈ
	//fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir")) //DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
	return strings.ToTitleSpecial(c, s)
}

/**
修剪
x := "!!!@@@你好,!@#$ Gophers###$$$"
*/
// 将 s 左侧和右侧中匹配 cutset 中的任一字符的字符去掉
//
func Trim(s string, cutset string) string {
	//fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-")) //你好,!@#$ Gophers
	return strings.Trim(s, cutset)
}

// 将 s 左侧的匹配 cutset 中的任一字符的字符去掉
func TrimLeft(s string, cutset string) string {
	//fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-")) //你好,!@#$ Gophers###$$$
	return strings.TrimLeft(s, cutset)
}

// 将 s 右侧的匹配 cutset 中的任一字符的字符去掉
func TrimRight(s string, cutset string) string {
	//fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-")) //!!!@@@你好,!@#$ Gophers
	return strings.TrimRight(s, cutset)
}

// 如果 s 的前缀为 prefix 则返回去掉前缀后的 string , 否则 s 没有变化。
func TrimPrefix(s, prefix string) string {
	//fmt.Println(strings.TrimPrefix(x, "!")) //!!@@@你好,!@#$ Gophers###$$$
	return strings.TrimPrefix(s, prefix)
}

// 如果 s 的后缀为 suffix 则返回去掉后缀后的 string , 否则 s 没有变化。
func TrimSuffix(s, suffix string) string {
	//fmt.Println(strings.TrimSuffix(x, "$")) //!!!@@@你好,!@#$ Gophers###$$
	return strings.TrimSuffix(s, suffix)
}

// 将 s 左侧和右侧的间隔符去掉。常见间隔符包括：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL)
func TrimSpace(s string) string {
	//fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) //Hello, Gophers
	return strings.TrimSpace(s)
}

// 将 s 左侧和右侧的匹配 f 的字符去掉
func TrimFunc(s string, f func(rune) bool) string {
	//f := func(r rune) bool {
	//    return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	//}
	//fmt.Println(strings.TrimFunc(x, f)) //你好
	return strings.TrimFunc(s, f)
}

// 将 s 左侧的匹配 f 的字符去掉
func TrimLeftFunc(s string, f func(rune) bool) string {
	//f := func(r rune) bool {
	//    return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	//}
	//fmt.Println(strings.TrimLeftFunc(x, f)) //你好,!@#$ Gophers###$$$
	return strings.TrimLeftFunc(s, f)
}

// 将 s 右侧的匹配 f 的字符去掉
func TrimRightFunc(s string, f func(rune) bool) string {
	//f := func(r rune) bool {
	//    return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	//}
	//fmt.Println(strings.TrimRightFunc(x, f)) //!!!@@@你好
	return strings.TrimRightFunc(s, f)
}

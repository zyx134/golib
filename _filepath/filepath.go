package _filepath

import (
	"path/filepath"
)

/**
解析路径名字符串
/home/polaris/studygolang.go，Dir 返回 /home/polaris，而 Base 返回 studygolang.go。
/home/polaris/studygolang/，Dir 返回 /home/polaris/studygolang（这与 Unix 中的 dirname 不一致，dirname 会返回 /home/polaris），而 Base 返回 studygolang
*/
//将一个路径名字符串分解成目录
func Dir(path string) string {
	return filepath.Dir(path)
}

//将一个路径名字符串分解成文件名
func Base(path string) string {
	return filepath.Base(path)
}

//获取文件名后缀
//扩展名是路径中最后一个从 . 开始的部分，包括 .。如果该元素没有 . 会返回空字符串。
func Ext(path string) string {
	return filepath.Ext(path)
}

/**
相对路径和绝对路径
*/
//判断是否绝对路径
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

//获取绝对路径
func Abs(path string) (string, error) {
	return filepath.Abs(path)
}

//获取相对路径
//将 basepath 和该路径用路径分隔符连起来的新路径在词法上等价于 targpath。也就是说，Join(basepath, Rel(basepath, targpath)) 等价于 targpath。如果成功执行，返回值总是相对于 basepath 的，即使 basepath 和 targpath 没有共享的路径元素。如果两个参数一个是相对路径而另一个是绝对路径，或者 targpath 无法表示为相对于 basepath 的路径，将返回错误。
func Rel(basepath, targpath string) (string, error) {
	//fmt.Println(filepath.Rel("/home/polaris/studygolang", "/home/polaris/studygolang/src/logic/topic.go"))
	// src/logic/topic.go <nil>
	//fmt.Println(filepath.Rel("/home/polaris/studygolang", "/data/studygolang"))
	// ../../../data/studygolang <nil>
	return filepath.Rel(basepath, targpath)
}

/**
路径的切分和拼接
*/
//Split 函数根据最后一个路径分隔符将路径 path 分隔为目录和文件名两部分（dir 和 file）。如果路径中没有路径分隔符，函数返回值 dir 为空字符串，file 等于 path；反之，如果路径中最后一个字符是 /，则 dir 等于 path，file 为空字符串。返回值满足 path == dir+file。dir 非空时，最后一个字符总是 /。
func Split(path string) (dir, file string) {
	//fmt.Println(filepath.Split("/home/polaris/studygolang")) // /home/polaris/ studygolang
	//fmt.Println(filepath.Split("/home/polaris/studygolang/")) //  /home/polaris/studygolang/ ""
	//fmt.Println(filepath.Split("studygolang")) // /home/polaris/ "" studygolang
	return filepath.Split(path)
}

//函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符。结果是经过 Clean 的，所有的空字符串元素会被忽略。对于拼接路径的需求，我们应该总是使用 Join 函数来处理
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

//我们需要分割 PATH 或 GOPATH 之类的环境变量（这些路径被特定于 OS 的列表分隔符连接起来），filepath.SplitList 就是这个用途
//注意，与 strings.Split 函数的不同之处是：对 ""，SplitList 返回[]string{}，而 strings.Split 返回 []string{""}。SplitList 内部调用的是 strings.Split。
func SplitList(path string) []string {
	return filepath.SplitList(path)
}

/**
规整化路径
*/
//Clean 函数通过单纯的词法操作返回和 path 代表同一地址的最短路径。
//
//它会不断的依次应用如下的规则，直到不能再进行任何处理：
//
//将连续的多个路径分隔符替换为单个路径分隔符
//剔除每一个 . 路径名元素（代表当前目录）
//剔除每一个路径内的 .. 路径名元素（代表父目录）和它前面的非 .. 路径名元素
//剔除开始于根路径的 .. 路径名元素，即将路径开始处的 /.. 替换为 /（假设路径分隔符是 /）
//返回的路径只有其代表一个根地址时才以路径分隔符结尾，如 Unix 的 / 或 Windows 的 C:\。
//如果处理的结果是空字符串，Clean 会返回 .，代表当前路径。
func Clean(path string) string {
	return filepath.Clean(path)
}

/**
符号链接指向的路径名
*/
func EvalSymlinks(path string) (string, error) {
	//// 在当前目录下创建一个 studygolang.txt 文件和一个 symlink 目录，在 symlink 目录下对 studygolang.txt 建一个符号链接 studygolang.txt.2
	//fmt.Println(filepath.EvalSymlinks("symlink/studygolang.txt.2")) //studygolang.txt <nil>
	//fmt.Println(os.Readlink("symlink/studygolang.txt.2")) //../studygolang.txt <nil>
	return filepath.EvalSymlinks(path)
}

/**
文件路径匹配
*/
//pattern:
//    { term }
//term:
//    '*'         匹配 0 或多个非路径分隔符的字符
//    '?'         匹配 _1 个非路径分隔符的字符
//    '[' [ '^' ] { character-range } ']'
//                  字符组（必须非空）
//    c           匹配字符 c（c != '*', '?', '\\', '['）
//    '\\' c      匹配字符 c
//character-range:
//    c           匹配字符 c（c != '\\', '-', ']'）
//    '\\' c      匹配字符 c
//    lo '-' hi   匹配区间[lo, hi]内的字符
func Match(pattern, name string) (matched bool, err error) {
	return filepath.Match(pattern, name)
}

//函数返回所有匹配了 模式字符串 pattern 的文件列表或者 nil（如果没有匹配的文件）。pattern 的语法和 Match 函数相同。pattern 可以描述多层的名字，如 /usr/*/bin/ed（假设路径分隔符是 /）。
//
//注意，Glob 会忽略任何文件系统相关的错误，如读目录引发的 I/O 错误。唯一的错误和 Match 一样，在 pattern 不合法时，返回 filepath.ErrBadPattern。返回的结果是根据文件名字典顺序进行了排序的。
//
//Glob 的常见用法，是读取某个目录下所有的文件，比如写单元测试时，读取 testdata 目录下所有测试数据：
//filepath.Glob("testdata/*.input")
func Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}

/**
遍历目录
*/
func Walk(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}

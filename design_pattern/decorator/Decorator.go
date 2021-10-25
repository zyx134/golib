package decorator

import (
	"log"
	"strings"
)

//
// 装饰模式
//    动态地给一个对象添加一些额外的职责，同时又不改变其结构
//

// 接口
type MessageBuilder interface {
	Build(messages ...string) string
}

// 基本信息构造器
type BaseMessageBuilder struct {
}

func (b *BaseMessageBuilder) Build(messages ...string) string {
	return strings.Join(messages, ",")
}

// 引号装饰器
type QuoteMessageBuilderDecorator struct {
	Builder MessageBuilder
}

func (q *QuoteMessageBuilderDecorator) Build(messages ...string) string {
	return "\"" + q.Builder.Build(messages...) + "\""
}

// 大括号装饰器
type BraceMessageBuilderDecorator struct {
	Builder MessageBuilder
}

func (b *BraceMessageBuilderDecorator) Build(messages ...string) string {
	return "{" + b.Builder.Build(messages...) + "}"
}

// 或者

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func Double(n int) int {
	return n * 2
}

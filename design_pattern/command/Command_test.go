package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	cmdHandler := &CommandHandler{CmdMap: make(map[string]Command)}
	postCtx := &CmdContext{
		CmdType: "post",
		Args:    "post",
	}
	getCtx := &CmdContext{
		CmdType: "get",
		Args:    "get",
	}
	nullCtx := &CmdContext{
		CmdType: "null",
		Args:    "null",
	}
	cmdHandler.Register("post", &PostCommand{})
	cmdHandler.Register("get", &GetCommand{})

	fmt.Println(cmdHandler.Handle(postCtx))
	fmt.Println(cmdHandler.Handle(getCtx))
	fmt.Println(cmdHandler.Handle(nullCtx))
}

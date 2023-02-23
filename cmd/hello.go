package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	helloCmd = cobra.Command{
		Use: "hello", //命令
		Run: func(cmd *cobra.Command, args []string) {
			// 真正 执行hello 命令运行的代码
			fmt.Println("hello", time.Now())
		}}
)

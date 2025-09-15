package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	// 处理减号作为命令的特殊情况
	// 在解析命令前检查参数，将 "-" 转换为 "sub"
	if len(os.Args) > 1 && os.Args[1] == "-" {
		fmt.Println(os.Args)
		os.Args[1] = "sub"
	}

	app := &cli.Command{
		Name:        "gocli",
		Usage:       "命令计算器",
		Description: "urface/cli 实现的计算器",
		Commands: []*cli.Command{
			//+
			{
				Name:    "add",
				Aliases: []string{"+"},
				Arguments: []cli.Argument{
					&cli.Float64Arg{Name: "num1", UsageText: "数字1"},
					&cli.Float64Arg{Name: "num2", UsageText: "数字2"},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					num1 := command.Float64Arg("num1")
					num2 := command.Float64Arg("num2")
					fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
					return nil
				},
			},
			//-
			{
				Name:    "sub",
				Aliases: []string{"-"},
				// 添加此配置以跳过对减号的flag解析
				Arguments: []cli.Argument{
					&cli.Float64Arg{Name: "num1", UsageText: "数字1"},
					&cli.Float64Arg{Name: "num2", UsageText: "数字2"},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					//这里这样用，命令行解析时对 - 符号的特殊处理导致的。在命令行工具中，- 通常被视为选项（flag）的前缀，而不是命令或参数，这会导致解析器无法正确识别你的减法命令。
					//在程序开头处理一下-的命令
					//执行会解析不到参数 .\clitest.exe - 1 2
					//				  0.00 -- 0.00 = 0.00
					num1 := command.Float64Arg("num1")
					num2 := command.Float64Arg("num2")
					fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
					return nil
				},
			},
			//*
			{
				Name:    "mul",
				Aliases: []string{"*"},
				Arguments: []cli.Argument{
					&cli.Float64Arg{Name: "num1", UsageText: "数字1"},
					&cli.Float64Arg{Name: "num2", UsageText: "数字2"},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					num1 := command.Float64Arg("num1")
					num2 := command.Float64Arg("num2")
					fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
					return nil
				},
			},
			// %
			{
				Name:    "mod",
				Aliases: []string{"%"},
				Arguments: []cli.Argument{
					&cli.Int64Arg{Name: "num1", UsageText: "数字1"},
					&cli.Int64Arg{Name: "num2", UsageText: "数字2"},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					num1 := command.Int64Arg("num1")
					num2 := command.Int64Arg("num2")
					if num2 == 0 {
						return fmt.Errorf("除数不能为0")
					}
					fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)
					return nil
				},
			},
			// /
			{
				Name:    "div",
				Aliases: []string{"/"},
				Arguments: []cli.Argument{
					&cli.Float64Arg{Name: "num1", UsageText: "数字1"},
					&cli.Float64Arg{Name: "num2", UsageText: "数字2"},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					num1 := command.Float64Arg("num1")
					num2 := command.Float64Arg("num2")
					if num2 == 0 {
						return fmt.Errorf("除数不能为0")
					}
					fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

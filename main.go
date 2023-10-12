package main

import (
	"fmt"
	"os"

	"evapp/bootstrap"
	btsConfig "evapp/config"
	"evapp/internal/cmd"
	"evapp/pkg/config"
	"evapp/pkg/console"

	"github.com/spf13/cobra"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "evapp",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有字命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)
			// 初始化 Logger
			bootstrap.SetupLogger()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
	)

	// 配置默认运行 web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)
	// 执行
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dbdoc/bootstrap"
	"dbdoc/conf"
	"dbdoc/log"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "启动数据字典测试环境",
	Long:  `启动数据字典测试环境，测试环境主要用于发布到服务器供前段同时联调测试使用`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		err := conf.InitConfig("./conf/conf-test.yaml")
		if err != nil {
			log.PrintlnError("conf init fail,error:%v", err)
			os.Exit(1)
		}
		bootstrap.Start()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

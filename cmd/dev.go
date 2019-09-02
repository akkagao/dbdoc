package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dbdoc/bootstrap"
	"dbdoc/conf"
	"dbdoc/log"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "启动数据字典本地开发环境",
	Long:  `启动数据字典本地开发环境`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev called")
		err := conf.InitConfig("./conf/conf-dev.yaml")
		if err != nil {
			log.PrintlnError("conf init fail,error:%v", err)
			os.Exit(1)
		}
		bootstrap.Start()
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}

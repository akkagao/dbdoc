package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dbdoc/bootstrap"
	"dbdoc/conf"
	"dbdoc/log"
)

// prodCmd represents the prod command
var prodCmd = &cobra.Command{
	Use:   "prod",
	Short: "启动数据字典正式线上系统",
	Long:  `启动数据字典正式线上系统。这是对外用户公开使用的。需要慎重操作`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prod called")
		err := conf.InitConfig("./conf/conf-prod.yaml")
		if err != nil {
			log.PrintlnError("conf init fail,error:%v", err)
			os.Exit(1)
		}
		bootstrap.Start()
	},
}

func init() {
	rootCmd.AddCommand(prodCmd)
}

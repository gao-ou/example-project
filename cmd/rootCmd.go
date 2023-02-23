package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// Used for flags.
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "example-project", //
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("这里是run 在执行代码")
	},
}

// cobra 初始化
func init() {
	// 执行cobra 初始化之前 先让viper 读配置
	// 这里的 initConfig 是一个函数，让 cobra 帮你执行函数
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// 绑定 hello 命令
	// go run main.go hello
	rootCmd.AddCommand(&helloCmd)
}

func initConfig() {
	// 判断 go run 时输入的 配置文件路径是否为空，不为空
	fmt.Println("cfgFile", cfgFile)

	if cfgFile != "" {
		// ，不为空
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// 为空
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

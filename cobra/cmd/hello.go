package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var helloFlags struct {
	message string
}

var helloCmd = &cobra.Command{
	Use:   "hello [NAME]",
	Args:  cobra.MinimumNArgs(1), // 引数が1つ未満の場合はエラー
	Short: "hello command",
	Long: `hello command.
say hello to your friends!`,
	Run: func(cmd *cobra.Command, args []string) {
		switch globalFlags.language {
		case "ja":
			sayHelloInJapanese(args)
		case "en":
			sayHelloInEnglish(args)
		default:
			fmt.Println("please specify the language (en/ja).")
			os.Exit(1)
		}
		fmt.Println(helloFlags.message)
	},
}

func sayHelloInJapanese(args []string) {
	for _, name := range args {
		fmt.Printf("%sさん, ", name)
	}
	fmt.Println("こんにちは!")
}

func sayHelloInEnglish(args []string) {
	fmt.Print("Hello,")
	for i, name := range args {
		fmt.Print(" " + name)
		if len(args) != 2 && i != len(args)-1 {
			fmt.Print(",")
		}
		if i == len(args)-2 {
			fmt.Print(" and")
		}
	}
	fmt.Println("!")
}

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().StringVarP(&helloFlags.message, "message", "m", "", "Help message for toggle")
}

// generate_completion.go
package main

import (
	"os"
	"github.com/spf13/cobra"
)

func SetupCompletion() {
	// cobra.Commandを仮で作成（フラグ解析のためだけに使う）
	var dummyCmd = &cobra.Command{
		Use: "gopassgen",
	}
	dummyCmd.Flags().Bool("generate-completions", false, "Generate shell completion scripts")
	dummyCmd.ParseFlags(os.Args[1:])

	generate, _ := dummyCmd.Flags().GetBool("generate-completions")
	if !generate {
		return
	}

	// 実際の生成処理
	os.MkdirAll("completions/bash", 0755)
	os.MkdirAll("completions/zsh", 0755)
	os.MkdirAll("completions/fish", 0755)
	os.MkdirAll("completions/powershell", 0755)

	dummyCmd.GenBashCompletionFileV2("completions/bash/gopassgen", true)
	dummyCmd.GenZshCompletionFile("completions/zsh/_gopassgen")
	dummyCmd.GenFishCompletionFile("completions/fish/gopassgen.fish", true)
	dummyCmd.GenPowerShellCompletionFile("completions/powershell/gopassgen.ps1")

	os.Exit(0) // 補完生成したらその場で終了（通常処理に行かない）
}

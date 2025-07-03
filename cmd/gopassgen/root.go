// cmd/gopassgen/root.go
package main

import (
    "fmt"
    "github.com/Jmiwa/gopassgen/internal/generator"
    "github.com/spf13/cobra"
    "os"
)

var (
    length     int
    useDigits  bool
    useSymbols bool
    count      int
)

// rootCmd defines CLI
var rootCmd = &cobra.Command{
    Use:   "gopassgen",
    Short: "Generate secure passwords",
    Run: func(cmd *cobra.Command, args []string) {
        for i := 0; i < count; i++ {
            password, err := generator.Generate(length, useDigits, useSymbols)
            if err != nil {
                fmt.Println("Error:", err)
                os.Exit(1)
            }
            fmt.Println(password)
        }
    },
}

// Execute launches the command
func Execute() {
    cobra.OnInitialize()

    // フラグ定義
    rootCmd.Flags().IntVarP(&length, "length", "l", 12, "パスワードの長さ")
    rootCmd.Flags().BoolVarP(&useDigits, "digits", "d", false, "数字を含める")
    rootCmd.Flags().BoolVarP(&useSymbols, "symbols", "s", false, "記号を含める")
    rootCmd.Flags().IntVarP(&count, "count", "c", 1, "生成するパスワードの個数")

    // completion サブコマンドを追加
    var completionCmd = &cobra.Command{
        Use:       "completion [bash|zsh|fish|powershell]",
        Short:     "シェル補完スクリプトを生成します",
        Long:      "1 引数に対象シェルを指定して、補完スクリプトを標準出力に出力します。\n例: gopassgen completion bash > /etc/bash_completion.d/gopassgen",
        Args:      cobra.ExactValidArgs(1),
        ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
        RunE: func(cmd *cobra.Command, args []string) error {
            switch args[0] {
            case "bash":
                return rootCmd.GenBashCompletionFileV2("completions/bash/gopassgen", true)
            case "zsh":
                return rootCmd.GenZshCompletionFile("completions/zsh/_gopassgen")
            case "fish":
                return rootCmd.GenFishCompletionFile("completions/fish/gopassgen.fish", true)
            case "powershell":
                return rootCmd.GenPowerShellCompletionFile("completions/powershell/gopassgen.ps1")
            }
            return nil
        },
    }
    rootCmd.AddCommand(completionCmd)

    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

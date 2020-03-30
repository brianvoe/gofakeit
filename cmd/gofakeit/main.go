package main

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v4/cmd/gofakeit/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofakeit",
	Short: "Random fake data generator",
}

func main() {
	funcs, err := internal.ExportedFuncs(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get exported functions: %v", err)
		os.Exit(1)
	}
	for _, fun := range funcs {
		cmd := funcToCommand(fun)
		rootCmd.AddCommand(cmd)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func funcToCommand(fun *ast.FuncDecl) *cobra.Command {
	name := fun.Name.Name
	commentLines := make([]string, 0)
	if fun.Doc != nil {
		for _, line := range fun.Doc.List {
			commentLines = append(commentLines, line.Text)
		}
	}
	return &cobra.Command{
		Use:   strings.ToLower(name),
		Short: fmt.Sprintf("Generate random %s", name),
		Long:  strings.Join(commentLines, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			result := internal.Execute(name, nil)
			fmt.Fprintln(os.Stdout, result)
		},
	}
}

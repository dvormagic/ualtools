package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dvormagic/ualtools/pkg/config"
)

func init() {
	CmdCache.AddCommand(CmdCachePrint)
}

var CmdCachePrint = &cobra.Command{
	Use:   "print",
	Short: "Print the directory where the artifacts of the tools are cached.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s/.ualtools/cache-%s\n", config.Home(), config.ProjectName())
		return nil
	},
}

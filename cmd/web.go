package cmd

import (
	"fmt"

	"github.com/getchipman/bolt-api/app/context"
	"github.com/spf13/cobra"
)

var apiCMD = &cobra.Command{
	Use:   "web",
	Short: "Start Web Application",
	RunE: func(cmd *cobra.Command, args []string) error {

		// --> New app Context
		ctx := context.New().WithLogger()

		fmt.Println(ctx)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCMD)
}

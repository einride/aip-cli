package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_optimizer_v1_OptimizerService = &cobra.Command{
	Use: "einride.optimizer.v1.Optimizer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.optimizer.v1.Optimizer called")
	},
}

func init() {
	rootCmd.AddCommand(einride_optimizer_v1_OptimizerService)
}

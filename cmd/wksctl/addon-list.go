package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/weaveworks/wksctl/pkg/addons"
)

var addonListCmd = &cobra.Command{
	Use:   "list",
	Short: "List addons",
	Run:   addonListRun,
}

func init() {
	addonCmd.AddCommand(addonListCmd)
}

func addonListRun(cmd *cobra.Command, args []string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, tabWidth, ' ', 0)

	addons := addons.List()
	for _, addon := range addons {
		fmt.Fprintf(w, "%s\t%s\n", addon.ShortName, addon.Name)
	}

	w.Flush()
}

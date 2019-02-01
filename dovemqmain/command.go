package dovemqmain

import (
	"github.com/spf13/cobra"
)


var (
	rootCmd = &cobra.Command{
		Use:        "dovemq",
		Short:      "dovemq server",
		SuggestFor: []string{"dovemq"},
	}
)

func init() {
	rootCmd.AddCommand(newServerCommand())
}


func newServerCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "nodes",
		Short: "Server related command",
	}

	return lpc
}
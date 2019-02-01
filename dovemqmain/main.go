package dovemqmain


import (
	"os"
	"fmt"
)
func Main() {
	if len(os.Args) >1 {
		if err := rootCmd.Execute(); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
}
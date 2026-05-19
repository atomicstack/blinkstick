package flex

import (
	"github.com/spf13/cobra"

	"github.com/atomicstack/blinkstick"
	"github.com/atomicstack/blinkstick/cli/blink/internal"
)

var cmdFlexList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick flex",
	Run: func(cmd *cobra.Command, args []string) {
		b := blinkstick.Flex{}
		internal.DisplayDevices(b.List())
	},
}

package strip

import (
	"github.com/spf13/cobra"

	"github.com/atomicstack/blinkstick"
	"github.com/atomicstack/blinkstick/cli/blink/internal"
)

var cmdStripList = &cobra.Command{
	Use:     "list",
	Short:   "List all blinkstick strip",
	Aliases: []string{"square"},
	Run: func(cmd *cobra.Command, args []string) {
		b := blinkstick.Strip{}
		internal.DisplayDevices(b.List())
	},
}

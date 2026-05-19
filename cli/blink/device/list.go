package device

import (
	"github.com/spf13/cobra"

	"github.com/atomicstack/blinkstick"
	"github.com/atomicstack/blinkstick/cli/blink/internal"
)

func init() {
}

var cmdDeviceList = &cobra.Command{
	Use:   "list",
	Short: "List all blinkstick device: blinkstick device list",
	Run: func(cmd *cobra.Command, args []string) {
		devices := blinkstick.List()
		internal.DisplayDevices(devices)
	},
}

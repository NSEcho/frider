package app

import (
	"errors"
	"github.com/apex/log"
	"github.com/frida/frida-go/frida"
	"github.com/spf13/cobra"
)

var frCmd = &cobra.Command{
	Use:   "fr",
	Short: "Get frontmost application info",
	RunE: func(cmd *cobra.Command, args []string) error {
		dev := frida.USBDevice()
		if dev == nil {
			return errors.New("no USB device detected")
		}

		app, err := dev.FrontmostApplication(frida.ScopeMinimal)
		if err != nil {
			return err
		}
		defer app.Clean()

		log.Infof("Frontmost application: %s (%d)", app.Name(), app.PID())

		return nil
	},
}

func init() {
	AppCmd.AddCommand(frCmd)
}

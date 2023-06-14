package app

import (
	"errors"
	"github.com/apex/log"
	"github.com/frida/frida-go/frida"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		dev := frida.USBDevice()
		if dev == nil {
			return errors.New("no USB device connected")
		}
		apps, err := dev.EnumerateApplications("", frida.ScopeMinimal)
		if err != nil {
			return err
		}
		defer dev.Clean()

		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}

		if all {
			log.Infof("Reading non started applications")
		}

		for _, app := range apps {
			func() {
				defer app.Clean()
				switch app.PID() {
				case 0:
					if all {
						log.WithFields(log.Fields{
							"identifier": app.Identifier(),
							"pid":        app.PID(),
						}).Infof("%s", app.Name())
					}
				default:
					log.WithFields(log.Fields{
						"identifier": app.Identifier(),
						"pid":        app.PID(),
					}).Infof("%s", app.Name())
				}
			}()
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "print only running apps")
	AppCmd.AddCommand(listCmd)
}

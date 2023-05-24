package app

import (
	"fmt"
	"github.com/frida/frida-go/frida"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		dev := frida.USBDevice()
		apps, err := dev.EnumerateApplications("", frida.ScopeMinimal)
		if err != nil {
			return err
		}
		defer dev.Clean()

		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}

		for _, app := range apps {
			func() {
				defer app.Clean()
				switch app.PID() {
				case 0:
					if all {
						fmt.Printf("[*] %d %s (%s)\n", app.PID(), app.Name(), app.Identifier())
					}
				default:
					fmt.Printf("[*] %d %s (%s)\n", app.PID(), app.Name(), app.Identifier())

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

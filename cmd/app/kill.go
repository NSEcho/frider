package app

import (
	"errors"
	"fmt"
	"github.com/frida/frida-go/frida"
	"github.com/spf13/cobra"
	"strings"
)

var killCmd = &cobra.Command{
	Use:   "kill [application Name]",
	Short: "Kill specific process or application",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing application name")
		}

		appName := strings.Join(args, " ")
		dev := frida.USBDevice()
		if dev == nil {
			return errors.New("no USB device connected")
		}

		apps, err := dev.EnumerateApplications("", frida.ScopeMinimal)
		if err != nil {
			return errors.New("error enumerating applications")
		}

		pid := -1
		for _, app := range apps {
			if app.Name() == appName {
				pid = app.PID()
			}
		}

		if pid == -1 {
			return errors.New("no such applications")
		}

		if err := dev.Kill(pid); err != nil {
			return err
		}

		fmt.Printf("[*] Killed %s application with PID %d\n",
			appName, pid)
		return nil
	},
}

func init() {
	AppCmd.AddCommand(killCmd)
}

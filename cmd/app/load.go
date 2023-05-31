package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/apex/log"
	"github.com/frida/frida-go/frida"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var loadCmd = &cobra.Command{
	Use:   "load [script name]",
	Short: "Load script from the database to the application",
	RunE: func(cmd *cobra.Command, args []string) error {
		promptDevice, err := cmd.Flags().GetBool("device")
		if err != nil {
			return err
		}

		var device *frida.Device

		if promptDevice {
			mgr := frida.NewDeviceManager()
			devices, err := mgr.EnumerateDevices()
			if err != nil {
				return err
			}
			var selected int
			var choices []string
			for _, d := range devices {
				choices = append(choices, fmt.Sprintf("[%-6s] %s (%s)", strings.ToUpper(d.DeviceType().String()), d.Name(), d.ID()))
			}
			prompt := &survey.Select{
				Message: "Select what device to connect to:",
				Options: choices,
			}
			if err := survey.AskOne(prompt, &selected); err == terminal.InterruptErr {
				return err
			}
			device = devices[selected]
		} else {
			dev := frida.USBDevice()
			if dev == nil {
				return errors.New("no USB device connected")
			}
			device = dev
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		var scriptContent string

		if file != "" {
			log.Infof("Reading script from %s", file)
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			data, err := io.ReadAll(f)
			if err != nil {
				return err
			}
			scriptContent = string(data)
		} else {

			db, err := database.NewDatabase()
			if err != nil {
				return err
			}
			defer db.Close()

			var scriptName string
			var choices []string

			scripts, err := db.Scripts()
			if err != nil {
				return err
			}

			for _, s := range scripts {
				choices = append(choices, fmt.Sprintf("%s (%s)", s.Name, s.Category))
			}

			var selected int
			prompt := &survey.Select{
				Message: "Select which script to load",
				Options: choices,
			}

			if err := survey.AskOne(prompt, &selected); err == terminal.InterruptErr {
				return err
			}

			scriptName = scripts[selected].Name

			log.Infof("Reading script %s from database", scriptName)

			dbScript, found, err := db.ScriptByName(scriptName)
			if err != nil {
				return err
			}

			if !found {
				return errors.New("no such application")
			}
			scriptContent = string(dbScript.Content)
		}

		appName, err := cmd.Flags().GetString("app")
		if err != nil {
			return err
		}

		defer device.Clean()

		name := appName

		frontmost, err := cmd.Flags().GetBool("running")
		if err != nil {
			return err
		}

		if frontmost {
			frontmostApp, err := device.FrontmostApplication(frida.ScopeMinimal)
			if err != nil {
				return err
			}
			name = frontmostApp.Name()
		}

		log.Infof("Attaching to %s", name)

		session, err := device.Attach(name, nil)
		if err != nil {
			return err
		}
		defer session.Clean()

		log.Infof("Attached to %s", appName)

		script, err := session.CreateScript(scriptContent)
		if err != nil {
			return err
		}
		defer script.Clean()

		script.On("message", func(message string) {
			msg, err := frida.ScriptMessageToMessage(message)
			if err != nil {
				panic(err)
			}

			switch msg.Type {
			case frida.MessageTypeLog:
				log.Infof("SCRIPT LOG: %s", msg.Payload)
			case frida.MessageTypeError:
				log.Errorf("SCRIPT ERROR: %s", msg.Payload)
			}
		})

		if err := script.Load(); err != nil {
			return err
		}

		reader := bufio.NewReader(os.Stdin)
		reader.ReadLine()

		return nil
	},
}

func init() {
	loadCmd.Flags().StringP("script", "s", "", "name of the script from database")
	loadCmd.Flags().StringP("file", "f", "", "path to the script")
	loadCmd.Flags().StringP("app", "a", "", "which application to attach to")
	loadCmd.Flags().BoolP("running", "r", true, "attach to the frontmost application")
	loadCmd.Flags().BoolP("device", "d", false, "interactive prompt for the device")
	AppCmd.AddCommand(loadCmd)
}

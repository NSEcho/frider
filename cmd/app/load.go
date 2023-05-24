package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/frida/frida-go/frida"
	"github.com/nsecho/frider/internal/database"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var loadCmd = &cobra.Command{
	Use:   "load [script name]",
	Short: "Load script from the database to the application",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		scriptName, err := cmd.Flags().GetString("script")
		if err != nil {
			return err
		}

		if file == "" && scriptName == "" {
			return errors.New("you need to pass either file or script from database")
		}

		var scriptContent string

		if file != "" {
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

		dev := frida.USBDevice()
		if dev == nil {
			return errors.New("no usb device detected")
		}

		defer dev.Clean()
		session, err := dev.Attach(appName, nil)
		if err != nil {
			return err
		}
		defer session.Clean()

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
				fmt.Printf("[*] LOG: %s\n", msg.Payload)
			case frida.MessageTypeError:
				fmt.Printf("[*] ERROR: %s\n", msg.Payload)
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
	AppCmd.AddCommand(loadCmd)
}

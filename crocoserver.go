package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "crocoserver"
	app.Usage = "A Docker based app server for non developers"

	app.Commands = []cli.Command{
		{
			Name:  "apps",
			Usage: "List all installed and available apps",
			Action: func(c *cli.Context) error {
				store := NewAppStore("./my.db")
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 0, 8, 1, '\t', 0)

				fmt.Fprintln(w, "APP\tINSTALLED")
				for _, app := range store.Apps() {
					fmt.Fprintf(w, "%s\t%t\n", app.Name, store.IsInstalled(app.Name))
				}
				w.Flush()
				return nil
			},
		},
		{
			Name:  "install",
			Usage: "Install app",
			Action: func(c *cli.Context) error {
				store := NewAppStore("./my.db")
				appName := c.Args().First()
				store.Install(appName)
				return nil
			},
		},
		{
			Name:  "uninstall",
			Usage: "Uninstall app",
			Action: func(c *cli.Context) error {
				store := NewAppStore("./my.db")
				appName := c.Args().First()
				store.Uninstall(appName)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

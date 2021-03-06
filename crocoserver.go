package main

import (
	"fmt"
	"log"
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
			Name:  "gui",
			Usage: "Run a web server with a GUI",
			Action: func(c *cli.Context) error {
				apiContext := ApiContext{store: NewAppStore()}
				err := apiContext.Serve(":8081")
				if err != nil {
					log.Fatalf("Cannot start HTTP server", err)
				}
				return nil
			},
		},
		{
			Name:  "apps",
			Usage: "List all installed and available apps",
			Action: func(c *cli.Context) error {
				store := NewAppStore()
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 0, 8, 1, '\t', 0)

				fmt.Fprintln(w, "APP\tINSTALLED\tLEVEL")
				for _, app := range store.Apps() {
					fmt.Fprintf(w, "%s\t%t\t%s\n", app.Name, store.IsInstalled(app.Name), app.Metadata.Level)
				}
				w.Flush()
				return nil
			},
		},
		{
			Name:  "install",
			Usage: "Install app",
			Action: func(c *cli.Context) error {
				store := NewAppStore()
				appName := c.Args().First()
				store.Install(appName)
				return nil
			},
		},
		{
			Name:  "uninstall",
			Usage: "Uninstall app",
			Action: func(c *cli.Context) error {
				store := NewAppStore()
				appName := c.Args().First()
				store.Uninstall(appName)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

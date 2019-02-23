package main

import (
	"errors"

	"github.com/blueskan/flixer/config"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "flixer [sub]",
		Short: "Flixer automatically creates Web UI for CLI Prompts",
	}

	runCmd := &cobra.Command{
		Use:   "run [strategy]",
		Short: "Hey flixer do all magic for me",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("You must provide output strategy, it can be one of these: stdout|file")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			template, _ := cmd.Flags().GetString("template")
			renderPath, _ := cmd.Flags().GetString("render-path")
			obtainPath, _ := cmd.Flags().GetString("obtain-path")
			url, _ := cmd.Flags().GetString("url")

			config := config.Config{
				Port:     port,
				Template: template,
				Routes: config.RouteDefinitions{
					RenderTemplateRoute: renderPath,
					ObtainInputRoute:    obtainPath,
				},
				Url: url,
			}

			bootstrap(config)
		},
	}

	runCmd.PersistentFlags().String("port", "9000", "HTTP port")
	runCmd.PersistentFlags().String("template", "flixer.html", "Flixer template path")
	runCmd.PersistentFlags().String("render-path", "/", "Rendering path for get")
	runCmd.PersistentFlags().String("obtain-path", "/obtain", "Obtain path for post")
	runCmd.PersistentFlags().String("url", "http://localhost", "Url which open by browser")

	rootCmd.AddCommand(runCmd)
	rootCmd.Execute()
}

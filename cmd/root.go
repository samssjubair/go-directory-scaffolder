package cmd

import (
	"os"

	"directory-scaffolder/internal"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dryRun bool

var rootCmd = &cobra.Command{
	Use:   "scaffold [template.yaml]",
	Short: "Create folder structure from YAML template",
	Long: `A CLI tool that creates folders and files based on a YAML template.

Example:
  scaffold project.yaml
  scaffold --dry-run project.yaml`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tmpl, err := internal.LoadTemplate(args[0])
		if err != nil {
			color.Red("❌ Error loading template: %v", err)
			return err
		}

		if dryRun {
			color.Yellow("🔍 Dry run mode - showing what would be created:")
			internal.PrintStructure(tmpl)
			return nil
		}

		if err := internal.CreateStructure(tmpl); err != nil {
			color.Red("❌ Error creating structure: %v", err)
			return err
		}

		color.Green("✅ Project scaffold created: %s", tmpl.Name)
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show what would be created without actually creating files")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
} 
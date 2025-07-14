package cmd

import (
	"fmt"
	"os"

	"directory-scaffolder/internal"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dryRun bool
var templateName string
var listTemplates bool

var rootCmd = &cobra.Command{
	Use:   "scaffold [template.yaml]",
	Short: "Create folder structure from YAML template",
	Long: `A CLI tool that creates folders and files based on a YAML template.

Examples:
  scaffold project.yaml                    # Use custom YAML template
  scaffold --template react-app            # Use built-in React template
  scaffold --template go-api --name my-api # Use built-in template with custom name
  scaffold --list-templates               # List all built-in templates
  scaffold --dry-run project.yaml         # Preview what would be created`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Handle list templates flag
		if listTemplates {
			return listBuiltinTemplates()
		}

		// Handle template name flag
		if templateName != "" {
			return handleBuiltinTemplate(templateName, args)
		}

		// Handle custom YAML file
		if len(args) == 0 {
			return cmd.Help()
		}

		return handleCustomTemplate(args[0])
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show what would be created without actually creating files")
	rootCmd.Flags().StringVarP(&templateName, "template", "t", "", "Use built-in template (react-app, go-api, node-express)")
	rootCmd.Flags().BoolVarP(&listTemplates, "list-templates", "l", false, "List all available built-in templates")
}

func listBuiltinTemplates() error {
	templates := internal.ListBuiltinTemplates()
	
	color.Cyan("📋 Available built-in templates:")
	fmt.Println()
	
	for _, tmpl := range templates {
		color.Green("  %s", tmpl.Name)
		color.White("    %s", tmpl.Description)
		fmt.Println()
	}
	
	color.Yellow("Usage: scaffold --template <template-name>")
	return nil
}

func handleBuiltinTemplate(templateName string, args []string) error {
	tmpl, err := internal.LoadBuiltinTemplate(templateName)
	if err != nil {
		color.Red("❌ Error loading template: %v", err)
		return err
	}

	// Override template name if provided as argument
	if len(args) > 0 {
		tmpl.Name = args[0]
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
}

func handleCustomTemplate(templatePath string) error {
	tmpl, err := internal.LoadTemplate(templatePath)
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
} 
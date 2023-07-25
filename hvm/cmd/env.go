package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/chanzuckerberg/happy/hvm/installer"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Calculate environment variables for eval() by the calling shell",
	Long: `
Output to STDOUT a list of env vars which should be eval'ed by the calling shell. This is
used to automatically set PATH and other variables via shell hooks.
	`,
	RunE: calcEnvironment,
}

func init() {
	rootCmd.AddCommand(envCmd)

}

// TODO: Split up this function into smaller functions
//
// IMPORTANT: The Stdout of this function is meant to be read by the calling shell.
// Make sure that anything written to Stdout is valid shell code or a comment.
// If you need to make a message to the user, write it to Stderr.
//
// This function is usually called by the shell hook scripts on chpwd.
func calcEnvironment(cmd *cobra.Command, args []string) {

	versionsBase := path.Join(os.UserHomeDirectory(), ".czi", "versions")

	basePath := stripManagedPathsFromPath(versionsBase, os.Getenv("PATH"))
	managedPath := ""

	happyConfig, err := config.GetHappyConfigForCmd(cmd)
	if err != nil {
		// remove managed paths from $PATH
		fmt.Printf("export PATH=%s", basePath)
		return
	} else {
		projectRoot := happyConfig.GetProjectRoot()
		if config.DoesHappyVersionLockFileExist(projectRoot) {
			versionFile, err := config.LoadHappyVersionLockFile(projectRoot)
			if err != nil {
				// remove managed paths from $PATH
				fmt.Printf("export PATH=%s", basePath)
				return
			}

			versionPaths := []string{}
			// iterate lockfile and set $PATH as appropriate
			for k, v := range versionFile.Require {

				org := strings.Split(k, "/")[0]
				project := strings.Split(k, "/")[1]

				// Look for an environment variable named HVM_<PACKAGE> and use the
				// version specified in the env var instead of the one in the lock file.
				// This allows for easier testing.
				override := os.Getenv(fmt.Sprintf("HVM_%s_%s", strings.ToUpper(org), strings.ToUpper(project)))

				if override != "" {
					v = override
				}

				swPath := path.Join(versionsBase, k, v)

				if _, err := os.Stat(swPath); os.IsNotExist(err) {

					org, project := strings.Split(k, "/")[0], strings.Split(k, "/")[1]

					if os.Getenv("HVM_AUTOINSTALL_PACKAGES") == "1" {
						fmt.Fprintf(os.Stderr, "%s version %s is not installed. Downloading it now. Please wait.\n", k, v)
						installer.InstallPackage(org, project, v, runtime.GOOS, runtime.GOARCH, swPath)
					} else {
						fmt.Fprintf(os.Stderr, "Error: %s version %s is not installed. Please run 'hvm install %s'. Set env HVM_AUTOINSTALL_PACKAGES=1 to do this automatically in the future.\n", k, v, v)
					}

				}

				versionPaths = append(versionPaths, swPath)
			}
			managedPath = strings.Join(versionPaths, ":")
		} else {
			fmt.Printf("export PATH=%s", basePath)
			return
		}
	}

	fmt.Printf("export PATH=%s", strings.Join([]string{managedPath, basePath}, ":"))

}

// Return a string of $PATH with all hvm-managed paths removed
func stripManagedPathsFromPath(versionsBasePath, currentPath string) string {

	components := strings.Split(currentPath, ":")
	newComponents := []string{}

	for _, component := range components {
		if !strings.Contains(component, versionsBasePath) {
			newComponents = append(newComponents, component)
		}
	}

	return strings.Join(newComponents, ":")
}

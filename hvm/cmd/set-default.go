package cmd

import (
	"fmt"

	linkmanager "github.com/chanzuckerberg/happy/hvm/linkManager"
	"github.com/spf13/cobra"
)

var setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Symlink the specified version of happy to $HOME/.czi/bin to be used as default",
	Long: `Create a symbolic link $HOME/.czi/bin/ pointing to the specified version of happy. Assuming
$HOME/.czi/bin is set appropriately in your $PATH, this version will be used by default when running 'happy'
outside of a project, or when a happy version config is not present.
	`,
	Run: setDefaultVersion,
}

func init() {
	rootCmd.AddCommand(setDefaultCmd)

}

func setDefaultVersion(cmd *cobra.Command, args []string) {
	versionTag := args[0]

	err := linkmanager.SetBinLink(versionTag)

	if err != nil {
		fmt.Println("Error setting bin link: ", err)
	}
}

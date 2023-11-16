package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bootstrap-cli",
	Short: "A minimalistic CLI to bootstrap projects with different frameworks",
	Long:  GetAsciiArt() + "\nA minimalistic CLI to bootstrap projects with different frameworks.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetAsciiArt() string {

	var (
		colorArr = [5]string{"\u001b[33m", "\u001b[33m", "\u001b[33m", "\u001b[31m", "\u001b[35m"}
		turnoff  = "\u001b[0m\n"
		buf      strings.Builder
		bytes    []byte
		err      error
	)

	if bytes, err = os.ReadFile("asciiArt.txt"); err != nil {
		return "no file"
	}
	strSlice := strings.Split(string(bytes), "\n")
	for i, color := range colorArr {
		buf.WriteString(fmt.Sprintf("%s %s %s", color, strSlice[i], turnoff))
	}
	return buf.String()

}

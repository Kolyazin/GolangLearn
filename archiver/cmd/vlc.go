package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var errEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(errEmptyPath)
	}
	filePath := args[0]
	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	// packed := Encode(data)
	packed := "" + string(data)

	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	// /path/to/file/myfile.txt -> myfile.vlc
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

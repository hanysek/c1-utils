/*
Copyright 2020 Jan Šimůnek github.com/hanysek
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
)

const (
	synopsis      = "Move RAW files without JPG counterpart to Trash folder"
	captureSubdir = "Capture"
	trashSubdir   = "Trash"
)

type CleanRawsCommand struct {
	Ui cli.Ui
}

func (c *CleanRawsCommand) Help() string {
	helpText := fmt.Sprintf(`
Usage: c1-utils clean-raws [options] <session_dir>

	%s

`, synopsis)

	return strings.TrimSpace(helpText)
}

func (c *CleanRawsCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println(c.Help())
		return -1
	}

	// Get session path from command args
	sessionPath := args[len(args)-1]
	if isDirNotExist(sessionPath) {
		fmt.Printf("%s is not valid folder\n", sessionPath)
		return -1

	}

	// Capture session subdir
	capturePath := fmt.Sprintf("%s/%s", sessionPath, captureSubdir)
	if isDirNotExist(capturePath) {
		fmt.Printf("%s is not valid folder\n", capturePath)
		return -1

	}

	// Trash session subdir
	trashPath := fmt.Sprintf("%s/%s", sessionPath, trashSubdir)
	if isDirNotExist(trashPath) {
		fmt.Printf("%s is not valid folder\n", trashPath)
		return -1
	}

	// Get list of JPG files in Trash directory
	trashFiles, err := getListOfFiles(trashPath, ".JPG")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return -1
	}

	// Get list of RAF files in Trash directory
	captureFiles, err := getListOfFiles(capturePath, ".RAF")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return -1
	}

	if len(trashFiles) == 0 || len(captureFiles) == 0 {
		return 0
	}

	// Go for all files in Trash and try to delete counterpart in Capture dir
	for _, trashFile := range trashFiles {
		err := moveCounterpartFile(captureFiles, trashFile, trashPath)
		if err != nil {
			return 1
		}
	}

	return 0
}

func (c *CleanRawsCommand) Synopsis() string {
	return synopsis
}

func isDirNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func getListOfFiles(path string, ext string) ([]string, error) {
	files := []string{}
	walkDir := func(fn string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if !fi.IsDir() && filepath.Ext(fn) == ext {
			files = append(files, fn)
		}

		return nil
	}

	err := filepath.Walk(path, walkDir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func moveCounterpartFile(captureFiles []string, trashFilePath string, trashPath string) error {
	baseFile := filepath.Base(trashFilePath)
	i := findCaptureFile(captureFiles, baseFile)
	if i == -1 {
		return nil
	}

	captureFile := captureFiles[i]

	captureBaseFile := filepath.Base(captureFile)
	destination := trashPath + "/" + captureBaseFile
	fmt.Println("Moving ", captureFile, " to ", destination)
	err := os.Rename(captureFile, destination)

	if err != nil {
		fmt.Printf("Error moving file: %s\n", err)
		return err
	}
	return nil
}

func findCaptureFile(captureFiles []string, trashFile string) int {
	tf := filepath.Base(trashFile)
	for i, captureFile := range captureFiles {
		cf := filepath.Base(captureFile)
		if fileNameWithoutExtension(tf) == fileNameWithoutExtension(cf) {
			return i
		}
	}
	return -1
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

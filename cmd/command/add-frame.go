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
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strings"

	"github.com/mitchellh/cli"
)

const (
	addFrameSynopsis = "Add frame around photo"
	defaultFrameSize = 20
	frameFileSuffix  = " - Frame"
)

type AddFrameCommand struct {
	Ui cli.Ui
}

func (c *AddFrameCommand) Help() string {
	helpText := fmt.Sprintf(`
Usage: c1-utils add-frame [options] <photo>

	%s

Options:

	-size				Frame size in pixels

`, addFrameSynopsis)

	return strings.TrimSpace(helpText)
}

func (c *AddFrameCommand) Run(args []string) int {
	var frameSize int

	cmdFlags := flag.NewFlagSet("add-frame", flag.ExitOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.IntVar(&frameSize, "size", defaultFrameSize, "frame size in pixels")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	if len(args) < 1 {
		c.Ui.Error("Photo file must be specified")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	iPath := args[len(args)-1]
	oPath := addSuffixToFilePath(iPath, frameFileSuffix)

	// Open photo
	img, err := openImage(iPath)
	if err != nil {
		return 1
	}

	// Create subimage
	imgDimension := img.Bounds()
	subImage := img.(SubImager).SubImage(imgDimension.Inset(frameSize))

	// Draw new image with white rectangle and subimage rectangle
	newImage := image.NewRGBA(imgDimension)
	color := color.White
	draw.Draw(newImage, imgDimension, &image.Uniform{color}, image.Point{}, draw.Src)
	draw.Draw(newImage, newImage.Bounds(), subImage, image.Point{}, draw.Src)

	// Sava new photo
	err = saveImage(newImage, oPath)
	if err != nil {
		return 1
	}

	return 0
}

func (c *AddFrameCommand) Synopsis() string {
	return addFrameSynopsis
}

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
	"image/jpeg"
	"os"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
)

const (
	addFrameSynopsis = "Add frame around photo"
	defaultFrameSize = 20
	suffix           = " - Frame"
	quality          = 85
)

type AddFrameCommand struct {
	Ui cli.Ui
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
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
		c.Ui.Error("Image file must be specified")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	iPath := args[len(args)-1]
	oPath := addSuffixToFilePath(iPath, suffix)

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
	return synopsis
}

func openImage(path string) (image.Image, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening photo")
	}

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		return nil, errors.Wrap(err, "Error decoding photo. Is it a real JPEG?")
	}

	return img, nil
}

func saveImage(img image.Image, fpath string) error {
	f, err := os.Create(fpath)
	if err != nil {
		return errors.Wrap(err, "Cannot create file: "+fpath)
	}
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return errors.Wrap(err, "Failed to encode the image as JPEG")
	}
	return nil
}

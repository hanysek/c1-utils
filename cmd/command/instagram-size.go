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
	instagramSizeSynopsis = "Adjust size for Instagram - makes photos size to 1350x1350px"
	instagramSuffix       = " - InstSize"
)

type InstagramSizeCommand struct {
	Ui cli.Ui
}

func (c *InstagramSizeCommand) Help() string {
	helpText := fmt.Sprintf(`
Usage: c1-utils instagram-size [options] <photo>

	%s

Options:

`, addFrameSynopsis)

	return strings.TrimSpace(helpText)
}

func (c *InstagramSizeCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("instagram-size", flag.ExitOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
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
	oPath := addSuffixToFilePath(iPath, instagramSuffix)

	// Open photo
	img, err := openImage(iPath)
	if err != nil {
		return 1
	}

	// This util does not resize image. Max size must be 1350x1350px
	// TODO: Implement image resize
	if img.Bounds().Dx() > instagramDimension || img.Bounds().Dy() > instagramDimension {
		c.Ui.Error("Photo must be max 1350x1350px")
		return 1
	}

	// Center photo
	x := (instagramDimension - img.Bounds().Dx()) / 2
	y := (instagramDimension - img.Bounds().Dy()) / 2
	newDimension := image.Rect(x, y, img.Bounds().Dx()+x, img.Bounds().Dy()+y)

	// Draw new image with white rectangle and original image
	newImage := image.NewRGBA(image.Rect(0, 0, instagramDimension, instagramDimension))
	color := color.White
	draw.Draw(newImage, newImage.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
	draw.Draw(newImage, newDimension, img, image.Point{}, draw.Src)

	// Sava new photo
	err = saveImage(newImage, oPath)
	if err != nil {
		return 1
	}

	return 0
}

func (c *InstagramSizeCommand) Synopsis() string {
	return instagramSizeSynopsis
}

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
	"image"
	"image/jpeg"
	"os"

	"github.com/pkg/errors"
)

const (
	quality            = 85
	instagramDimension = 1350
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
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

Capture One Utils
=================

Utils for photos management if you use Capture One for FUJI photos. It is written in Golang.

## Usage

### Clean RAF files

You can use this if you shoot in RAW + JPEG and delete in C1 JPEG only.

Command searching for JPEG files in Trash folder and moves RAF files from Capture folder to Trash folder as well.

Run:

	c1-utils clean-raws <capture_dir> <trash-dir>

### Add frame

Adds frame around photo

Run:

	c1-utils add-frame -size 100 <photo>

### Instagram size

Adjusts size for Instagram - makes photos size to 1350x1350px.

Run:

	c1-utils instagram-size <photo>

## Build

Run:

	make build

## Install into GOPATH

	make install

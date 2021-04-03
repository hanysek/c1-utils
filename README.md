Capture One Utils
=================

Utils for photos management if you use Capture One for FUJI photos. It is written in Golang.

## Usage

### Clean RAF files

You can use this if you shoot in RAW + JPEG and delete JPEG pair only.

This simple utility search for RAF files which have JPG files in Trash folder.

Run:

	c1-utils clean-raws <capture_dir> <trash-dir>

Command searching for JPEG files in Trash folder and removes RAF from Capture folder. 

## Build

Run:

	make build

## Install into GOPATH

	make install

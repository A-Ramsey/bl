# Better ls (bl)

[![Go](https://github.com/A-Ramsey/bl/actions/workflows/go.yml/badge.svg)](https://github.com/A-Ramsey/bl/actions/workflows/go.yml)

This is a command written in go designed to improve upon ls by adding better options for the output

## How to install
Assuming you have go already setup, clone this repo cd into it and run `go install .`

## Features:
- Tree based output, outputs the file tree by default with a depth of 3, ignoring .git and node_modules folders. Control the depth it outputs by using the `-d` flag
  - e.g. `bl -d 5` to show 5 levels deep or `bl -d 1` will act like a regular ls and just do the current folder

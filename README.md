# Kindle Clippings Manager

Clippings is a CLI Tool written in Go with the Cobra framework. It aims to easy the management of notes, highlights and bookmarks made in a e-books with Kindle.

To achieve this, the tool parses the `My Clippings.txt` file, which resides in the Kindle's filesystem, to a JSON structure. Then, you can query and update your JSON database through the command line.

### Note
This tool was only tested against a 6th Generation Kindle Touch. If the `My Clippings.txt` pattern changes in other version please feel free to open a issue to request support or a pull request to contribute.

# Installation

First, get the source code with the following command:

```shell
$ git clone git@github.com:caiorcferreira/clippings-cli.git
```
 
 The CLI uses Go Modules to manage its dependencies. As such, you can place this project anywhere in your machine.
 
 Then, install the dependencies and vendor them with: `$ go mod vendor`
 
 Third, build the tool: `go build`
 
 Finally, copy the created binary to a local in your `PATH`.
 
 # Contributing
 
 Please feel free to review the code and open request to fixes or new functionality.
 
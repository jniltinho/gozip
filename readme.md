# gozip

Create a self extracting executable by calling the `Unzip` func on its own binary.

## Install

The `gozip` command can be used to add zipped content to your binary.

```bash
go get github.com/jniltinho/gozip
```

## Usage

```
Usage of gozip:
package main
import ("github.com/jniltinho/gozip")

gozip.Unzip("filename.zip", "filename_folder")
```



# GOSCII

library and cli for convert image to ASCII art

## Installation

```bash
# package
go get github.com/Perkovec/goscii

# cli
go install github.com/Perkovec/goscii/cli/goscii
```

## CLI usage

```bash
goscii [args]
```

| Argument | Default | Description |
| -------- | ------- | ----------- |
| -h       |         | help |
| -i       |         | input image file |
| -c       | 0       | output columns (0 - terminal width) |
| -r       | 0       | output rows (0 - terminal height) |
| -f       | contain | fit algorithm (contain, cover, height, width, fill) |

## Package methods

### NewConverter

Create new converter instance with options

```go
package main

import (
    "fmt"
    "github.com/Perkovec/goscii"
)

func main() {
    converter, err := goscii.NewConverter()
}
```

### Convert

Convert image to ASCII art

```go
package main

import (
    "fmt"
    "github.com/Perkovec/goscii"
    "image"
    "os"
    _ "image/jpeg"
)

func main() {
    converter, err := goscii.NewConverter()
    if err != nil {
        panic(err)
    }

    imageFile, err := getImageFromFilePath("image.jpg")
    if err != nil {
        panic(err)
    }

    artRows := converter.Convert(imageFile)
   
    for _, row := range artRows {
        fmt.Println(row)
    }
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
```
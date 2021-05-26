// Package funks implements basic functions that I use frequently to save time on replicating code across different projects.
package funks

import (
	"fmt"
	"image"
	"os"
)

func LoadImage(f string) (image.Image, error) {
    file, err := os.Open(f)
    if err != nil {
        return nil, fmt.Errorf("unable to open image file %s: %w", f, err)
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return nil, fmt.Errorf("unable to decode image file %s: %w", f, err)
    }

    return img, nil
}

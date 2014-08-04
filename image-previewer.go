package main

import (
	"github.com/disintegration/imaging"
	"runtime"
	"image"
	"image/color"
	"fmt"
	"flag"
	"path/filepath"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	smallSize := flag.Int("small-size", 250, "small size")
	bigSize := flag.Int("big-size", 1200, "big size")
	flag.Parse()
	originalFiles := flag.Args()

	saveThumbs(createThumbs(originalFiles, *smallSize), "small")
	saveThumbs(createThumbs(originalFiles, *bigSize), "big")
}

func saveThumbs(thumbs map[string]image.Image, suffix string) {
	for fileName,thumb := range(thumbs) {
		dst := imaging.New(thumb.Bounds().Dx(), thumb.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
		extension := filepath.Ext(fileName)
		name := fileName[0:len(fileName)-len(extension)]
		err := imaging.Save(dst, fmt.Sprintf("%s_%s.jpg", name, suffix))
		if (err != nil) {
			panic(err)
		}
	}
}

func createThumbs(originalFiles []string, size int) (map[string]image.Image) {
	thumbs := make(map[string]image.Image)
	for _,file := range originalFiles {
		img, err := imaging.Open(file)
		if err != nil {
			panic(err)
		}
		thumb := imaging.Fit(img, size, size, imaging.Lanczos)
		thumbs[file] = thumb
	}

	return thumbs
}

package day20

import (
	"image"
	"image/color"
	"image/gif"
	"os"
	"strconv"
	"strings"
)

type pixel struct {
	X, Y int
}

type egami struct {
	raw           []string
	height, width int
	pixels        map[pixel]bool
	litPixels     int
	infIsLit      bool
}

func Puzzle(input *[]byte, enhanceCount int) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n\n")
	rawImage := strings.Split(strings.TrimSpace(in[1]), "\n")
	enhanceAlgo := in[0]
	pic := egami{
		raw:       rawImage,
		height:    len(rawImage),
		width:     len(rawImage[0]),
		pixels:    map[pixel]bool{},
		litPixels: 0,
		infIsLit:  false,
	}
	images := []*image.Paletted{}
	delays := []int{}
	for i := 0; i < enhanceCount; i++ {
		pic.enhance(&enhanceAlgo)
		images = append(images, pic.GetImage(enhanceCount-i))
		delays = append(delays, 25)
	}
	file, err := os.OpenFile("2021/go/day20/20.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	gif.EncodeAll(file, &gif.GIF{Image: images, Delay: delays})
	return pic.litPixels
}

func (i *egami) GetImage(padding int) *image.Paletted {
	palette := []color.Color{
		color.Black,
		color.White,
		color.Transparent,
	}
	w, h := i.width+padding*2, i.height+padding*2
	img := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{w, h}}, palette)
	for y := 0; y < h; y++ {
		for x := 0; x < h; x++ {
			img.Set(x, y, color.Transparent)
			if lit, ok := i.pixels[pixel{x - padding, y - padding}]; ok {
				if lit {
					img.Set(x, y, color.Black)
				} else {
					img.Set(x, y, color.White)
				}
			}
		}
	}
	return img
}

func (i *egami) enhance(enhanceAlgo *string) {
	i.expand()
	i.process()
	newImage := make(map[pixel]bool)
	lit := 0
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			current := pixel{X: x, Y: y}
			binaryIndex := ""
			indexKeys := [9]pixel{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}}
			for _, n := range indexKeys {
				bit := "0"
				if lit, ok := i.pixels[pixel{X: current.X + n.X, Y: current.Y + n.Y}]; ok {
					if lit {
						bit = "1"
					}
				} else if i.infIsLit {
					bit = "1"
				}
				binaryIndex += bit
			}
			index, _ := strconv.ParseInt(binaryIndex, 2, 0)
			if (*enhanceAlgo)[index] == '#' {
				newImage[current] = true
				lit++
			} else {
				newImage[current] = false
			}
		}
	}
	i.pixels = newImage
	i.rawImage()
	i.litPixels = lit
	if (*enhanceAlgo)[0] == '#' {
		if i.infIsLit {
			i.infIsLit = false
		} else {
			i.infIsLit = true
		}
	}
}

func (i *egami) expand() {
	i.height, i.width = len(i.raw)+2, len((i.raw)[0])+2
	yPadding, padding := "", "."
	if i.infIsLit {
		padding = "#"
	}
	for n := 0; n < i.width; n++ {
		yPadding += padding
	}
	expandedRawImage := []string{}
	expandedRawImage = append(expandedRawImage, yPadding)
	for _, line := range i.raw {
		expandedRawImage = append(expandedRawImage, padding+line+padding)
	}
	expandedRawImage = append(expandedRawImage, yPadding)
	i.raw = expandedRawImage
}

func (i *egami) process() {
	pixels := make(map[pixel]bool)
	for y, line := range i.raw {
		for x, p := range line {
			if p == '#' {
				pixels[pixel{X: x, Y: y}] = true
			} else {
				pixels[pixel{X: x, Y: y}] = false
			}
		}
	}
	i.pixels = pixels
}

func (i *egami) rawImage() {
	rawImage := []string{}
	for y := 0; y < i.height; y++ {
		line := ""
		for x := 0; x < i.width; x++ {
			p, r := pixel{X: x, Y: y}, "."
			if i.pixels[p] {
				r = "#"
			}
			line += r
		}
		rawImage = append(rawImage, line)
	}
	i.raw = rawImage
}

// Lissajous generates GIF animations of random Lissajous figures.
// Link: https://docs.hacknode.org/gopl-zh/ch1/ch1-04.html
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type txt2file struct {
	filename string
	content []byte
}

func (tf *txt2file) Write(p []byte) (n int, err error)  {
	fmt.Println("txt2file")
	//fmt.Println(string(p))
	if tf.content == nil{
		tf.content = []byte{}
	}
	tf.content = append(tf.content, p...)
	return 0, nil
}

// 写入文件
func (tf *txt2file) output()  {
	ioutil.WriteFile(tf.filename, tf.content, 0777)
}

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	lissajous(os.Stdout)

	// 将内容输入到文本
	tf := &txt2file{filename:"./lissajous.gif"}
	lissajous(tf)
	tf.output()
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
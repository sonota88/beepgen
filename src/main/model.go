package model

import (
	"fmt"
	"math"
	"os"
)

type config struct {
	vol   float64
	srate int
}

func Main(srate int, freq float64, dur int, vol float64, inst string) {
	conf := config{
		vol:   vol,
		srate: srate}

	tmax := conf.srate * dur / 1000
	for t := 0; t < tmax; t++ {
		bs := []uint8{calcAmp(conf, freq, t, inst)}
		os.Stdout.Write(bs)
	}
}

func fmod(x float64, y float64) float64 {
	div := x / y
	divInt := math.Trunc(div)
	return x - (y * divInt)
}

func toCycleRatio(t int, srate int, freq float64) float64 {
	samplesPerCycle := float64(srate) / freq
	mod := fmod(float64(t), samplesPerCycle)
	return mod / samplesPerCycle
}

func calcAmp(conf config, freq float64, t int, inst string) uint8 {
	cycleRatio := toCycleRatio(t, conf.srate, freq)

	var signedVal float64

	if inst == "sq" {
		signedVal = oscSq(cycleRatio)
	} else if inst == "tri" {
		signedVal = oscTri(cycleRatio)
	} else {
		fmt.Fprintf(os.Stderr, "Must not happen")
		os.Exit(1)
	}

	return uint8(signedVal*conf.vol + 128)
}

// @return -128.0 <= retval <= 127.0
func oscSq(r float64) float64 {
	if r < 0.5 {
		return 127.0
	} else {
		return -128.0
	}
}

// @return -128.0 <= retval <= 127.0
func oscTri(r float64) float64 {
	if r < 0.25 {
		return 512.0*r + 0.0
	} else if r < 0.75 {
		return -512.0*r + 255.0
	} else {
		return 512.0*r - 512.0
	}
}

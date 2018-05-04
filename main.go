package main

import (
	"./src/main"
	"./src/main/maparse"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Fprintf(os.Stderr, "")

	opts, err := maparse.ParseArgs(os.Args[1:], []string{})
	if err != nil {
		panic(err)
	}

	// Hz
	freq := getFrequency(opts)

	// msec
	dur := getDuration(opts)

	vol := getVolume(opts)

	inst := getInstrument(opts)

	srate := getSamplingRate(opts)

	model.Main(srate, freq, dur, vol, inst)
}

func getFrequency(opts map[string]string) float64 {
	freq := getFloatOrDefault(opts, "f", 750.0)

	if freq > 0 {
		// OK
	} else {
		fmt.Fprintf(os.Stderr, "Invalid argument: 0 < frequency\n")
		os.Exit(1)
	}

	return freq
}

func getSamplingRate(opts map[string]string) int {
	return getIntOrDefault(opts, "sr", 44100)
}

func getDuration(opts map[string]string) int {
	dur := getIntOrDefault(opts, "l", 500)

	if dur == 0 {
		fmt.Fprintf(os.Stderr, "Invalid argument: 0 < duration\n")
		os.Exit(1)
	}

	return dur
}

func getVolume(opts map[string]string) float64 {
	vol := getIntOrDefault(opts, "v", 5)

	if vol < 0 || vol > 100 {
		fmt.Fprintf(os.Stderr, "Invalid argument: 0 <= volume <= 100\n")
		os.Exit(1)
	}

	return float64(vol) / 100.0
}

func getInstrument(opts map[string]string) string {
	inst := getOrDefault(opts, "i", "sq")

	if inst == "sq" || inst == "tri" {
		// OK
	} else {
		fmt.Fprintf(os.Stderr, "Invalid argument: instrument must be sq or tri\n")
		os.Exit(1)
	}

	return inst
}

func getOrDefault(opts map[string]string, key string, defaultVal string) string {
	v, hasKey := opts[key]
	if hasKey {
		return v
	} else {
		return defaultVal
	}
}

func getIntOrDefault(opts map[string]string, key string, defaultVal int) int {
	str, hasKey := opts[key]

	if !hasKey {
		return defaultVal
	}

	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}

func getFloatOrDefault(opts map[string]string, key string, defaultVal float64) float64 {
	str, hasKey := opts[key]

	if !hasKey {
		return defaultVal
	}

	x, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return x
}

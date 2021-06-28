package main

import (
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type display struct {
	width  int
	height int
	xStart int
}

func newDisplay(xrandrStr string) display {
	regex := regexp.MustCompile("[0-9]{1,4}x[0-9]{1,4}\\+[0-9]{1,4}\\+[0-9]{1,4}")
	nums := regex.FindString(xrandrStr)
	parts := strings.SplitN(nums, "+", 2)
	widthHeight := strings.Split(parts[0], "x")
	offsets := strings.Split(parts[1], "+")
	w, _ := strconv.Atoi(widthHeight[0])
	h, _ := strconv.Atoi(widthHeight[1])
	leftOffset, _ := strconv.Atoi(offsets[0])

	return display{
		width:  w,
		height: h,
		xStart: leftOffset,
	}
}

type displayConfig struct {
	displays []display
}

func newDisplayConfig() displayConfig {
	cmd := exec.Command("xrandr", "--query")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile("connected .*[0-9]{1,4}x[0-9]{1,4}\\+[0-9]{1,4}\\+[0-9]{1,4}")
	dispsStrs := regex.FindAllString(string(stdout), -1)
	var disps []display

	for _, disStr := range dispsStrs {
		disps = append(disps, newDisplay(disStr))
	}

	sort.Slice(disps, func(i, j int) bool {
		return disps[i].xStart < disps[j].xStart
	})

	return displayConfig{
		displays: disps,
	}
}

func (conf displayConfig) getDisplay(geo windowGeometry) display {
	if len(conf.displays) == 1 {
		return conf.displays[0]
	}

	midX := geo.x + (geo.width / 2)

	for i := len(conf.displays) - 1; i >= 0; i-- {
		if conf.displays[i].xStart+10 <= midX {
			return conf.displays[i]
		}
	}

	return conf.displays[0]
}

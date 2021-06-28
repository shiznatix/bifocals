package main

import (
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type anchorSide int
type windowGeometry struct {
	x      int
	y      int
	width  int
	height int
}

const (
	LEFT = iota
	RIGHT
	TOP
	BOTTOM
	FULLSCREEN
)

func getAnchorSide() anchorSide {
	if len(os.Args) < 2 {
		return LEFT
	}

	switch os.Args[1] {
	case "right":
		return RIGHT
	case "top":
		return TOP
	case "bottom":
		return BOTTOM
	case "fullscreen":
		return FULLSCREEN
	default:
		return LEFT
	}
}

func xdotool(args ...string) string {
	cmd := exec.Command("xdotool", args...)
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return strings.Trim(string(stdout), "\n")
}

func getWindowGeometry(windowId string) windowGeometry {
	geo := xdotool("getwindowgeometry", windowId)
	lines := strings.Split(geo, "\n")
	posRegex := regexp.MustCompile("[0-9]+,[0-9]+")
	pos := posRegex.Find([]byte(lines[1]))
	xy := strings.Split(string(pos), ",")
	xInt, err := strconv.Atoi(xy[0])
	if err != nil {
		panic(err)
	}
	yInt, err := strconv.Atoi(xy[1])
	if err != nil {
		panic(err)
	}

	sizeRegex := regexp.MustCompile("[0-9]+x[0-9]+")
	dimens := sizeRegex.Find([]byte(lines[2]))
	wh := strings.Split(string(dimens), "x")
	widthInt, err := strconv.Atoi(wh[0])
	if err != nil {
		panic(err)
	}
	heightInt, err := strconv.Atoi(wh[1])
	if err != nil {
		panic(err)
	}

	return windowGeometry{
		x:      xInt,
		y:      yInt,
		width:  widthInt,
		height: heightInt,
	}
}

func getResizeWidth(geo windowGeometry, disp display) int {
	size1 := disp.width / 3
	size2 := disp.width / 2
	size3 := size1 * 2
	c := 10

	if geo.width < size1-c {
		return size1
	}
	if geo.width < size2-c {
		return size2
	}
	if geo.width < size3-c {
		return size3
	}

	return size1
}

func getResizeHeight(geo windowGeometry, disp display) int {
	size1 := disp.height / 3
	size2 := disp.height / 2
	size3 := size1 * 2
	c := 20

	if geo.height < size1-c {
		return size1
	}
	if geo.height < size2-c {
		return size2
	}
	if geo.height < size3-c {
		return size3
	}

	return size1
}

func main() {
	conf := newDisplayConfig()
	toSide := getAnchorSide()
	windowId := xdotool("getactivewindow")
	geo := getWindowGeometry(windowId)
	disp := conf.getDisplay(geo)

	width := 0
	height := 0
	xStart := 0
	yStart := 0

	switch toSide {
	case RIGHT:
		width = getResizeWidth(geo, disp)
		height = disp.height - 50
		xStart = disp.xStart + disp.width - width + 10
		yStart = 0
	case LEFT:
		width = getResizeWidth(geo, disp)
		height = disp.height - 50
		xStart = disp.xStart
		yStart = 0
	case TOP:
		width = disp.width
		height = getResizeHeight(geo, disp)
		xStart = disp.xStart
		yStart = 0
	case BOTTOM:
		width = disp.width
		height = getResizeHeight(geo, disp)
		xStart = disp.xStart
		yStart = disp.height - height
	case FULLSCREEN:
		width = disp.width
		height = disp.height - 50
		xStart = disp.xStart
		yStart = 0
	default:
		panic("Unknown side")
	}

	//fmt.Println("windowId:", windowId)
	//fmt.Println("disp h:", disp.height, "w:", disp.width, "xStart:", disp.xStart)
	//fmt.Println("toSide:", toSide)
	//fmt.Println("geo.x:", geo.x, "geo.y:", geo.y, "geo.width", geo.width, "geo.height", geo.height)
	//fmt.Println("height:", height)
	//fmt.Println("width:", width)
	//fmt.Println("xStart:", xStart)
	//fmt.Println("yStart:", yStart)

	xdotool("windowsize", windowId, strconv.Itoa(width), strconv.Itoa(height))
	xdotool("windowmove", windowId, strconv.Itoa(xStart), strconv.Itoa(yStart))
}

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

func main() {
	display := microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	buttonLeft := machine.BUTTONA
	buttonLeft.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	buttonRight := machine.BUTTONB
	buttonRight.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	display.ClearDisplay()
	display.SetPixel(0, 0, microbitmatrix.Brightness9)

	// Track previous button states for edge detection
	prevLeft := true
	prevRight := true

	then := time.Now()
	for {
		if time.Since(then).Milliseconds() > 50 {
			then = time.Now()

			currentLeft := buttonLeft.Get()
			currentRight := buttonRight.Get()

			// Detect button press (transition from true to false)
			if prevLeft && !currentLeft {
				// Button A pressed - turn all on
				setAllPixels(display, microbitmatrix.BrightnessFull)
			}

			if prevRight && !currentRight {
				// Button B pressed - turn all off
				setAllPixels(display, microbitmatrix.BrightnessOff)
			}

			prevLeft = currentLeft
			prevRight = currentRight
		}

		display.Display()
	}
}

func setAllPixels(display microbitmatrix.Device, brightness color.RGBA) {
	for row := int16(0); row < 5; row++ {
		for col := int16(0); col < 5; col++ {
			display.SetPixel(col, row, brightness)
		}
	}
}

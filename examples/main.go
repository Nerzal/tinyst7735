package main

import (
	"machine"
	"time"

	"github.com/Nerzal/tinyst7735/color"

	"github.com/Nerzal/tinyst7735"
)

func main() {
	println("starting")
	// config.SCK = PB5
	// config.SDO = PB3
	// config.SDI = PB4
	machine.SPI0.Configure(machine.SPIConfig{Frequency: 127500, SCK: machine.D5, SDO: machine.D3, SDI: machine.D4})
	println("spi configured")
	time.Sleep(200 * time.Millisecond)

	display := tinyst7735.New(machine.SPI0, machine.D6, machine.D7, machine.D8, machine.D9)
	display.Configure(tinyst7735.Config{})

	println("display configured")
	display.EnableBacklight(true)

	// for a := byte('!'); a < 127; a++ {
	// 	print("input: ", a, " ")
	// 	b, _ := machine.SPI0.Transfer(a)
	// 	println("data read: ", b)
	// 	time.Sleep(300 * time.Millisecond)
	// }

	width, height := display.Size()

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	for {

		println("start to fill screen")

		display.FillScreen(black)
		time.Sleep(4 * time.Second)

		println("start to draw rectangles")

		display.FillRectangle(0, 0, width/2, height/2, white)
		display.FillRectangle(width/2, 0, width/2, height/2, red)
		display.FillRectangle(0, height/2, width/2, height/2, green)
		display.FillRectangle(width/2, height/2, width/2, height/2, blue)
		display.FillRectangle(width/4, height/4, width/2, height/2, black)

		time.Sleep(2 * time.Second)
	}
}

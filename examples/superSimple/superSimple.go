package main

import (
	"github.com/seemywingz/gg"
)

func main() {

	gg.Init(800, 600, "Simple Triangle")

	tri := gg.NewPointsObject(
		gg.NewPosition(0, 0, -5),
		gg.Triangle,
		0,
		gg.Shader["basic"],
	)

	for !gg.ShouldClose() {
		gg.Update()
		tri.Draw()
		gg.SwapBuffers()
	}
}

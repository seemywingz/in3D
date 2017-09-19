package main

// Card :
type Card struct {
	Front *DrawnObjectData
	Back  *DrawnObjectData
}

// New : Create New Card
func (Card) New(p Position) *Card {

	cardfront := DrawnObjectData{}.New(p, cardFront, texture["tk"], shader["phong"])
	cardback := DrawnObjectData{}.New(p, cardBack, texture["back"], shader["phong"])
	cl := func(d *DrawnObjectData) {
		// d.XRotation++
		d.YRotation++
	}
	cardfront.DrawLogic = cl
	cardback.DrawLogic = cl
	return &Card{cardfront, cardback}
}

// Draw : draw the card (makes Card a DrawnObject Interface)
func (c *Card) Draw() {
	c.Front.Draw()
	c.Back.Draw()
}

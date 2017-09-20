package gg

// Card :
type Card struct {
	Front *DrawnObjectData
	Back  *DrawnObjectData
}

// NewCard : Create New Card
func NewCard(p Position, ftexture, btexture, shader uint32) *Card {

	cardfront := NewDrawnObject(p, cardFront, ftexture, shader)
	cardback := NewDrawnObject(p, cardBack, btexture, shader)
	cl := func(d *DrawnObjectData) {
		// d.XRotation++
		d.YRotation += 0.5
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

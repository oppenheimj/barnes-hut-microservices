package internal

import (
	"math"
)

type body struct {
	position location
	velocity location
	force    location
	mass     float64
	netForce float64
}

func (b *body) addForce(otherBody *body) {
	G := 6.67e-2 // gravitational constant e-11
	dx := otherBody.position.x - b.position.x
	dy := otherBody.position.y - b.position.y
	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	F := (G * b.mass * otherBody.mass) / math.Pow(distance, 2)

	b.force.x += F * dx / distance
	b.force.y += F * dy / distance

	// fmt.Println(b.mass, otherBody.mass, b.force)
}

func (b *body) applyForce() {
	b.velocity.x += b.force.x / b.mass
	b.velocity.y += b.force.y / b.mass

	b.position.x += b.velocity.x
	b.position.y += b.velocity.y

	b.force = location{}
}

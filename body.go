package main

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

func (b *body) addForce(otherBody body) {
	G := 6.67e-11 // gravitational constant
	dx := otherBody.position.x - b.position.x
	dy := otherBody.position.y - b.position.y
	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	F := (G * b.mass * otherBody.mass) / math.Pow(distance, 2)
	b.force.x += F * dx / distance
	b.force.y += F * dy / distance
}

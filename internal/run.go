package internal

import (
	"math/rand"
	"time"
)

var width = 800.0
var n = 2000
var ticks = 500
var g = 6.67e-6 // e-11
var theta = 0.5

// Run the simulation
func Run() {
	bodies := generateRandomBodies(n)
	data := [][]location{}

	for t := 0; t < ticks; t++ {
		data = append(data, extractLocations(bodies))
		root := node{location: location{x: 0, y: 0}, width: width}
		for i := range bodies {
			if root.contains(&bodies[i]) {
				root.addBody(&bodies[i])
			}
		}

		root.calculateCentersOfMass()

		for i := range bodies {
			if root.contains(&bodies[i]) {
				root.calculateForceOnBody(&bodies[i])
				bodies[i].applyForce()
			}
		}
	}

	generateGif(data)
}

func extractLocations(bodies []body) []location {
	locations := []location{}
	for i := range bodies {
		locations = append(locations, bodies[i].position)
	}

	return locations
}

func generateRandomBodies(n int) []body {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomLocation := func() location {
		return location{
			x: r.Float64() * width,
			y: r.Float64() * width,
		}
	}

	points := []body{}

	for i := 0; i < n; i++ {
		points = append(points, body{position: randomLocation(), mass: r.Float64() * 10000})
	}

	return points
}

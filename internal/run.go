package internal

import (
	"math/rand"
	"time"
	// "sync"
)

// Run the simulation
func Run() {
	bodies := generateRandomBodies(500)
	data := [][]location{}

	for t := 0; t < 1000; t++ {
		data = append(data, extractLocations(bodies))
		root := node{location: location{x: 0, y: 0}, width: 800.0}
		for i := range bodies {
			if root.contains(&bodies[i]) {
				root.addBody(&bodies[i])
			}
		}

		for i := range bodies {
			root.calculateForceOnBody(&bodies[i])
			bodies[i].applyForce()
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
			x: r.Float64() * 800,
			y: r.Float64() * 800,
		}
	}

	points := []body{}

	for i := 0; i < n; i++ {
		points = append(points, body{position: randomLocation(), mass: r.Float64() * 100000})
	}

	return points
}

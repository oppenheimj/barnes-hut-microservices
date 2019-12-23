package main

import (
	"math"
)

type node struct {
	children     []node
	centerOfMass location
	totalMass    float64
	width        float64
	location     location
	body         body
}

var theta = 0.5

func (n *node) addBody(b body) {
	recursiveCall := func() {
		i := n.locationToChildrenIndex(b.position)
		n.children[i].addBody(b)
	}

	if n.isLeaf() {
		if n.hasBody() {
			n.convertToInternal()
			recursiveCall()
		} else {
			n.body = b
		}
	} else {
		n.recalculateCenterOfMass(b)
		recursiveCall()
	}
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) hasBody() bool {
	return n.body != (body{})
}

func (n *node) convertToInternal() {
	childLocations := childLocations(n.location, n.width)

	for i := 0; i < 4; i++ {
		n.children = append(n.children, node{location: childLocations[i], width: n.width / 2})
	}

	if n.hasBody() {
		pTmp := n.body
		n.body = body{}
		n.addBody(pTmp)
	}
}

func (n *node) recalculateCenterOfMass(b body) {
	n.totalMass += b.mass
	n.centerOfMass = location{
		x: (n.centerOfMass.x*n.totalMass + b.position.x*b.mass) / n.totalMass,
		y: (n.centerOfMass.y*n.totalMass + b.position.y*b.mass) / n.totalMass,
	}
}

// [nw, ne, sw, se]
func childLocations(l location, w float64) []location {
	return []location{
		location{x: l.x - w/4, y: l.y + w/4},
		location{x: l.x + w/4, y: l.y + w/4},
		location{x: l.x - w/4, y: l.y - w/4},
		location{x: l.x + w/4, y: l.y - w/4},
	}
}

func (n *node) locationToChildrenIndex(l location) int {
	if l.x < n.location.x {
		if l.y < n.location.y {
			return 2
		}
		return 0
	}

	if l.y < n.location.y {
		return 3
	}
	return 1
}

func (n *node) calculateForceOnBody(b body) {
	if n.body != b && (body{}) != b {
		if n.hasBody() {
			b.addForce(n.body)
		} else {
			threshold := n.width / math.Sqrt(math.Pow(n.centerOfMass.x-b.position.x, 2)+math.Pow(n.centerOfMass.y-b.position.y, 2))
			if threshold < theta {
				b.addForce(body{position: n.centerOfMass, mass: n.totalMass})
			} else {
				for i := range n.children {
					n.children[i].calculateForceOnBody(b)
				}
			}
		}
	}
}

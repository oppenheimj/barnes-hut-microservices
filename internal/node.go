package internal

import (
	"math"
)

type node struct {
	children     []node
	centerOfMass location
	totalMass    float64
	width        float64
	location     location
	body         *body
}

var solarMass = 2e30 // 1 - 10 for stars, 1000000x for black hole
var theta = 0.5

func (n *node) addBody(b *body) {
	// fmt.Println("In addBody() with", len(n.children), n.body, b, n.isLeaf(), n.hasBody())
	if n.isLeaf() {
		if n.hasBody() {
			n.convertToInternal()
			// fmt.Println("After convertToInternal()", len(n.children), n.body)
			i := n.locationToChildrenIndex(b.position)
			n.children[i].addBody(b)
		} else {
			n.body = b
		}
	} else {
		n.recalculateCenterOfMass(b)
		i := n.locationToChildrenIndex(b.position)
		n.children[i].addBody(b)
	}
}

func (n *node) contains(b *body) bool {
	return b.position.x < n.location.x+n.width/2 && b.position.x > n.location.x-n.width/2 && b.position.y < n.location.y+n.width/2 && b.position.y > n.location.y-n.width/2
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) hasBody() bool {
	return n.body != nil
}

func (n *node) convertToInternal() {
	// fmt.Println(n.location, n.width)
	childLocations := childLocations(n.location, n.width)
	// fmt.Println(childLocations)
	for i := 0; i < 4; i++ {
		// fmt.Println(childLocations[i], n.width)
		n.children = append(n.children, node{location: childLocations[i], width: n.width / 2})
	}

	pTmp := n.body
	n.body = nil
	n.addBody(pTmp)
}

func (n *node) recalculateCenterOfMass(b *body) {
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

func (n *node) calculateForceOnBody(b *body) {
	if n.hasBody() && n.body != b {
		b.addForce(n.body)
	} else {
		threshold := n.width / math.Sqrt(math.Pow(n.centerOfMass.x-b.position.x, 2)+math.Pow(n.centerOfMass.y-b.position.y, 2))
		if threshold < theta {
			b.addForce(&body{position: n.centerOfMass, mass: n.totalMass})
		} else {
			for i := range n.children {
				n.children[i].calculateForceOnBody(b)
			}
		}
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	input := []*Location{
		&Location{13, -13, -2},
		&Location{16, 2, -15},
		&Location{7, -18, -12},
		&Location{-3, -8, -8},
	}
	moons := []*Moon{}
	for _, loc := range input {
		moon := &Moon{loc, &Velocity{0, 0, 0}}
		moons = append(moons, moon)
	}

	steps := 1000
	for i := 0; i < steps; i++ {
		scanMoons(moons)
	}

	totalEnergy := 0
	for _, moon := range moons {
		potentialEnergy := int(
			math.Abs(float64(moon.loc.x)) +
				math.Abs(float64(moon.loc.y)) +
				math.Abs(float64(moon.loc.z)))
		kineticEnergy := int(
			math.Abs(float64(moon.vel.x)) +
				math.Abs(float64(moon.vel.y)) +
				math.Abs(float64(moon.vel.z)))

		totalEnergy += potentialEnergy * kineticEnergy
	}

	fmt.Println("Part 1 : ", totalEnergy)
}

func scanMoons(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		moon := moons[i]
		for j := i + 1; j < len(moons); j++ {
			neighbor := moons[j]
			modifyVelocity(moon, neighbor)
		}
		moon.loc.x += moon.vel.x
		moon.loc.y += moon.vel.y
		moon.loc.z += moon.vel.z
	}
}

func modifyVelocity(moon1 *Moon, moon2 *Moon) {
	if moon1.loc.x > moon2.loc.x {
		moon1.vel.x--
		moon2.vel.x++
	} else if moon1.loc.x < moon2.loc.x {
		moon1.vel.x++
		moon2.vel.x--
	}

	if moon1.loc.y > moon2.loc.y {
		moon1.vel.y--
		moon2.vel.y++
	} else if moon1.loc.y < moon2.loc.y {
		moon1.vel.y++
		moon2.vel.y--
	}

	if moon1.loc.z > moon2.loc.z {
		moon1.vel.z--
		moon2.vel.z++
	} else if moon1.loc.z < moon2.loc.z {
		moon1.vel.z++
		moon2.vel.z--
	}
}

type Moon struct {
	loc *Location
	vel *Velocity
}
type Location struct {
	x int
	y int
	z int
}
type Velocity = Location

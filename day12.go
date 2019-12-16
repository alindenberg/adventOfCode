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
		moons = append(moons, moon(loc, &Velocity{0, 0, 0}))
	}

	// only run one part bc of memory sharing and original state gets lost
	fmt.Println("Part 1 or 2 ?")
	var partNum int
	fmt.Scanf("%d", &partNum)
	if partNum == 1 {
		steps := 1000
		for i := 0; i < steps; i++ {
			scanMoons(moons)
		}
		fmt.Println("Part 1 : ", getTotalEnergy(moons))
	} else {
		fmt.Println("Part 2 : ", findRepeatOrbit(moons))
	}

}

func getTotalEnergy(moons []*Moon) int {
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
	return totalEnergy
}
func scanMoons(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		moon := moons[i]
		for j := i + 1; j < len(moons); j++ {
			neighbor := moons[j]
			modifyVelocity(moon, neighbor, true, true, true)
		}
		moon.updatePosition("X")
		moon.updatePosition("Y")
		moon.updatePosition("Z")
	}
}

func modifyVelocity(moon1 *Moon, moon2 *Moon, modifyX bool, modifyY bool, modifyZ bool) {
	if modifyX {
		if moon1.loc.x > moon2.loc.x {
			moon1.vel.x--
			moon2.vel.x++
		} else if moon1.loc.x < moon2.loc.x {
			moon1.vel.x++
			moon2.vel.x--
		}
	}

	if modifyY {
		if moon1.loc.y > moon2.loc.y {
			moon1.vel.y--
			moon2.vel.y++
		} else if moon1.loc.y < moon2.loc.y {
			moon1.vel.y++
			moon2.vel.y--
		}
	}

	if modifyZ {
		if moon1.loc.z > moon2.loc.z {
			moon1.vel.z--
			moon2.vel.z++
		} else if moon1.loc.z < moon2.loc.z {
			moon1.vel.z++
			moon2.vel.z--
		}
	}
}

func findRepeatOrbit(moons []*Moon) int {
	timeDimensions := make(map[string]int)
	for _, dimension := range []string{"X", "Y", "Z"} {
		atStart := false
		for !atStart {
			for i := 0; i < len(moons); i++ {
				moon := moons[i]
				for j := i + 1; j < len(moons); j++ {
					neighbor := moons[j]
					modifyVelocity(moon, neighbor, dimension == "X", dimension == "Y", dimension == "Z")
				}
				moon.updatePosition(dimension)
			}
			timeDimensions[dimension]++

			allMoonsAtStart := true
			for _, moon := range moons {
				if !moon.atStart(dimension) {
					allMoonsAtStart = false
					break
				}
			}
			atStart = allMoonsAtStart
		}
	}

	return lcmOf3(timeDimensions["X"], timeDimensions["Y"], timeDimensions["Z"])
}

func (moon *Moon) updatePosition(dimension string) {
	switch dimension {
	case "X":
		moon.loc.x += moon.vel.x
	case "Y":
		moon.loc.y += moon.vel.y
	case "Z":
		moon.loc.z += moon.vel.z
	}
}

func (moon *Moon) atStart(dimension string) bool {
	switch dimension {
	case "X":
		if moon.startingLocation.x == moon.loc.x && moon.vel.x == 0 {
			return true
		}
		return false
	case "Y":
		if moon.startingLocation.y == moon.loc.y && moon.vel.y == 0 {
			return true
		}
		return false
	case "Z":
		if moon.startingLocation.z == moon.loc.z && moon.vel.z == 0 {
			return true
		}
		return false
	}
	return false
}

func lcmOf3(x int, y int, z int) int {
	return lcm(x, lcm(y, z))
}

func lcm(x int, y int) int {
	return (x * y) / GCD(x, y)
}

func GCD(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

func moon(loc *Location, vel *Velocity) *Moon {
	return &Moon{
		loc,
		vel,
		*loc,
	}
}

type Moon struct {
	loc              *Location
	vel              *Velocity
	startingLocation Location
}
type Location struct {
	x int
	y int
	z int
}
type Velocity = Location

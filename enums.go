package main

import "fmt"

type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}

type PlanetName uint

func (pn PlanetName) name() string {
	return planets[pn]
}
func (pn PlanetName) ordinal() int {
	return int(pn)
}
func (pn PlanetName) String() string {
	return planets[pn]
}
func (pn PlanetName) values() *[]string {
	return &planets
}

const (
	MERCURY PlanetName = iota
	VENUS
	EARTH
	MARS
	JUPITER
	SATURN
	URANUS
	NEPTUNE
	PLUTO
)

var planets = []string{"MERCURY", "VENUS",
	"EARTH",
	"MARS",
	"JUPITER",
	"SATURN",
	"URANUS",
	"NEPTUNE",
	"PLUTO"}

func main() {
	fmt.Println(MERCURY)
	fmt.Println(MERCURY.name())
	fmt.Println(MERCURY.ordinal())
	fmt.Println(*MERCURY.values())
}

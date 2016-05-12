package CarStuff

type Car struct {
	NumberOfDoors int
	Cylinders     int
}

func (c *Car) GetDoors() int {
	return c.NumberOfDoors
}

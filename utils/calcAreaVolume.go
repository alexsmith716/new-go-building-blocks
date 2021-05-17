package utils


type CalcAreaVolume struct{ length, width, height int }

// area: l x w
// volume: l x w x h

// attach inaccessible method `area()` to struct `CalcAreaVolume`
// modifying the original struct fields
func (c CalcAreaVolume) calculateArea() int {
	return c.length * c.width
}

// attach accessible method `SetSize()` to struct `CalcAreaVolume`
// **** modifying the original struct fields here ****
func (c *CalcAreaVolume) SetStructFields(l, w, h int) {
	c.length = l
	c.width = w
	c.height = h
}

// attach accessible method `GetArea()` to struct `CalcAreaVolume`
// modifying the original struct fields
func (c CalcAreaVolume) MethodArea() int {
	return c.calculateArea()
}

// attach accessible method `GetVolume()` to struct `CalcAreaVolume`
// modifying the original struct fields
func (c CalcAreaVolume) MethodVolume() int {
	return c.length * c.width * c.height
}

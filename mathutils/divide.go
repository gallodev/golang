package mathutils

func GetPercentOfLifeTime(age, maxAge float32) float32 {
	return (divide(age, maxAge) * 100)
}

// Divide returns the division of x and y
func divide(x,y float32) float32 {
	return x / y
}
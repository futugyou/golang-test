package tempconv

func CTof(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

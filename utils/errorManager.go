package errormanager

//Check verify if have errors
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

package sequentialgenerator

func GenerateInts() []int {
	r := []int{}
	count := 10
	// count := 40
	for i := 0; i < count; i++ {
		r = append(r, i)
	}
	return r
}

func GenerateIntsCon() <-chan int {
	rc := make(chan int)
	go func() {
		defer close(rc)

		for i := 0; i < 10; i++ {
			rc <- i
		}
	}()
	return rc
}

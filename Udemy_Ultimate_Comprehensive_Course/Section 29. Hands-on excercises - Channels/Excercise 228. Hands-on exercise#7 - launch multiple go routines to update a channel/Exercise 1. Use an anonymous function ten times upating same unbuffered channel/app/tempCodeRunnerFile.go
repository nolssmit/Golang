	for k := 0; k < 100; k++ {  // because we know the number of values in C
		fmt.Println(k, <-c)
	}
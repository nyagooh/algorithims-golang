package checkpoint

 func Max(a []int)int{
	max := a[0]
	for _, ch := range a {
		if ch > max {
			max = ch
		}
	}
	return max
 }
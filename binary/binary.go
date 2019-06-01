package binary


func HammingWeight(x int) int {
	var counter int = 0

	for x != 0 {
		if x%2 == 1 {
			counter++
		}

		x = x >> 1
	}

	return counter
}

func PowOf2(pow int) int {
	return 1 << uint(pow)
}

func Log2(bits int) int {
	log := 0
	if ((bits & 0xffff0000) != 0) {
		bits >>= 16
		log = 16
	}
	if (bits >= 256) {
		bits >>= 8
		log += 8
	}
	if (bits >= 16) {
		bits >>= 4
		log += 4
	}
	if (bits >= 4) {
		bits >>= 2
		log += 2
	}

	return log + (bits >> 1)
}


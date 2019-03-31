package main

func main() {

}

func judgeCircle(moves string) bool {

	if len(moves)%2 != 0 {
		return false
	}

	horizon, vec := 0, 0
	for _, move := range moves {
		switch move {
		case 'L':
			horizon += 1
		case 'R':
			horizon -= 1
		case 'U':
			vec += 1
		case 'D':
			vec -= 1
		}
	}

	return horizon == 0 && vec == 0
}

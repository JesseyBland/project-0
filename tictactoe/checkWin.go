package main

func checkWin() bool {
	for i := 0; i < len(win); i++ {
		wConditions()
		if win[i] {
			return true
		}
	}
	return false
}

package main

func repeatedCharacter(s string) byte {
	var output byte

	tmpmap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := tmpmap[s[i]]; !ok {
			tmpmap[s[i]] = 1
		} else {
			output = s[i]
			break
		}
	}

	return output
}

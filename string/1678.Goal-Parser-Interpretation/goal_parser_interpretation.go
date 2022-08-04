package main

func interpret(command string) string {
	var result string

	for i := 0; i < len(command); i++ {
		if string(command[i]) == "(" && string(command[i+1]) == ")" {
			result += "o"
			i++
			continue
		} else if string(command[i]) != "(" && string(command[i]) != ")" {
			result += string(command[i])
		}
	}

	return result
}

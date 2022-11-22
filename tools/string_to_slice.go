package tools

func StringToSlice(rawString string) []string {
	var result []string
	var tmp string
	count := 0

	for count < len(rawString) {
		if rawString[count] == '"' {
			tmp = ""
			count++

			for {
				if rawString[count] == '"' {
					result = append(result, tmp)
					break
				} else {
					tmp += string(rawString[count])
					count++
				}
			}
		}

		count++
	}

	return result
}

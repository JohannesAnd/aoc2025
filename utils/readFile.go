package utils

func ReadFile(path string) ([]string, error) {
	lr, err := NewLineReader(path)

	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		lines = append(lines, line)
	}

	return lines, nil
}

package methods

func DefaultStringReplacer(input string, defaultOutput string) string {
	if len(input) > 0 {
		return input
	}

	return defaultOutput
}

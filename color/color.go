package color

func Green(text string) string {
	return "\033[32m" + text + "\033[0m"
}

func Red(text string) string {
	return "\033[31m" + text + "\033[0m"
}

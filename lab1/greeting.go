package greeting

func Greet(name string, style string) string {
	if name == "" {
		name = "Poridhian"
	}

	prefix := getPrefix(style)
	return prefix + name + "!"
}

func getPrefix(style string) string {
	switch style {
	case "casual":
		return "Hey, "
	case "formal":
		return "Good day, "
	default:
		return "Hello, "
	}
}

package movie

func FindName(imbdID string) string {
	switch imbdID {
	case "tt4154796":
		return "Avenger Endgame"
	case "tt1825683":
		return "Black Panther"
	}
	return "not found."
}

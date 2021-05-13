package main

func daySelectorIf(day int) string {
	if day == 1 {
		return "Sunday"
	} else if day == 2 {
		return "Monday"
	} else if day == 3 {
		return "Tuesday"
	} else if day == 4 {
		return "Wednesday"
	} else if day == 5 {
		return "Thursday"
	} else if day == 6 {
		return "Friday"
	} else if day == 7 {
		return "Saturday"
	} else {
		return "None"
	}
}

func daySelectorSwitch(day int) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "None"
	}
}

func main() {

}

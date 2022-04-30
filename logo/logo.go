package logo

import "fmt"

func Print() {
	var gotter string
	gotter = "   _____    ____    _______   _______   ______   _____\n"
	gotter += "  / ____|  / __ \\  |__   __| |__   __| |  ____| |  __ \\ \n"
	gotter += " | |  __  | |  | |    | |       | |    | |__    | |__) |\n"
	gotter += " | | |_ | | |  | |    | |       | |    |  __|   |  _  /\n"
	gotter += " | |__| | | |__| |    | |       | |    | |____  | | \\ \\\n"
	gotter += "  \\_____|  \\____/     |_|       |_|    |______| |_|  \\_\\\n"
	fmt.Println(gotter)
}

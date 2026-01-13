package util

import "fmt"

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func PrintTitle() {
	fmt.Println(" ________                 _             _    _____ _           _")
	fmt.Println("|__   __|                (_)           | |  / ____| |         | |")
	fmt.Println("   | | ___ _ __ _ __ ___  _ _ __   __ _| | | |    | |__   __ _| |_")
	fmt.Println("   | |/ _ \\ '__| '_ ` _ \\| | '_ \\ / _` | | | |    | '_ \\ / _` | __|")
	fmt.Println("   | |  __/ |  | | | | | | | | | | (_| | | | |____| | | | (_| | |_")
	fmt.Println("   |_|\\___|_|  |_| |_| |_|_|_| |_|\\__,_|_|  \\_____|_| |_|\\__,_|\\__|")
}

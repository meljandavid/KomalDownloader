package frontend

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCreds() (em, pw, m string, set []string) {
	fmt.Println("  _  __                     _   _          _        _ _        ")
	fmt.Println(" | |/ /                    | | | |        | |      | | |        ")
	fmt.Println(" | ' / ___  _ __ ___   __ _| | | |     ___| |_ ___ | | |_ ___   ")
	fmt.Println(" |  < / _ \\| '_ ` _ \\ / _` | | | |    / _ \\ __/ _ \\| | __/ _ \\  ")
	fmt.Println(" | . \\ (_) | | | | | | (_| | | | |___|  __/ || (_) | | || (_) | ")
	fmt.Println(" |_|\\_\\___/|_| |_| |_|\\__,_|_| |______\\___|\\__\\___/|_|\\__\\___/  ")

	fmt.Println()
	fmt.Println(" A letöltéshez bejelentkezés szükséges")
	fmt.Println()

	fmt.Printf(" email: ")
	fmt.Scanf("%s\n", &em)
	fmt.Printf(" jelszó: ")
	fmt.Scanf("%s\n", &pw)
	fmt.Printf(" hónap száma két jeggyel (pl 04): ")
	fmt.Scanf("%s\n", &m)

	fmt.Printf(" pontversenyek (szóközzel elválasztva): ")
	rdr := bufio.NewReader(os.Stdin)
	line, _ := rdr.ReadString('\n')
	line = line[:len(line)-2] // wtf?
	line = strings.ToUpper(line)
	set = strings.Split(line, " ")

	return
}

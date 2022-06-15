package frontend

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCreds() (em, pw, m string, set []string) {
	fmt.Printf("email: ")
	fmt.Scanf("%s\n", &em)
	fmt.Printf("password: ")
	fmt.Scanf("%s\n", &pw)
	fmt.Printf("month: ")
	fmt.Scanf("%s\n", &m)

	fmt.Printf("problemsets: ")
	rdr := bufio.NewReader(os.Stdin)
	line, _ := rdr.ReadString('\n')
	line = line[:len(line)-2] // wtf?
	set = strings.Split(line, " ")

	return
}

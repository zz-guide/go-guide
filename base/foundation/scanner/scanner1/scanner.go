package scanner1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TestLineFilters() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		//ucl := strings.ToUpper(string(scanner.token))

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

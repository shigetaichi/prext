package utils

import (
	"bufio"
	"fmt"
	"github.com/thoas/go-funk"
	"os"
)

func Cd(dir string) {
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dir
}

func Exists(name string) bool {
	f, err := os.Stat(name)
	return !(os.IsNotExist(err) || f.IsDir())
}

func Remove(name string) {
	if err := os.Remove(name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func YesOrNo(q string) bool {
	result := true
	fmt.Print(q)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			result = false
			break
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Print(q)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
func Choice(choices []string, fns ...func()) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		index := funk.IndexOfString(choices, scanner.Text())
		if index == -1 {
			fns[1]()
		} else {
			fns[index]()
		}
		break

	}
}

func YarnOrNpm(yarnFlag bool) string {
	if yarnFlag {
		return "yarn"
	} else {
		return "npm"
	}
}

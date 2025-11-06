package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getSessionCookie() string {
	bytes, err := os.ReadFile("./session")
	if err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("session=%s", string(bytes))
}

func buildMain(day string) string {
	return fmt.Sprintf(`
package main

import (
	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day%s",
		c.PlaceholderFunc,
		c.PlaceholderFunc,
	)
}
`, day)
}

func getInput(day string) []byte {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Cookie", getSessionCookie())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

// dirExists returns whether the given file or directory dirExists
func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Expected exactly 1 argument: number of day")
	}
	day := args[0]
	dirPath := fmt.Sprintf("./day%s", day)

	if exists, _ := dirExists(dirPath); exists {
		log.Fatalf("Day %s already exists\n", day)
	}

	input := getInput(day)

	fmt.Printf("Creating directory %s\n", dirPath)
	os.Mkdir(dirPath, 0775)

	fmt.Printf("Creating file %s\n", dirPath+"/main.go")
	os.WriteFile(dirPath+"/main.go", []byte(buildMain(day)), 0775)

	fmt.Printf("Creating empty file %s\n", dirPath+"/input-sample.go")
	os.WriteFile(dirPath+"/input-sample.txt", []byte{}, 0775)

	fmt.Printf("Creating file %s\n", dirPath+"/input.txt")
	os.WriteFile(dirPath+"/input.txt", input, 0775)
}

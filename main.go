package main

import (
	"os"
	"strings"
	"os/exec"
)

func main() {
	file, _ := os.ReadFile("input.txt")

	program := strings.ReplaceAll(string(file), "package", "package main")
	program = strings.ReplaceAll(program, "import", "import \"fmt\"")

	temp := "temp.go" // create a temporary file to hold the program while you run it.
	os.WriteFile(temp, []byte(program), 0644)
	defer os.Remove(temp) // remove file from your files to avoid errors
	cmd := exec.Command("go", "run", temp)
	out, _ := cmd.CombinedOutput() // combine output
	os.WriteFile("output.txt", []byte(out), 0644)
}

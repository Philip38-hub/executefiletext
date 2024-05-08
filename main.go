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

	temp := "temp.go"
	os.WriteFile(temp, []byte(program), 0644)
	defer os.Remove(temp)
	cmd := exec.Command("go", "run", temp)
	out, _ := cmd.CombinedOutput()
	os.WriteFile("output.txt", []byte(out), 0644)
}

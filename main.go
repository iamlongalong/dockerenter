package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

var (
	templates = &promptui.SelectTemplates{
		Label:    "✨ {{ . | green}}",
		Active:   "➤ {{ .Name | cyan  }} {{ .Image | yellow }} {{ .ID }}",
		Inactive: "  {{.Name | faint}} {{ .Image }} {{ .ID }}",
	}
)

func main() {
	transferArgs := os.Args[1:]

	runCmdArgs := []string{"exec", "-it", ""}

	contailerName := choose()

	if len(transferArgs) == 0 {
		runCmd := "bash"
		if !testForBash(contailerName) {
			runCmd = "sh"
		}

		runCmdArgs = append(runCmdArgs, runCmd)
	} else {
		runCmdArgs = append(runCmdArgs, transferArgs...)
	}

	runCmdArgs[2] = contailerName

	cmd := exec.Command("docker", runCmdArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func testForBash(name string) bool {
	cmd := exec.Command("docker", "exec", name, "which", "bash")

	_, err := cmd.Output()
	if err != nil {
		return false
	}

	return true
}

func choose() string {
	containers, err := getContainers()
	if err != nil {
		log.Printf("get contailers fail : %s", err)
		os.Exit(1)
	}

	prompt := promptui.Select{
		Label:        "select docker container ~",
		Items:        containers,
		Templates:    templates,
		Size:         10,
		HideSelected: true,
		Searcher: func(input string, index int) bool {
			node := containers[index]
			content := fmt.Sprintf("%s %s %s", node.Name, node.Image, node.ID)

			return strings.Contains(content, input)
		},
	}

	idx, _, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}

	return containers[idx].Name
}

type dockerContainer struct {
	ID    string
	Image string
	Name  string
}

func getContainers() ([]*dockerContainer, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{.ID}}\t{{.Image}}\t{{.Names}}")

	// 执行 Docker 命令，并获取输出结果
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	res := make([]*dockerContainer, 0)

	// 解析输出结果，获取每个 container 的 ID、Image 和 Name
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) == 3 {
			id := parts[0]
			image := parts[1]
			name := parts[2]
			c := &dockerContainer{
				ID:    id,
				Image: image,
				Name:  name,
			}

			iparts := strings.Split(image, "/")
			image = iparts[len(iparts)-1]
			c.Image = image

			res = append(res, c)
		} else {
			return nil, fmt.Errorf("scan fail of : %s", line)
		}
	}

	return res, nil
}

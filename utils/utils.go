package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the YAML file
type Config struct {
	ProjectName   string   `yaml:"project_name"`
	ProjectType   string   `yaml:"project_type"`
	Compiler      string   `yaml:"compiler"`
	CompilerFlags []string `yaml:"compiler_flags"`
	SrcPath       string   `yaml:"src_path"`
	SrcFiles      []string `yaml:"src_files"`
	BuildPath     string   `yaml:"build_path"`
	IncluderDirs  []string `yaml:"include_dirs"`
	LibDirs       []string `yaml:"lib_dirs"`
	Libs          []string `yaml:"libs"`
}

// Read user input with formatting
func readUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Cleans input by handling both comma and space-separated values properly
func cleanInput(input string) []string {
	input = strings.TrimSpace(input)
	rawItems := strings.Split(input, ",")

	var items []string
	for _, item := range rawItems {
		trimmedItem := strings.TrimSpace(item) // Ensure no leading spaces remain
		if trimmedItem != "" {                 // Prevent empty items
			items = append(items, trimmedItem)
		}
	}
	return items
}

// Setup YAML configuration interactively
func SetupYaml() {
	fmt.Println("\033[1;36mWelcome to the Nyati Build System Setup!\033[0m")
	fmt.Println("\033[1;33mPlease enter the following details to create your project configuration:\033[0m")

	config := &Config{}

	config.ProjectName = readUserInput("\033[1;32mEnter Project Name: \033[0m")
	config.ProjectType = readUserInput("\033[1;32mEnter Project Type (e.g., Console, Server): \033[0m")
	config.Compiler = readUserInput("\033[1;32mEnter Compiler (e.g., g++, clang++): \033[0m")

	flags := readUserInput("\033[1;32mEnter Compiler Flags (space/comma-separated, e.g., std=c++20 Wall): \033[0m")
	config.CompilerFlags = cleanInput(flags)

	config.SrcPath = readUserInput("\033[1;32mEnter Source Path (e.g., src): \033[0m")

	srcFiles := readUserInput("\033[1;32mEnter Source Files (space/comma-separated, e.g., main.cpp utils.cpp): \033[0m")
	config.SrcFiles = cleanInput(srcFiles)

	config.BuildPath = readUserInput("\033[1;32mEnter Build Path (e.g., build, press Enter for default '.'): \033[0m")
	if strings.TrimSpace(config.BuildPath) == "" {
		config.BuildPath = "."
	}

	includeDirs := readUserInput("\033[1;32mEnter Include Directories (space/comma-separated, e.g., include libs/include): \033[0m")
	config.IncluderDirs = cleanInput(includeDirs)

	libDirs := readUserInput("\033[1;32mEnter Library Directories (space/comma-separated, e.g., lib libs/lib): \033[0m")
	config.LibDirs = cleanInput(libDirs)

	libs := readUserInput("\033[1;32mEnter Libraries (space/comma-separated, e.g., ws2_32 mswsock): \033[0m")
	config.Libs = cleanInput(libs)

	// Convert to YAML
	cg, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("\033[1;31mError marshaling YAML:\033[0m", err)
		return
	}

	// Write YAML to file
	err = os.WriteFile("project.yaml", cg, 0644)
	if err != nil {
		fmt.Println("\033[1;31mError writing YAML file:\033[0m", err)
		return
	}

	fmt.Println("\033[1;32mProject.yaml created successfully!\033[0m")
}

// Initialize Configuration with default values
func InitConfig() {
	fmt.Println("\033[1;36mNyati Config System Loading...\033[0m")
	time.Sleep(2 * time.Second)
	fmt.Println("\033[1;32mNyati Config System Loaded!\033[0m")

	config := &Config{}
	// 	ProjectName:   "Nyati",
	// 	ProjectType:   "Console",
	// 	SrcPath:       "src",
	// 	SrcFiles:      []string{"main.cpp", "users.cpp"},
	// 	BuildPath:     ".",
	// 	Compiler:      "g++",
	// 	CompilerFlags: []string{"std=c++20"},
	// 	IncluderDirs:  []string{"include", "D:/Devlibs/include"},
	// 	LibDirs:       []string{"lib", "D:/Devlibs/lib"},
	// 	Libs:          []string{"ws2_32", "mswsock"},
	// }

	// Convert to YAML
	cg, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("\033[1;31mError marshaling YAML:\033[0m", err)
		return
	}

	fmt.Println("\033[1;34mGenerated YAML:\033[0m")
	fmt.Println(string(cg))

	err = os.WriteFile("project.yaml", cg, 0644)
	if err != nil {
		fmt.Println("\033[1;31mError writing YAML file:\033[0m", err)
		return
	}

	fmt.Println("\033[1;32mProject.yaml created!\033[0m")
}

// Function to update a field in the YAML file
func UpdateConfigField(field, value string) error {
	data, err := os.ReadFile("project.yaml")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("error parsing YAML: %v", err)
	}

	switch field {
	case "project_name":
		config.ProjectName = value
	case "project_type":
		config.ProjectType = value
	case "compiler":
		config.Compiler = value
	case "compiler_flags":
		config.CompilerFlags = cleanInput(value)
	case "src_path":
		config.SrcPath = value
	case "src_files":
		config.SrcFiles = cleanInput(value)
	case "build_path":
		config.BuildPath = value
	case "include_dirs":
		config.IncluderDirs = cleanInput(value)
	case "lib_dirs":
		config.LibDirs = cleanInput(value)
	case "libs":
		config.Libs = cleanInput(value)
	default:
		return fmt.Errorf("unknown field: %s", field)
	}

	updatedData, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("error encoding YAML: %v", err)
	}

	err = os.WriteFile("project.yaml", updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Println("\033[1;32mConfiguration updated successfully!\033[0m")
	return nil
}

// Build the project using shell commands
func Build() {
	config := &Config{}
	file, err := os.ReadFile("project.yaml")
	if err != nil {
		fmt.Println("\033[1;31mError reading file:\033[0m", err)
		return
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("\033[1;31mError unmarshaling YAML:\033[0m", err)
		return
	}

	fmt.Println("\033[1;36mBuilding...\033[0m")
	command := exec.Command(config.Compiler, append(config.CompilerFlags, "-o", config.BuildPath+"/"+config.ProjectName)...)

	cmd := exec.Command("bash", "-c", command.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("\033[1;31mBuild failed:\033[0m", err)
		return
	}

	fmt.Println("\033[1;32mBuild complete!\033[0m")
}

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

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

func InitConfig() {
	fmt.Println("Nyati Config System Loading...")
	time.Sleep(2 * time.Second)
	fmt.Println("Nyati Config System Loaded!")

	// config := &Config{
	// 	ProjectName:   "Nyati",
	// 	ProjectType:   "Console",
	// 	SrcPath:       "src",
	// 	SrcFiles:      []string{"main.cpp", "users.cpp"},
	// 	BuildPath:     "",
	// 	Compiler:      "g++",
	// 	CompilerFlags: []string{"std=c++20"},
	// 	IncluderDirs:  []string{"include", "D:/Devlibs/include"},
	// 	LibDirs:       []string{"lib", "D:/Devlibs/lib"},
	// 	Libs:          []string{"ws2_32", "mswsock"},
	// }

	config := &Config{}

	if config.BuildPath == "" || config.BuildPath == " " {
		config.BuildPath = "."
	}

	cg, err := yaml.Marshal(config)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(cg)
	fmt.Println(string(cg))

	err = os.WriteFile("project.yaml", cg, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Project.yaml created!\nReading File in 10 seconds!")
	// time.Sleep(10 * time.Second)

}

func ModifyConfig() {
	fmt.Println("Reading File To Modify Contents...")

	file, err := os.ReadFile("project.yaml")

	if err != nil {
		fmt.Println(err)
	}

	config := &Config{}
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err)
	}

	config.ProjectType = "server"

	config.AppendLibs("ws2_32")
	config.AppendLibs("glfw3")

	cg3, err := yaml.Marshal(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cg3)
	fmt.Println(string(cg3))

	err = os.WriteFile("project.yaml", cg3, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Modified files successfully!")
}

func (config *Config) AppendLibs(value string) {
	exists := false
	for _, lib := range config.Libs {
		if lib == value {
			exists = true
			fmt.Printf("%s is already included!\n", value)
			break
		}
	}

	if !exists {
		config.Libs = append(config.Libs, value)
	}
}

func Build() {
	config := &Config{}
	file, err := os.ReadFile("project.yaml")

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
	}

	var shell, flag, pathsep string

	if os.PathSeparator == '\\' { // Windows
		shell, flag, pathsep = "cmd", "/C", `\\`
	} else { // Linux/macOS
		shell, flag, pathsep = "bash", "-c", `\`
	}

	fmt.Println("Building...")
	fmt.Printf("Building %s with %s...\n", config.ProjectName, config.Compiler)
	fmt.Printf("Using flags: %v\n", config.CompilerFlags)
	fmt.Printf("Using include directories: %v\n", config.IncluderDirs)
	fmt.Printf("Using library directories: %v\n", config.LibDirs)
	fmt.Printf("Using libraries: %v\n", config.Libs)

	var command strings.Builder

	command.WriteString(config.Compiler)

	for _, flag := range config.CompilerFlags {
		cmd := " -" + flag
		command.WriteString(cmd)
	}

	for _, srcFile := range config.SrcFiles {
		cmd := " " + config.SrcPath + pathsep + srcFile
		command.WriteString(cmd)
	}

	for _, incdir := range config.IncluderDirs {
		cmd := " -I" + incdir
		command.WriteString(cmd)
	}

	for _, libdir := range config.LibDirs {
		cmd := " -L" + libdir
		command.WriteString(cmd)
	}

	for _, lib := range config.Libs {
		cmd := " -l" + lib
		command.WriteString(cmd)
	}

	command.WriteString(" -o " + config.BuildPath + "/" + config.ProjectName)

	// fmt.Println("Command: ", command.String())

	cmd := exec.Command(shell, flag, command.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		fmt.Println("Error executing command:", err)
	}

	// fmt.Println("Output:", string(output))

	runcmd := exec.Command(shell, flag, config.BuildPath+pathsep+config.ProjectName)
	runcmd.Stdout = os.Stdout
	runcmd.Stderr = os.Stderr

	if err != nil {
		fmt.Println(err)
	}

	runcmd.Run()
	// Here you would add the actual build command using the values from the config
	// For example, you could use os/exec to run a shell command to build the project
}

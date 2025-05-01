# Nyati Build System  

```
                    ███╗   ██╗██╗   ██╗ █████╗ ████████╗██╗
                    ████╗  ██║╚██╗ ██╔╝██╔══██╗╚══██╔══╝██║
                    ██╔██╗ ██║ ╚████╔╝ ███████║   ██║   ██║
                    ██║╚██╗██║  ╚██╔╝  ██╔══██║   ██║   ██║
                    ██║ ╚████║   ██║   ██║  ██║   ██║   ██║
                    ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝   ╚═╝
                                                       buffalo
```

Nyati Build System is an in-progress C/C++ build system designed to simplify the process of building and running your C/C++ projects. The name "Nyati" means buffalo in Shona, symbolizing strength and reliability. This project is built using Go and Cobra.  

## Features  

- **Project Initialization**: Use the `init` command to generate a `project.yaml` file, where you can define your project's properties.  
- **Build and Run**: Use the `dev` command to build and execute your project seamlessly.  

### Planned Features  

1. CLI command to add required YAML fields during initialization.  
2. Live reload capabilities for a smoother development experience.  

## Installation  

To use Nyati Build System, ensure you have Go installed on your system. Clone the repository and build the tool:  

```bash  
git clone https://github.com/your-repo/nyati-build-system.git  
cd nyati-build-system  
go build -o nyati  
```  

## Usage  

### Initialize a Project  

Run the following command to generate a `project.yaml` file:  

```bash  
nyati init  
```  

Edit the generated `project.yaml` to match your project's configuration.  

### Build and Run  

Use the `dev` command to build and run your project:  

```bash  
./nyati dev  
``` 
If you add it to PATH, in your project root dir run:

```bash
nyati dev
```

## `project.yaml` Structure  

Below is the generic structure of the `project.yaml` file: 
```yaml
project_name: ""
project_type: ""
compiler: ""
compiler_flags: []
src_path: ""
src_files: []
build_path: .
include_dirs: []
lib_dirs: []
libs: []
```


```yaml
project_name: "MyProject"
project_type: "executable"
compiler: "gcc"
compiler_flags:
    - "-Wall"
    - "-O2"
src_path: "src"
src_files:
    - "main.c"
    - "utils.c"
build_path: "build"
include_dirs:
    - "include"
lib_dirs:
    - "libs"
libs:
    - "m"
    - "pthread"
```

## Contributing  

Contributions are welcome! Feel free to open issues or submit pull requests to help improve Nyati Build System.  

## License  

This project is licensed under the MIT License.  

---  

Happy coding with Nyati Build System! 🚀  
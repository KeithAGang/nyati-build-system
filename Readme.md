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

Nyati Build System is a versatile **build system** designed to simplify the process of building and running projects across various programming stacks, with a primary focus on **C/C++**. However, it is designed to be extensible, making it adaptable to other setups that lack a simplified build system.

The name **"Nyati"** means buffalo in Shona, symbolizing strength and reliability.

This project is built using **Go** and **Cobra**.

---

## Features

- ✅ **Project Initialization (`init`)** – Generates a `project.yaml` file with default settings.
- ✅ **Project Setup (`setup`)** – Interactive CLI wizard for configuring `project.yaml`.
- ✅ **Modify Configuration (`modify`)** – Allows updating specific fields in `project.yaml` via CLI.
- ✅ **Build and Run (`dev`)** – Compiles and executes your project seamlessly.

### Planned Features

- 🔹 CLI command to dynamically add required YAML fields.
- 🔹 Live reload capabilities for faster iteration.
- 🔹 Support for additional programming stacks through plugins or extensions.

---

## Installation

To use Nyati Build System, ensure you have **Go** installed on your system.

Clone the repository and build the tool:

```bash
git clone https://github.com/KeithAGang/nyati-build-system
cd nyati-build-system
go build nyati.go
```

---

## Usage

### Initialize a Project (`init`)

Run the following command to generate an empty `project.yaml` file:

```bash
nyati init
```

This will create a `project.yaml` with default values, ready to be configured.

---

### Set Up Project (`setup`)

Run the interactive setup wizard to define your project's configuration:

```bash
nyati setup
```

This guides you through entering:

- Project name
- Compiler settings
- Source files
- Include directories
- Library dependencies, etc.

Once completed, it writes the configuration to `project.yaml`.

---

### Modify a Configuration Field (`modify`)

To update a specific field in `project.yaml`, use the `modify` command:

```bash
nyati modify <field_name> <new_value>
```

#### Field Names for Modification

| Field Name       | Description                          | Example Command                              |
|-------------------|--------------------------------------|----------------------------------------------|
| `project_name`    | Name of the project                 | `nyati modify project_name "MyApp"`          |
| `project_type`    | Type of project (Console, Server)   | `nyati modify project_type "Server"`         |
| `compiler`        | Compiler to use (gcc, clang++)      | `nyati modify compiler "clang++"`            |
| `compiler_flags`  | Compilation flags (comma-separated) | `nyati modify compiler_flags "-Wall, -O2"`   |
| `src_path`        | Source file directory               | `nyati modify src_path "src"`                |
| `src_files`       | Source files (comma-separated)      | `nyati modify src_files "main.cpp, utils.cpp"`|
| `build_path`      | Output directory for compiled binary| `nyati modify build_path "bin"`              |
| `include_dirs`    | Include directories (comma-separated)| `nyati modify include_dirs "include, thirdparty"`|
| `lib_dirs`        | Library directories                 | `nyati modify lib_dirs "libs, /usr/local/lib"`|
| `libs`            | Libraries to link (comma-separated) | `nyati modify libs "m, pthread"`            |

---

### Build and Run (`dev`)

Use the `dev` command to compile and execute your project:

```bash
nyati dev
```

If you've added Nyati to your system `PATH`, you can simply run:

```bash
nyati dev
```

---

## Extensibility

Nyati Build System is designed to be extensible, allowing developers to adapt it to other programming stacks or workflows. Future updates will include support for plugins or extensions to handle diverse build environments.

---

## Feedback & Contributions

Your feedback is highly valued! It helps improve Nyati Build System.

🟢 The tool has been tested with MinGW and Clang.  
🟢 If you encounter any issues or have suggestions, please let us know!

### Contributing

✅ Open issues or submit pull requests to help improve Nyati Build System.  
✅ Any contribution, big or small, is appreciated!

---

## License

This project is licensed under the MIT License.

---

## Final Thoughts

Nyati is built to be fast, lightweight, and developer-friendly.

Try it out and simplify your development workflow! 🚀 Let us know how we can make it even better.
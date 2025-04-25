Fast and Easy way to create a base-minimum directory layout for your Go project with the `gogogo` cli.

# Content
1. [About](#about)
2. [Dependencies](#dependencies)
3. [Installation](#installation)
4. [Examples](#examples)
5. [Warning](#warning)
6. [License](#license)

# About
`gogogo` is a CLI tool that helps you create a simple & basic directory layout using either the "pkg" directory or the "mvc" structure (Model, View, Control), and populate it with files like `README.md`, `.gitignore`, `Makefile`, and `main.go`.

example of the final layout:
```shell
provided-directory-path/
├── .git        ## if -g flag was provided
├── .gitignore
├── bin/
├── pkg/        ## default layout
├── Makefile
├── README.md
└── go.mod
```

# Dependencies
1. [Go programming language (v1.24.1 or later)](https://go.dev/doc/install).
2. [git](https://git-scm.com/downloads).
3. [cobra library](https://github.com/spf13/cobra).

# Installation
To start using the `gogogo` cli just run the following command in your terminal. It will build and install `gogogo` to your `$GOBIN` directory:
```shell
go install github.com/ABDsheikho/gogogo@latest
```

# Examples
1. example 1:
    ```shell
    gogogo . github.com/example/test
    ```
    This initializes a new Go module in the current directory, using "github.com/example/test" as the module path.
2. example 2:
    ```shell
    gogogo ./new-folder/sub-folder github.com/example/test -l mvc -g
    ```
    This will initialize new Go module at sub-folder directory (and it will create all required directories if they don't exist), and it will create MVC folders layout inside sub-folder, then it will initialize git repository.

# Warning:
It is not recommended to use `gogogo` in a non-empty directory, as the current implementation will delete the entire directory (including its contents) if an error occurs during the process.

# License:
This software is licensed under CC BY-NC-SA.

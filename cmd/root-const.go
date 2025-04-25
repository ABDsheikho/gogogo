package cmd

const (
	shortDescription string = "Fast and Easy way to create a base-minimum directory layout for your Go project"
	longDescription  string = `gogogo is a CLI tool that helps you create a simple & basic directory layout using either "pkg" 
or "mvc" as base layout, and populate it with files like README, gitignore, Makefile and main.go 

` + "example of the final layout:\n" + "\u001b[33m" +
		`provided-directory-path/
├── .git        ## if -g flag was provided
├── .gitignore
├── bin/
├── pkg/        ## default layout
├── Makefile
├── README.md
└── go.mod
` + "\u001b[0m" + "\n" + warning

	warning string = "\u001b[31mWarning!\u001b[0m" +
		`: It is not recommended to use "gogogo" in a non-empty directory, as the current 
          implementation will delete the entire directory (including its contents) if an error
          occurs during the process.`

	example string = "ex.1:" + "\n  \u001b[33mgogogo . github.com/example/test\u001b[0m" + `
    This initializes a new Go module in the current directory, using "github.com/example/test" as 
    as the module path.

ex.2:` + "\n  \u001b[33mgogogo ./new-folder/sub-folder github.com/example/test -l mvc -g\u001b[0m" + `
    This will initialize new Go module at sub-folder directory (and it will create all required 
    directories if they don't exist), and it will create MVC folders layout inside sub-folder, 
    then it will initialize git repository.`
)

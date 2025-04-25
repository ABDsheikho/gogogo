package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// `validateArgs“ process slice of strings that represents arguments passed to the `gogogo` cli
// and return `dir` which is the absolute path to the desired directory, and `modPath`
// which is the module path.
//
// if an error occurred while processing `dir` to its absolute path,
// `validateArgs` will print the error to stderr, and then exit the program with code 1.
func validateArgs(args []string) (dir, modPath string) {
	dir, err := filepath.Abs(args[0])
	modPath = args[1]

	// set abortAndClear arg to false because no directories or files are made yet.
	handelErr(err, false, "")
	return
}

// `initialModule` initialize go.mod file via go CLI and using provided argument for modPath.
//
// this function takes two arguments, `dir` which is the absolute path of the module
// directory, and `modPath` which is the module path. and then run `go mod init modPath`.
//
// if and error occurred while running `go init mod` command,
// `initialModule` will print the error to stderr, and then exit the program with code 1.
// and the provided `dir` will be wiped-out and removed, because it was assumed that
// it's initially empty.
func initialModule(dir string, modPath string) {
	// run go CLI to initial module, this is equivalent to running `go mod init modPath`
	cmd := exec.Command("go", "mod", "init", modPath)
	_, err := cmd.CombinedOutput()
	handelErr(err, true, dir)
}

// `initialRepository` initialize git repository via git CLI at the provided directory.
//
// this function takes one argument, `dir` which is the absolute path of the module directory,
// and then run `git init dir`.
//
// if and error occurred while running `git init dir` command,
// `initialRepository` will print the error to stderr, and then exit the program with code 1.
// and the provided `dir` will be wiped-out and removed, because it was assumed that
// it's initially empty.
func initialRepository(dir string) {
	// run go CLI to initial module, this is equivalent to running `git init modPath`
	cmd := exec.Command("git", "init", dir)
	_, err := cmd.CombinedOutput()
	handelErr(err, true, dir)
}

// `mkDir` creates directory with `dirName` inside `parent` directory.
//
// this function takes two arguments, `dirName` which is the name of the desired directory,
// and `parent` which is the parent directory.
//
// if and error occurred while running `mkDir` function,
// `mkDir` will print the error to stderr, and then exit the program with code 1.
// and the provided `parent` will be wiped-out and removed, because it was assumed that
// it's initially empty.
func mkDir(parent, dirName string) {
	err := os.Mkdir(filepath.Join(parent, dirName), os.ModePerm)
	handelErr(err, true, parent)
}

// `mkDir` creates file named `fileName` inside `dir` directory with content as slice of `lines`.
//
// this function takes two main arguments, `dir` which is the name of the current directory,
// `fileName` which is the name of the desired file, and the reset are variadic arguments that
// represents slice of lines that needs to be written to the file.
//
// if and error occurred while running `mkFile` function,
// `mkFile` will print the error to stderr, and then exit the program with code 1.
// and the provided `dir` will be wiped-out and removed, because it was assumed that
// it's initially empty.
func mkFile(dir, fileName string, lines ...string) {
	path := filepath.Join(dir, fileName)
	file, err := os.Create(path)
	handelErr(err, true, dir)
	defer file.Close()

	_, err = file.WriteString(strings.Join(lines, "\n") + "\n") // last \n is trailing newline
	handelErr(err, true, dir)
}

// `mkLayout` creates layout of directories using either "pkg" or "mvc" structure inside `dir`
// directory.
//
// this function takes two arguments, `dir` which is the current directory that will be populated
// with layout, and `layout` which represents the desired initial directories needed.
//
// `mkLayout` currently handle only two cases "pkg" and "mvc".
// any other cases will be already handled by `validateFlags` function, but in case something
// went wrong, `mkLayout` also handle a default case with an error and deletion of the directory.
func mkLayout(dir, layout string) {
	switch layout {
	case "pkg":
		mkDir(dir, "pkg")
	case "mvc":
		mkDir(dir, "model")
		mkDir(dir, "view")
		mkDir(dir, "control")
	default:
		handelErr(fmt.Errorf("%sError! unhandled layout value!%s check `mkLayout` function",
			"\u001b[31m", "\u001b[0m",
		),
			true, dir,
		)
	}
	mkDir(dir, "bin")
}

// `handelErr` is a utility function that handles any error, and print the error to stderr.
//
// this function takes three arguments, `err` which is the error that needs to be handled,
// `abortAndClear` which is a boolean flag indicts whether every change should be enrolled and
// deleted, and `dir` which is the current directory that the error occurred.
//
// if you set `abortAndClear` to false, then it's recommended to set `dir` to an empty string ""
func handelErr(err error, abortAndClear bool, dir string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// fmt.Println(err)
		if abortAndClear {
			os.RemoveAll(dir)
		}
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type cmdFlags struct {
	layout  string
	initGit bool
}

func (cf *cmdFlags) ignoreCases() {
	cf.layout = strings.ToLower(cf.layout)
}

var cf = &cmdFlags{}

var rootCmd = &cobra.Command{
	Use:     "gogogo directory mod-path",
	Short:   shortDescription,
	Long:    longDescription,
	Example: example,
	Args:    cobra.ExactArgs(2),
	PreRun:  validateFlags,
	Run:     gogogo,
}

func init() {
	rootCmd.Flags().BoolVarP(&cf.initGit, "init-git", "g", false, `initial git repository immediately or not. (default false)`)
	rootCmd.Flags().StringVarP(&cf.layout, "layout", "l", "pkg",
		fmt.Sprintf(`set directory layout as either %s"pkg"%s or %s"mvc"%s.`,
			"\u001b[33m", "\u001b[0m",
			"\u001b[33m", "\u001b[0m",
		),
	)
}

func Execute() {
	err := rootCmd.Execute()
	handelErr(err, false, "")
}

func validateFlags(_ *cobra.Command, _ []string) {
	// according to cobra process, this function will run before execution, thus we better handle
	// any checks here.
	cf.ignoreCases()
	switch cf.layout {
	case "pkg": // good, do nothing.
	case "mvc": // good, do nothing.
	default: // No
		fmt.Fprintf(
			os.Stderr,
			"%sWRONG INPUT%s: flag value for -l (--layout) can only be %s\"pkg\"%s or %s\"mvc\"%s\n",
			"\u001b[31m", "\u001b[0m",
			"\u001b[33m", "\u001b[0m",
			"\u001b[33m", "\u001b[0m",
		)
		os.Exit(1)
	}
}

func gogogo(cmd *cobra.Command, args []string) {
	dir, modPath := validateArgs(args)
	err := os.MkdirAll(dir, os.ModePerm)
	// set abortAndClear arg to false because no directories or files are made yet.
	handelErr(err, false, "")

	// step into dir, this is equivalent to running `cd dir`
	err = os.Chdir(dir)
	handelErr(err, true, dir)

	initialModule(dir, modPath)
	mkLayout(dir, cf.layout)
	mkFile(dir, "README.md")
	mkFile(dir, "Makefile")
	mkFile(dir, "main.go",
		"package main", // set package for main.go as main
	)
	mkFile(dir, ".gitignore",
		`bin/`, // ignore everything in bin directory
	)

	if cf.initGit {
		initialRepository(dir)
	}
	fmt.Println("\u001b[32mDone\u001b[0m")
}

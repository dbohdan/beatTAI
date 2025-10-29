package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const leapsecsInstallPath = "/usr/local/etc/"

var prevHeadingLevel int = -1

func main() {
	var (
		noGbt    bool
		noGbtgui bool
		noI9w    bool
	)

	flag.BoolVar(&noGbt, "no-gbt", false, "Skip gbt test")
	flag.BoolVar(&noGbtgui, "no-gbtgui", false, "Skip gbtgui test")
	flag.BoolVar(&noI9w, "no-i9w", false, "Skip i9w test")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if !noGbt {
		runGbt(dir)
	}

	if !noGbtgui {
		runGbtgui(dir)
	}

	if !noI9w {
		runI9w(dir)
	}
}

func runGbt(rootDir string) {
	heading(2, "gbt")
	heading(3, "Building and running")

	gbtDir := filepath.Join(rootDir, "gbt")

	if err := runCommand(gbtDir, "go", "build"); err != nil {
		log.Fatal(err)
	}

	if err := runCommand(gbtDir, "./gbt"); err != nil {
		log.Fatal(err)
	}
}

func runGbtgui(rootDir string) {
	heading(2, "gbtgui")
	heading(3, "Building")

	gbtguiDir := filepath.Join(rootDir, "gbtgui")

	if err := runCommand(gbtguiDir, "go", "build"); err != nil {
		log.Fatal(err)
	}
}

func runI9w(rootDir string) {
	heading(2, "i9w")
	heading(3, "Building libtai")

	libtaiDir := filepath.Join(rootDir, "i9w", "vendor", "libtai")

	if err := runCommand(libtaiDir, "make"); err != nil {
		log.Fatal(err)
	}

	heading(3, `Installing "leapsecs.dat"`)

	if err := runCommand(libtaiDir, "sudo", "install", "-d", leapsecsInstallPath); err != nil {
		log.Fatalf("failed to create directory for leapsecs.dat: %v", err)
	}

	if err := runCommand(libtaiDir, "sudo", "install", "-m", "0644", "leapsecs.dat", leapsecsInstallPath); err != nil {
		log.Fatalf("failed to install leapsecs.dat: %v", err)
	}

	heading(3, `Building and running "beattai"`)
	i9wDir := filepath.Join(rootDir, "i9w")

	if err := runCommand(i9wDir, "make"); err != nil {
		log.Fatal(err)
	}

	if err := runCommand(i9wDir, "./beattai"); err != nil {
		log.Fatal(err)
	}
}

func heading(level int, message string) {
	if level < prevHeadingLevel {
		fmt.Println()
	}
	switch level {
	case 2:
		fmt.Print("=== ")
	case 3:
		fmt.Print("- ")
	}
	fmt.Println(message)

	prevHeadingLevel = level
}

func runCommand(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

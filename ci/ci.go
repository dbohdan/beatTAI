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
		fmt.Println("=== Building and running gbt")

		gbtDir := filepath.Join(dir, "gbt")
		if err := os.Chdir(gbtDir); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("go", "build"); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("./gbt"); err != nil {
			log.Fatal(err)
		}
	}

	if !noGbtgui {
		fmt.Println("=== Building gbtgui")

		gbtguiDir := filepath.Join(dir, "gbtgui")
		if err := os.Chdir(gbtguiDir); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("go", "build"); err != nil {
			log.Fatal(err)
		}
	}

	if !noI9w {
		fmt.Println("=== Building libtai and installing 'leapsecs.dat'")

		libtaiDir := filepath.Join(dir, "i9w", "vendor", "libtai")
		if err := os.Chdir(libtaiDir); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("make"); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("sudo", "install", "-d", leapsecsInstallPath); err != nil {
			log.Fatalf("failed to create directory for leapsecs.dat: %v", err)
		}

		if err := runCommand("sudo", "install", "-m", "0644", "leapsecs.dat", leapsecsInstallPath); err != nil {
			log.Fatalf("failed to install leapsecs.dat: %v", err)
		}

		fmt.Println("=== Building and running i9w")
		i9wDir := filepath.Join(dir, "i9w")
		if err := os.Chdir(i9wDir); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("make"); err != nil {
			log.Fatal(err)
		}

		if err := runCommand("./beattai"); err != nil {
			log.Fatal(err)
		}
	}
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

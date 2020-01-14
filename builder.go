package gowtk

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

// The Builder takes a gowtk project directory and compiles it into a web project. It expects a working go
// toolchain in the current path environment. It sports a small build cache, which avoids building and embedding
// files if nothing changed. It copies everything from the static folder into the build directory.
type Builder struct {
	mutex            sync.Mutex
	projectDir       string
	outDir           string
	currentBuildHash string
}

// NewBuilder creates a new Builder instance with obligatory settings
func NewBuilder(projectDir string, outDir string) *Builder {
	return &Builder{projectDir: projectDir, outDir: outDir}
}

// goRoot returns the GOROOT folder
func goRoot() (string, error) {
	cmd := exec.Command("go", "env", "GOROOT")
	cmd.Env = os.Environ()
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get goroot: %w", err)
	}
	return string(stdoutStderr), nil
}

// Resets the build hash to force a rebuild on next Build call
func (b *Builder) Reset() {
	b.currentBuildHash = ""
}

// Build triggers a new build
func (b *Builder) Build() error {
	files, err := listFiles(b.projectDir)
	if err != nil {
		return fmt.Errorf("failed to list files %s: %w", b.projectDir, err)
	}

	hash, err := calculateHash(files)
	if err != nil {
		return fmt.Errorf("failed to calculate hash %s: %w", b.projectDir, err)
	}

	if hash == b.currentBuildHash {
		return nil
	}

	// clear build folder
	err = os.RemoveAll(b.outDir)
	if err != nil {
		return fmt.Errorf("failed to clean build dir %s: %w", b.outDir, err)
	}

	err = os.MkdirAll(b.outDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create build dir %s: %w", b.outDir, err)
	}

	// perform wasm build
	cmd := exec.Command("go", "build", "-o", filepath.Join(b.outDir, "app.wasm"))
	cmd.Env = append(os.Environ(),
		"GOOS=js",
		"GOARCH=wasm", // this value is used
	)
	cmd.Dir = b.projectDir
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to compile webassembly %s: %w", b.projectDir, err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	// copy go-compiler and specific wasm adapter script for our app.wasm
	goRoot, err := goRoot()
	if err != nil {
		return fmt.Errorf("failed to grab goRoot: %w", err)
	}
	_, err = copy(filepath.Join(goRoot, "misc/wasm/wasm_exec.js"), filepath.Join(b.outDir, "wasm_exec.js"))

	// copy our entire static folder
	staticFrom := filepath.Join(b.projectDir, "static")
	staticTo := filepath.Join(b.outDir, "static")
	err = copyDir(staticFrom, staticTo)
	if err != nil {
		return fmt.Errorf("failed to copy static folder from %s -> %s: %w", staticFrom, staticTo, err)
	}

	return nil
}

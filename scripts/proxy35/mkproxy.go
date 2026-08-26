package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	proxy := flag.String("proxy", "", "proxy root directory")
	version := flag.String("version", "v0.1.0", "module version")
	flag.Parse()
	if *proxy == "" || flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "usage: mkproxy -proxy DIR [-version v0.1.0] path=dir ...\n")
		os.Exit(2)
	}
	infoTime := time.Date(2026, 8, 26, 0, 0, 0, 0, time.UTC)
	for _, arg := range flag.Args() {
		path, dir, ok := strings.Cut(arg, "=")
		if !ok || path == "" || dir == "" {
			fmt.Fprintf(os.Stderr, "mkproxy: bad spec %q (want path=dir)\n", arg)
			os.Exit(2)
		}
		if err := publish(*proxy, path, *version, dir, infoTime); err != nil {
			fmt.Fprintf(os.Stderr, "mkproxy: %s@%s: %v\n", path, *version, err)
			os.Exit(1)
		}
		fmt.Printf("published %s@%s\n", path, *version)
	}
}

func publish(proxy, path, version, dir string, t time.Time) error {
	escaped, err := module.EscapePath(path)
	if err != nil {
		return err
	}
	vdir := filepath.Join(proxy, filepath.FromSlash(escaped), "@v")
	if err := os.MkdirAll(vdir, 0o755); err != nil {
		return err
	}
	modBytes, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(vdir, version+".mod"), modBytes, 0o644); err != nil {
		return err
	}
	info, err := json.Marshal(struct {
		Version string
		Time    time.Time
	}{Version: version, Time: t})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(vdir, version+".info"), append(info, '\n'), 0o644); err != nil {
		return err
	}
	zf, err := os.Create(filepath.Join(vdir, version+".zip"))
	if err != nil {
		return err
	}
	err = zip.CreateFromDir(zf, module.Version{Path: path, Version: version}, dir)
	closeErr := zf.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	listPath := filepath.Join(vdir, "list")
	existing, readErr := os.ReadFile(listPath)
	if readErr != nil && !os.IsNotExist(readErr) {
		return readErr
	}
	found := false
	for _, line := range strings.Split(string(existing), "\n") {
		if line == version {
			found = true
			break
		}
	}
	if found {
		return nil
	}
	f, err := os.OpenFile(listPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	_, err = io.WriteString(f, version+"\n")
	cerr := f.Close()
	if err != nil {
		return err
	}
	return cerr
}

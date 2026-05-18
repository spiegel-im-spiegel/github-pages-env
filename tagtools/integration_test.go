package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github-pages-env/tagtools/internal/tagslist"
	"github-pages-env/tagtools/internal/toptags"
)

func TestTagslistGolden(t *testing.T) {
	out := filepath.Join(t.TempDir(), "tagslist.csv")
	cfg := tagslist.DefaultConfig()
	cfg.ContentDir = "../content"
	cfg.Out = out
	cfg.InheritMeansFrom = "../.github/workflows/tagslist.csv"

	if err := tagslist.Run(cfg); err != nil {
		t.Fatalf("tagslist.Run() error = %v", err)
	}

	expected, err := os.ReadFile("testdata/golden/tagslist.csv")
	if err != nil {
		t.Fatalf("read expected: %v", err)
	}
	actual, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read actual: %v", err)
	}
	if !bytes.Equal(expected, actual) {
		lineNo, expectedLine, actualLine := firstDiffLine(expected, actual)
		t.Fatalf("tagslist output mismatch against golden\nline=%d\nexpected=%q\nactual=%q", lineNo, expectedLine, actualLine)
	}
}

func TestToptagsGolden(t *testing.T) {
	out := filepath.Join(t.TempDir(), "toptags.json")
	cfg := toptags.DefaultConfig()
	cfg.ContentDir = "../content"
	cfg.Out = out
	cfg.Today = "2026-05-18"

	if err := toptags.Run(cfg); err != nil {
		t.Fatalf("toptags.Run() error = %v", err)
	}

	expected, err := os.ReadFile("testdata/golden/toptags.json")
	if err != nil {
		t.Fatalf("read expected: %v", err)
	}
	actual, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read actual: %v", err)
	}
	if !bytes.Equal(expected, actual) {
		lineNo, expectedLine, actualLine := firstDiffLine(expected, actual)
		t.Fatalf("toptags output mismatch against golden\nline=%d\nexpected=%q\nactual=%q", lineNo, expectedLine, actualLine)
	}
}

func firstDiffLine(expected, actual []byte) (int, string, string) {
	eLines := strings.Split(string(expected), "\n")
	aLines := strings.Split(string(actual), "\n")

	maxLen := len(eLines)
	if len(aLines) > maxLen {
		maxLen = len(aLines)
	}

	for i := 0; i < maxLen; i++ {
		e := ""
		if i < len(eLines) {
			e = eLines[i]
		}
		a := ""
		if i < len(aLines) {
			a = aLines[i]
		}
		if e != a {
			return i + 1, trimLine(e), trimLine(a)
		}
	}

	return 0, "", ""
}

func trimLine(s string) string {
	if len(s) <= 160 {
		return s
	}
	return fmt.Sprintf("%s...", s[:160])
}

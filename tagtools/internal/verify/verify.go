package verify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/goark/errs"

	"github-pages-env/tagtools/internal/tagslist"
	"github-pages-env/tagtools/internal/toptags"
)

// Config represents verify command options.
type Config struct {
	ContentDir       string `pflag:"content-dir,c,content directory to scan"`
	ExpectedTagslist string `pflag:"expected-tagslist,,expected tagslist CSV path"`
	ExpectedToptags  string `pflag:"expected-toptags,,expected toptags JSON path"`
	InheritMeansFrom string `pflag:"inherit-means-from,m,CSV path to inherit means from"`
	TopN             int    `pflag:"top-n,n,number of top tags"`
	Today            string `pflag:"today,t,override today date (YYYY-MM-DD)"`
	Window           string `pflag:"window,w,window duration (currently only 1y)"`
	Debug            bool   `pflag:"debug,d,output verify result as JSON"`
}

func DefaultConfig() Config {
	return Config{
		ContentDir:       "content",
		ExpectedTagslist: ".github/workflows/tagslist.csv",
		ExpectedToptags:  "data/toptags.json",
		InheritMeansFrom: ".github/workflows/tagslist.csv",
		TopN:             15,
		Today:            "",
		Window:           "1y",
		Debug:            false,
	}
}

// Failure represents verification details with first mismatch information.
type Failure struct {
	Stage            string
	ExpectedPath     string
	ActualPath       string
	Line             int
	ExpectedLine     string
	ActualLine       string
	ReproduceCommand string
	Cause            error
}

func (f *Failure) Error() string {
	if f.Cause != nil {
		return fmt.Sprintf("%s verification failed: %v\nreproduce: %s", f.Stage, f.Cause, f.ReproduceCommand)
	}
	return fmt.Sprintf(
		"%s verification failed\n  expected: %s\n  actual:   %s\n  first diff line: %d\n  expected line: %q\n  actual line:   %q\nreproduce: %s",
		f.Stage,
		f.ExpectedPath,
		f.ActualPath,
		f.Line,
		f.ExpectedLine,
		f.ActualLine,
		f.ReproduceCommand,
	)
}

type mismatch struct {
	ExpectedPath string
	ActualPath   string
	Line         int
	ExpectedLine string
	ActualLine   string
}

func (m *mismatch) Error() string {
	return fmt.Sprintf(
		"output mismatch\n  expected: %s\n  actual:   %s\n  first diff line: %d\n  expected line: %q\n  actual line:   %q",
		m.ExpectedPath,
		m.ActualPath,
		m.Line,
		m.ExpectedLine,
		m.ActualLine,
	)
}

type debugPayload struct {
	Status           string `json:"status"`
	Stage            string `json:"stage,omitempty"`
	Message          string `json:"message,omitempty"`
	ExpectedPath     string `json:"expected_path,omitempty"`
	ActualPath       string `json:"actual_path,omitempty"`
	Line             int    `json:"line,omitempty"`
	ExpectedLine     string `json:"expected_line,omitempty"`
	ActualLine       string `json:"actual_line,omitempty"`
	ReproduceCommand string `json:"reproduce_command"`
}

// JSONSuccess returns JSON output for successful verification.
func JSONSuccess(cfg Config) (string, error) {
	b, err := json.Marshal(debugPayload{Status: "ok", ReproduceCommand: reproduceCommand(cfg)})
	if err != nil {
		return "", errs.Wrap(err)
	}
	return string(b), nil
}

// JSONError returns JSON output for failed verification.
func JSONError(err error, cfg Config) (string, error) {
	var f *Failure
	if !errors.As(err, &f) {
		f = toFailure(err, "verify", reproduceCommand(cfg))
	}
	p := debugPayload{
		Status:           "error",
		Stage:            f.Stage,
		Message:          failureMessage(f),
		ExpectedPath:     f.ExpectedPath,
		ActualPath:       f.ActualPath,
		Line:             f.Line,
		ExpectedLine:     f.ExpectedLine,
		ActualLine:       f.ActualLine,
		ReproduceCommand: f.ReproduceCommand,
	}
	b, jerr := json.Marshal(p)
	if jerr != nil {
		return "", errs.Wrap(jerr)
	}
	return string(b), nil
}

func failureMessage(f *Failure) string {
	if f.Cause != nil {
		return fmt.Sprintf("%s verification failed: %v", f.Stage, f.Cause)
	}
	return fmt.Sprintf(
		"%s verification failed (first diff line: %d)",
		f.Stage,
		f.Line,
	)
}

// Run executes output verification against expected files.
func Run(cfg Config) error {
	repro := reproduceCommand(cfg)

	tmpDir, err := os.MkdirTemp("", "tagtools-verify-*")
	if err != nil {
		return errs.Wrap(err)
	}
	defer os.RemoveAll(tmpDir)

	tagslistOut := filepath.Join(tmpDir, "tagslist.csv")
	if err := tagslist.Run(tagslist.Config{
		ContentDir:       cfg.ContentDir,
		Out:              tagslistOut,
		InheritMeansFrom: cfg.InheritMeansFrom,
	}); err != nil {
		return &Failure{Stage: "tagslist", ReproduceCommand: repro, Cause: errs.Wrap(err)}
	}
	if err := compareFile(cfg.ExpectedTagslist, tagslistOut); err != nil {
		return toFailure(err, "tagslist", repro)
	}

	toptagsOut := filepath.Join(tmpDir, "toptags.json")
	if err := toptags.Run(toptags.Config{
		ContentDir: cfg.ContentDir,
		Out:        toptagsOut,
		TopN:       cfg.TopN,
		Today:      cfg.Today,
		Window:     cfg.Window,
	}); err != nil {
		return &Failure{Stage: "toptags", ReproduceCommand: repro, Cause: errs.Wrap(err)}
	}
	if err := compareFile(cfg.ExpectedToptags, toptagsOut); err != nil {
		return toFailure(err, "toptags", repro)
	}

	return nil
}

func compareFile(expectedPath, actualPath string) error {
	expected, err := os.ReadFile(expectedPath)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("path", expectedPath))
	}
	actual, err := os.ReadFile(actualPath)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("path", actualPath))
	}

	if bytes.Equal(expected, actual) {
		return nil
	}

	lineNo, expectedLine, actualLine := firstDiffLine(expected, actual)
	return &mismatch{
		ExpectedPath: expectedPath,
		ActualPath:   actualPath,
		Line:         lineNo,
		ExpectedLine: expectedLine,
		ActualLine:   actualLine,
	}
}

func toFailure(err error, stage, reproduce string) *Failure {
	var mm *mismatch
	if errors.As(err, &mm) {
		return &Failure{
			Stage:            stage,
			ExpectedPath:     mm.ExpectedPath,
			ActualPath:       mm.ActualPath,
			Line:             mm.Line,
			ExpectedLine:     mm.ExpectedLine,
			ActualLine:       mm.ActualLine,
			ReproduceCommand: reproduce,
		}
	}
	return &Failure{Stage: stage, ReproduceCommand: reproduce, Cause: err}
}

func reproduceCommand(cfg Config) string {
	parts := []string{
		"./bin/tagtools",
		"verify",
		"--content-dir", strconv.Quote(cfg.ContentDir),
		"--expected-tagslist", strconv.Quote(cfg.ExpectedTagslist),
		"--expected-toptags", strconv.Quote(cfg.ExpectedToptags),
		"--inherit-means-from", strconv.Quote(cfg.InheritMeansFrom),
		"--top-n", strconv.Itoa(cfg.TopN),
		"--today", strconv.Quote(cfg.Today),
		"--window", strconv.Quote(cfg.Window),
	}
	if cfg.Debug {
		parts = append(parts, "--debug")
	}
	return strings.Join(parts, " ")
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

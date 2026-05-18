package frontmatter

import (
	"bufio"
	"os"
	"regexp"

	"github.com/goark/errs"
)

var (
	reDateLine = regexp.MustCompile(`^date\s*=`)
	reDate     = regexp.MustCompile(`"([0-9]{4}-[0-9]{2}-[0-9]{2})`)
	reTagsLine = regexp.MustCompile(`^tags\s*=`)
	reQuoted   = regexp.MustCompile(`"[^"]+"`)
)

// Meta holds front matter fields used by tagtools.
type Meta struct {
	Date string
	Tags []string
}

// ParseFile parses TOML front matter and extracts only date and tags fields.
func ParseFile(path string) (Meta, error) {
	f, err := os.Open(path)
	if err != nil {
		return Meta{}, errs.Wrap(err, errs.WithContext("path", path))
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	meta := Meta{}
	fmCount := 0
	inTags := false

	for s.Scan() {
		line := s.Text()

		if line == "+++" {
			fmCount++
			if fmCount >= 2 {
				break
			}
			continue
		}
		if fmCount != 1 {
			continue
		}

		if reDateLine.MatchString(line) {
			m := reDate.FindStringSubmatch(line)
			if len(m) > 1 {
				meta.Date = m[1]
			}
		}

		if inTags {
			meta.Tags = append(meta.Tags, extractQuoted(line)...)
			if hasArrayEnd(line) {
				inTags = false
			}
			continue
		}

		if reTagsLine.MatchString(line) {
			meta.Tags = append(meta.Tags, extractQuoted(line)...)
			if !hasArrayEnd(line) {
				inTags = true
			}
		}
	}
	if err := s.Err(); err != nil {
		return Meta{}, errs.Wrap(err, errs.WithContext("path", path))
	}

	return meta, nil
}

func extractQuoted(line string) []string {
	matches := reQuoted.FindAllString(line, -1)
	res := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) >= 2 {
			res = append(res, m[1:len(m)-1])
		}
	}
	return res
}

func hasArrayEnd(line string) bool {
	for i := 0; i < len(line); i++ {
		if line[i] == ']' {
			return true
		}
	}
	return false
}

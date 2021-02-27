// +build run

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func getPaths(root, path string) []string {
	path = filepath.ToSlash(path)
	subPath := ""
	if strings.HasPrefix(path, "**/") {
		roots := []string{}
		if len(root) == 0 {
			roots = getDirectories("./")
		} else {
			roots = getDirectories(root)
		}
		subPath = path[3:]
		if len(roots) > 0 {
			paths := []string{}
			for _, rt := range roots {
				ps := getPaths(rt, subPath)
				if len(ps) > 0 {
					paths = append(paths, ps...)
				}
			}
			return paths
		} else {
			path = subPath
		}
	} else if i := strings.Index(path, "/**/"); i >= 0 {
		return getPaths(root+path[:i+1], path[i+1:])
	}

	if len(path) > 0 {
		paths := []string{}
		if ps, err := filepath.Glob(root + path); err == nil {
			for _, p := range ps {
				if info, err := os.Stat(p); err == nil {
					if m := info.Mode(); (m & os.ModeDir) != 0 {
						paths = append(paths, filepath.Clean(p)+string(filepath.Separator))
					} else {
						paths = append(paths, p)
					}
				}
			}
		}
		return paths
	}
	return []string{filepath.Clean(root) + string(filepath.Separator)}
}

func getDirectories(root string) []string {
	paths := &[]string{}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			path = filepath.ToSlash(path)
			*paths = append(*paths, filepath.ToSlash(filepath.Clean(path))+"/")
		}
		return nil
	})
	return *paths
}

func removeDuplicate(ary []string) []string {
	m := make(map[string]struct{})
	for _, e := range ary {
		m[e] = struct{}{}
	}
	u := make([]string, 0, len(ary))
	for i := range m {
		u = append(u, i)
	}
	sort.Strings(u)
	return u
}

func main() {
	paths := getPaths("", os.Args[1])
	paths = removeDuplicate(paths)
	for _, p := range paths {
		if strings.HasSuffix(p, string(filepath.Separator)) {
			fmt.Println("Dir: ", p)
		} else {
			fmt.Println("File: ", p)
		}
	}
}

// binlog parses a MySQL/MariaDB binary log and prints a log of queries executed.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	minLen = 6
	maxLen = 256
)

type Query struct {
	Query    string
	Database string
}

// Parse the binlog!
func main() {
	strs, err := GetStrings("binlog.000003")
	if err != nil {
		log.Fatal(err)
	}
	queries := ParseBinlogs(strs)
	for _, q := range queries {
		fmt.Printf("executed '%s' on database '%s'\n", q.Query, q.Database)
	}
}

func ParseBinlogs(strs []string) []Query {
	queries := make([]Query, 0)

	//header := strs[0] // 10.5.9-MariaDB-1:10.5.9+maria-focal-log
	//filename := strs[1] // binlog.000003

	for i := 2; i < len(strs); i += 2 {
		queries = append(queries, Query{Query: strs[i+1], Database: strs[i]})
	}

	return queries
}

// TrimTrailingJunk removes some binlog-specific trailing characters from strings.
func TrimTrailingJunk(in string) string {
	if strings.HasSuffix(in, "^") {
		return in[:len(in)-1]
	}
	if strings.HasSuffix(in, "NxDF") {
		return in[:len(in)-4]
	}
	return in
}

// GetStrings is a Go implementation of 'strings' from 'binutils'.
// Shoutout to github.com/robpike/strings.
func GetStrings(filename string) ([]string, error) {
	results := make([]string, 0)

	fd, err := os.Open(filename)
	if err != nil {
		return results, err
	}

	var (
		buf   = bufio.NewReader(fd)
		runes = make([]rune, 0)
		pos   = int64(0)
	)

	for {
		var (
			r   rune
			wid int
			err error
		)

		for ; ; pos += int64(wid) {
			// Read one rune (character) at a time.
			r, wid, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return results, err // actual read error
				}
				return results, nil
			}

			// If we encounter a non-ASCII character, save what we read and move on.
			if !strconv.IsPrint(r) || r > 0xFF {
				if len(runes) >= minLen {
					results = append(results, TrimTrailingJunk(string(runes)))
				}
				runes = runes[0:0] // Reset the buffer.
				continue
			}

			runes = append(runes, r) // Save the printable rune to the buffer.
		}
	}
}

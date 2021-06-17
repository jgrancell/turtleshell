package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jgrancell/turtleshell/cli"
)

func History(s *cli.Cli, components []string) (int, error) {
	if len(components) >= 1 {
		lookup := strings.Join(components, " ")
		history := s.History.Find(lookup)

		for _, entry := range history {
			fmt.Println(strconv.FormatInt(entry.Timestamp, 10)+":", entry.Value)
		}
		return 0, nil
	}

	fmt.Println("For complete shell history see the history file at", s.Configuration.HistoryFile)
	var limit int
	if len(s.History.Entries) < s.Configuration.HistoryLimit {
		limit = len(s.History.Entries)
	} else {
		limit = s.Configuration.HistoryLimit
	}

	history := s.History.Entries
	for i := 0; i < limit; i++ {
		entry := history[i]
		fmt.Println(strconv.FormatInt(entry.Timestamp, 10)+":", entry.Value)
	}
	return 0, nil
}

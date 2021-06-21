package history

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jgrancell/turtleshell/utils"
)

type History struct {
	Entries []HistoryEntry
	File    string
}

type HistoryEntry struct {
	Timestamp int64
	Value     string
}

func Load(path string) (*History, error) {
	var err error
	if err = utils.ValidateOrCreateFile(path); err != nil {
		return nil, err
	}

	h := &History{
		Entries: make([]HistoryEntry, 0),
		File:    path,
	}
	return h, nil
}

func (h *History) Append(text string) error {
	newEntry := HistoryEntry{
		Timestamp: time.Now().Unix(),
		Value:     text,
	}

	file, err := os.OpenFile(h.File, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	if _, err := file.WriteString("# Timestamp: " + strconv.FormatInt(newEntry.Timestamp, 10) + "\n"); err != nil {
		return err
	}
	if _, err := file.WriteString(newEntry.Value + "\n"); err != nil {
		return err
	}
	h.Entries = append(h.Entries, newEntry)
	return nil
}

func (h *History) Find(text string) []HistoryEntry {
	var entries []HistoryEntry
	for _, hist := range h.Entries {
		if strings.Contains(hist.Value, text) {
			entries = append(entries, hist)
		}
	}
	return entries
}

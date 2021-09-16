package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/coreos/go-systemd/v22/sdjournal"
	"github.com/stretchr/testify/assert"
)

func TestJournalGetEntryFromSpecificFile(t *testing.T) {

	//GIVEN
	j, _ := sdjournal.NewJournalFromFiles("data/logEntries.journal")

	//WHEN
	j.SeekHead()
	j.Next()
	entry, _ := j.GetEntry()

	//THEN
	assert.NotNil(t, entry, "Entry must be received")
	assert.NotNil(t, entry.Cursor, "Cursor has to be defined")
	assert.Equal(t, "My Message", entry.Fields[sdjournal.SD_JOURNAL_FIELD_MESSAGE], "My Message")
	//4 more entries should be found
	var result uint64
	result, _ = j.Next()
	assert.Equal(t, uint64(1), result, "Next should find result")
	result, _ = j.Next()
	assert.Equal(t, uint64(1), result, "Next should find result")
	result, _ = j.Next()
	assert.Equal(t, uint64(1), result, "Next should NOT find result")
	result, _ = j.Next()
	assert.Equal(t, uint64(0), result, "Next should NOT find result")

}

func writeSomeTestEntries() {

	var entries []sdjournal.JournalEntry

	entries = append(entries, sdjournal.JournalEntry{Cursor: "1",
		MonotonicTimestamp: 1321,
		RealtimeTimestamp:  12345,
		Fields:             map[string]string{sdjournal.SD_JOURNAL_FIELD_MESSAGE: "My Message"}})
	entries = append(entries, sdjournal.JournalEntry{Cursor: "2",
		MonotonicTimestamp: 1321,
		RealtimeTimestamp:  12345,
		Fields:             map[string]string{sdjournal.SD_JOURNAL_FIELD_MESSAGE: "My Message"}})
	entries = append(entries, sdjournal.JournalEntry{Cursor: "3",
		MonotonicTimestamp: 1321,
		RealtimeTimestamp:  12345,
		Fields:             map[string]string{sdjournal.SD_JOURNAL_FIELD_MESSAGE: "My Message"}})
	entries = append(entries, sdjournal.JournalEntry{Cursor: "4",
		MonotonicTimestamp: 1321,
		RealtimeTimestamp:  12345,
		Fields:             map[string]string{sdjournal.SD_JOURNAL_FIELD_MESSAGE: "My Message"}})

	writeJournalEntriesToTxt(entries)
}

//lib/systemd/systemd-journal-remote --output=/home/ammo/git/journald-go-example/data/logEntries1.journal ~/git/journald-go-example/data/logEntries.txt

func writeJournalEntriesToTxt(entries []sdjournal.JournalEntry) {
	f, err := os.Create("data/logEntries.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, entry := range entries {
		for key, value := range entry.Fields {
			_, err := fmt.Fprint(f, key, "=", value, "\n")
			if err != nil {

				log.Fatal(err)
			}
		}
		fmt.Fprint(f, sdjournal.SD_JOURNAL_FIELD_CURSOR, "=", entry.Cursor, "\n")
		fmt.Fprint(f, sdjournal.SD_JOURNAL_FIELD_MONOTONIC_TIMESTAMP, "=", entry.MonotonicTimestamp, "\n")
		fmt.Fprint(f, sdjournal.SD_JOURNAL_FIELD_REALTIME_TIMESTAMP, "=", entry.RealtimeTimestamp, "\n")
		fmt.Fprint(f, "unconfined", "\n")
		fmt.Fprintln(f)

	}

}

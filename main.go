package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/coreos/go-systemd/v22/sdjournal"
	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
)

type LogEntry struct {
	Message  string
	Severity string
}

func main() {
	fmt.Print("asdasdsad")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeWesteros), viewer.WithAddr("localhost:8087"))

	mgr := statsview.New()

	// Start() runs a HTTP server at `localhost:18066` by default.
	go mgr.Start()

	//  getJournalDLogs(1000000, 4)
	// entries, _ := GetJournalDLogs(10000, 4)
	GetJournalDLogs(100000, 4)
	// for _, entry := range entries {

	// fmt.Println("--------------------------------------")
	// fmt.Println("Message:", entry.Message)
	// fmt.Println("Severity:", entry.Severity)
	// fmt.Println("--------------------------------------")
	// }
	runtime.GC()
	fmt.Println("Time:", time.Now().UnixMilli())
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}

func GetJournalDLogs(numberOfEntries int, priority int) ([]LogEntry, error) {
	var journal, _ = sdjournal.NewJournal()

	var entries = make([]LogEntry, 0)
	if err := journal.SeekTail(); err != nil {
		return entries, err
	}
	if _, err := journal.Previous(); err != nil {
		return entries, err
	}

	readEntries := 0
	for readEntries < numberOfEntries {
		r, err := journal.Previous()
		if err != nil {
			return entries, err
		}
		if r == 0 {
			fmt.Println("Reached the end")
			break
		}
		entry, err := journal.GetEntry()
		if err != nil {
			return entries, err
		}

		severity, _ := strconv.Atoi((*entry).Fields[sdjournal.SD_JOURNAL_FIELD_PRIORITY])

		if severity == priority {
			var logEntry LogEntry = LogEntry{Message: (*entry).Fields[sdjournal.SD_JOURNAL_FIELD_MESSAGE], Severity: (*entry).Fields[sdjournal.SD_JOURNAL_FIELD_PRIORITY]}
			entries = append(entries, logEntry)
			readEntries++
		}
	}
	return entries, nil
}

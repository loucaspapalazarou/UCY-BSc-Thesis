package tools

import (
	"fmt"
	"os"
	"time"
)

var (
	TOTAL_ADD_ATOMIC_TIME int
	REQUESTS              int
)

func saveState(server_id string) error {

	filename := "scenario_results_" + server_id + ".txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	avg_add_atomic := float64(TOTAL_ADD_ATOMIC_TIME) / float64(REQUESTS) / float64(time.Millisecond)
	_, err = fmt.Fprintf(file, "REQUESTS=%d\nAVG_ADD_ATOMIC_TIME=%fms\n", REQUESTS, avg_add_atomic)

	if err != nil {
		return err
	}

	return nil
}

func IncrementAddAtomicTime(server_id string, t time.Duration) (int, int) {
	TOTAL_ADD_ATOMIC_TIME += int(t.Nanoseconds())
	REQUESTS++
	saveState(server_id)
	return TOTAL_ADD_ATOMIC_TIME, REQUESTS
}

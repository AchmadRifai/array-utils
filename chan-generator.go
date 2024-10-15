package arrayutils

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

func Int64sChanClosed(start, end int64) <-chan int64 {
	return IteratingChan(start, func(v int64) bool { return v <= end }, func(v int64) int64 { return v + 1 })
}

func Int64sChan(start, end int64) <-chan int64 {
	return IteratingChan(start, func(v int64) bool { return v < end }, func(v int64) int64 { return v + 1 })
}

func Int32sChanClosed(start, end int32) <-chan int32 {
	return IteratingChan(start, func(v int32) bool { return v <= end }, func(v int32) int32 { return v + 1 })
}

func Int32sChan(start, end int32) <-chan int32 {
	return IteratingChan(start, func(v int32) bool { return v < end }, func(v int32) int32 { return v + 1 })
}

func Int16sChanClosed(start, end int16) <-chan int16 {
	return IteratingChan(start, func(v int16) bool { return v <= end }, func(v int16) int16 { return v + 1 })
}

func Int16sChan(start, end int16) <-chan int16 {
	return IteratingChan(start, func(v int16) bool { return v < end }, func(v int16) int16 { return v + 1 })
}

func IntsChanClosed(start, end int) <-chan int {
	return IteratingChan(start, func(v int) bool { return v <= end }, func(v int) int { return v + 1 })
}

func IntsChan(start, end int) <-chan int {
	return IteratingChan(start, func(v int) bool { return v < end }, func(v int) int { return v + 1 })
}

type DataWithError[V interface{}] struct {
	Data V
	Err  error
}

func SqlToChan[V interface{}](row *sql.Rows, convert func(row *sql.Rows, sqlStep int) (V, error)) <-chan DataWithError[V] {
	dataChan := make(chan DataWithError[V])
	go func(row *sql.Rows) {
		defer close(dataChan)
		sqlStep := 0
		for row.NextResultSet() {
			sqlStep = sqlStep + 1
			if err := row.Err(); err != nil {
				dataChan <- DataWithError[V]{Err: err}
				continue
			}
			for row.Next() {
				v, err := convert(row, sqlStep)
				if err != nil {
					dataChan <- DataWithError[V]{Err: err}
				} else {
					dataChan <- DataWithError[V]{Data: v}
				}
			}
		}
	}(row)
	return dataChan
}

func IteratingChan[V interface{}](initial V, hasNext func(v V) bool, next func(v V) V) <-chan V {
	result := make(chan V)
	go func(initial V) {
		defer close(result)
		for position := initial; hasNext(position); position = next(position) {
			result <- position
		}
	}(initial)
	return result
}

func CsvToChan(csvPath string) <-chan string {
	strChan := make(chan string)
	go func(csvPath string) {
		file, err := os.Open(csvPath)
		defer func(file *os.File) {
			close(strChan)
			if r := recover(); r != nil {
				log.Println("", r)
				log.Println("", string(debug.Stack()))
			}
			file.Close()
		}(file)
		if err != nil {
			panic(err)
		}
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}
		for _, record := range records {
			strChan <- strings.Join(record, ", ")
		}
	}(csvPath)
	return strChan
}

package logging

import (
	"log"
	"time"
)

func logQuery(opeartion string, query interface{}, elapsed time.Duration, err error) {
	log.Printf("Operation: %s, Query: %v, Elapsed: %s, Error: %v", opeartion, query, elapsed, err)
}
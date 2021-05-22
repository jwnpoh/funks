// Package funks implements basic functions that I use frequently to save time on replicating code across different projects.
package funks

import (
	"encoding/csv"
	"fmt"
	"os"
    "io"
)

// ReadCSV reads a CSV file of filename f and returns a slice of records, with each record containing a slice of string. Returns [][]string and error.
func ReadCSV(f string) ([][]string, error) {
	file, err := os.Open(f)
	if err != nil {
        return nil, fmt.Errorf("error opening file %s - %v", f, err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	recs := make([][]string, 0, 50)
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
            return nil, fmt.Errorf("unable to read records in %s: %v", f, err)
		}
		recs = append(recs, rec)
	}
	return recs, err
}

// WriteCSV takes a slice of records recs, with each record containing a slice of string and writes to a CSV file filename. It returns an error if the write is unsuccessful. 
func WriteCSV(recs [][]string, filename string) error {
    out, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("error creating file %s - %v", filename, err)
    }

    w := csv.NewWriter(out)
    defer out.Close()

	for _, k := range recs {
		err = w.Write(k)
		if err != nil {
        return fmt.Errorf("unable to write records to %s - %v", filename, err)
		}
		w.Flush()
	}
	return nil
}

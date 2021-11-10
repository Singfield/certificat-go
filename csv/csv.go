package csv

import (
	"certificat-go/cert"
	"encoding/csv"
	"os"
)

func ParseCSV(fileName string) ([]*cert.Cert, error) {
	certs := make([]*cert.Cert, 0)
	f, err := os.Open(fileName)
	if err != nil {
		return certs, err
	}
	defer f.Close()
	r :=csv.NewReader(f)
	records, err := r.ReadAll()
	if err !=nil {
		return certs, err
	}

	for _, rec := range records {
		course := rec[0]
		name := rec[1]
		date := rec[2]
		c, err :=cert.New(course, name, date)
		if err !=nil {
			return certs, err
		}
		certs = append(certs, c)
	}
	return certs, nil
}
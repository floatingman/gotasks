package main

import (
	"encoding/csv"
	"os"
)

type RepositoryInterface interface {
	save(task Task) error
}

type TasksCsvRepository struct {
	Path string
}

func (csvRepository TasksCsvRepository) save(task Task) error {
	csvFile, err := os.OpenFile(csvRepository.Path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	err = writer.Write(task.toArrayString())
	writer.Flush()
	return err
}

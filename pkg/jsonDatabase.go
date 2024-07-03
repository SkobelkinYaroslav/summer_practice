package database

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"summer_practice/internal/domain"
	"sync"
)

type DataBase struct {
	fileName  string
	rows      []domain.Car
	lastIndex int
	mu        sync.Mutex
}

func New(fileName string) *DataBase {
	var fileRows []domain.Car
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return &DataBase{
			fileName:  fileName,
			rows:      nil,
			lastIndex: 0,
			mu:        sync.Mutex{},
		}
	}

	if err := json.Unmarshal(bytes, &fileRows); err != nil {
		return &DataBase{
			fileName:  fileName,
			rows:      nil,
			lastIndex: 0,
			mu:        sync.Mutex{},
		}
	}

	return &DataBase{
		fileName:  fileName,
		rows:      fileRows,
		lastIndex: fileRows[len(fileRows)-1].ID + 1,
		mu:        sync.Mutex{},
	}
}

func (db *DataBase) AddRow(row domain.Car) {
	db.mu.Lock()
	defer db.mu.Unlock()
	row.ID = db.lastIndex
	db.lastIndex++
	db.rows = append(db.rows, row)
	db.saveToFile(db.fileName)
}

func (db *DataBase) GetRow(id int) (domain.Car, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		return db.rows[index], nil
	}

	return domain.Car{}, domain.ErrNotFound
}

func (db *DataBase) UpdateRow(id int, newRow domain.Car) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		db.rows[index] = newRow
		db.saveToFile(db.fileName)
		return nil
	}

	return domain.ErrNotFound
}
func (db *DataBase) DeleteRow(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		db.rows = append(db.rows[:index], db.rows[index+1:]...)
		db.saveToFile(db.fileName)
		return nil
	}

	return domain.ErrNotFound
}

func (db *DataBase) saveToFile(fileName string) {
	data, err := json.MarshalIndent(db.rows, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
		return
	}

	if err = os.WriteFile(fileName, data, 0644); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

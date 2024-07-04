package database

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"summer_practice/internal/domain"
)

type operationType int

const (
	addOperation operationType = iota
	updateOperation
	deleteOperation
	readOperation
	readAllOperation
)

type operation struct {
	Type   operationType
	Row    domain.Car
	ID     int
	Result chan<- interface{} // Channel to send back read result
}

type DataBase struct {
	fileName  string
	rows      []domain.Car
	lastIndex int
	queue     chan operation
}

func New(fileName string) *DataBase {
	var fileRows []domain.Car
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fileRows = nil
	} else {
		err = json.Unmarshal(bytes, &fileRows)
		if err != nil {
			fileRows = nil
		}
	}

	var lastIndex int
	if fileRows != nil && len(fileRows) > 0 {
		lastIndex = fileRows[len(fileRows)-1].ID + 1
	}

	db := &DataBase{
		fileName:  fileName,
		rows:      fileRows,
		lastIndex: lastIndex,
		queue:     make(chan operation, 100),
	}

	go db.processQueue()
	return db
}

func (db *DataBase) AddRow(row domain.Car) {
	db.queue <- operation{Type: addOperation, Row: row}
}

func (db *DataBase) UpdateRow(id int, newRow domain.Car) {
	db.queue <- operation{Type: updateOperation, Row: newRow, ID: id}
}

func (db *DataBase) DeleteRow(id int) {
	db.queue <- operation{Type: deleteOperation, ID: id}
}

func (db *DataBase) GetRow(id int) (domain.Car, error) {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue <- operation{Type: readOperation, ID: id, Result: resultChan}

	result := <-resultChan
	switch res := result.(type) {
	case domain.Car:
		return res, nil
	case error:
		return domain.Car{}, res
	default:
		return domain.Car{}, domain.ErrInternalServerError
	}
}

func (db *DataBase) GetAllRows() ([]domain.Car, error) {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue <- operation{Type: readAllOperation, Result: resultChan}
	result := <-resultChan

	switch res := result.(type) {
	case []domain.Car:
		return res, nil
	case error:
		return nil, res
	default:
		return nil, domain.ErrInternalServerError
	}
}

func (db *DataBase) processQueue() {
	log.Println("processQueue started")
	for op := range db.queue {
		log.Println("processQueue hit")
		switch op.Type {
		case addOperation:
			db.addRow(op.Row)
		case updateOperation:
			db.updateRow(op.ID, op.Row)
		case deleteOperation:
			db.deleteRow(op.ID)
		case readOperation:
			db.readRow(op.ID, op.Result)
		case readAllOperation:
			db.readAllRows(op.Result)
		}
	}
}

func (db *DataBase) addRow(row domain.Car) {
	row.ID = db.lastIndex
	db.lastIndex++
	db.rows = append(db.rows, row)
	db.saveToFile(db.fileName)
}

func (db *DataBase) updateRow(id int, newRow domain.Car) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		db.rows[index] = newRow
		db.saveToFile(db.fileName)
	}
}

func (db *DataBase) deleteRow(id int) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		db.rows = append(db.rows[:index], db.rows[index+1:]...)
		db.saveToFile(db.fileName)
	}
}

func (db *DataBase) readRow(id int, result chan<- interface{}) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		result <- db.rows[index]
	} else {
		result <- domain.ErrNotFound
	}
}

func (db *DataBase) readAllRows(result chan<- interface{}) {
	if len(db.rows) == 0 {
		result <- domain.ErrNotFound
		return
	}
	result <- db.rows
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

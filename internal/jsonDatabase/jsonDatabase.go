package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"summer_practice/internal/domain"
)

type operationType int

const (
	addOperation operationType = iota
	putOperation
	deleteOperation
	readOperation
	readAllOperation
	updateFieldsOperation
)

type operation struct {
	Type           operationType
	Row            domain.Car
	ID             int
	FieldsToUpdate map[string]interface{}
	Result         chan<- interface{}
}

type DataBase struct {
	fileName  string
	rows      []domain.Car
	lastIndex int
	queue     []operation
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
		queue:     make([]operation, 0),
	}
	go db.processOperations()
	return db
}

func (db *DataBase) processOperations() {
	for {
		if len(db.queue) > 0 {
			op := db.queue[0]
			db.queue = db.queue[1:]
			switch op.Type {
			case addOperation:
				db.addRow(op.Row, op.Result)
			case putOperation:
				db.putRow(op.Row, op.Result)
			case deleteOperation:
				db.deleteRow(op.ID, op.Result)
			case readOperation:
				db.readRow(op.ID, op.Result)
			case readAllOperation:
				db.readAllRows(op.Result)
			case updateFieldsOperation:
				db.updateFieldsByID(op.ID, op.FieldsToUpdate, op.Result)
			}
		}
	}
}

func (db *DataBase) AddRow(row domain.Car) (domain.Car, error) {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue = append(db.queue, operation{Type: addOperation, Row: row, Result: resultChan})

	result := <-resultChan

	switch res := result.(type) {
	case domain.Car:
		return res, nil
	case error:
		return domain.Car{}, res
	default:
		fmt.Println(res)
		return domain.Car{}, domain.ErrInternalServerError
	}
}

func (db *DataBase) PutRow(newRow domain.Car) (domain.Car, error) {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue = append(db.queue, operation{Type: putOperation, Row: newRow, Result: resultChan})

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

func (db *DataBase) DeleteRow(id int) error {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue = append(db.queue, operation{Type: deleteOperation, ID: id, Result: resultChan})

	result := <-resultChan
	switch res := result.(type) {
	case nil:
		return nil
	case error:
		return res
	default:
		return domain.ErrInternalServerError
	}
}

func (db *DataBase) GetRow(id int) (domain.Car, error) {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue = append(db.queue, operation{Type: readOperation, ID: id, Result: resultChan})

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

	db.queue = append(db.queue, operation{Type: readAllOperation, Result: resultChan})

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

func (db *DataBase) UpdateRow(fieldsToUpdate map[string]interface{}) (domain.Car, error) {
	resultChan := make(chan interface{})
	defer close(resultChan)

	db.queue = append(db.queue, operation{Type: updateFieldsOperation, ID: fieldsToUpdate["id"].(int), FieldsToUpdate: fieldsToUpdate, Result: resultChan})

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

func (db *DataBase) addRow(row domain.Car, result chan<- interface{}) {
	row.ID = db.lastIndex
	db.lastIndex++
	db.rows = append(db.rows, row)
	if err := db.saveToFile(db.fileName); err != nil {
		result <- err
	}

	result <- row
}

func (db *DataBase) putRow(newRow domain.Car, result chan<- interface{}) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= newRow.ID
	})

	if index < len(db.rows) && db.rows[index].ID == newRow.ID {
		db.rows[index] = newRow
		if err := db.saveToFile(db.fileName); err != nil {
			result <- err
			return
		}
		result <- newRow
		return
	}
	result <- domain.ErrNotFound
}

func (db *DataBase) deleteRow(id int, result chan<- interface{}) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		db.rows = append(db.rows[:index], db.rows[index+1:]...)
		if err := db.saveToFile(db.fileName); err != nil {
			result <- err
			return
		}

		result <- nil
		return
	}

	result <- domain.ErrNotFound
}

func (db *DataBase) readRow(id int, result chan<- interface{}) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		result <- db.rows[index]
		return
	}
	result <- domain.ErrNotFound
}

func (db *DataBase) readAllRows(result chan<- interface{}) {
	if len(db.rows) == 0 {
		result <- domain.ErrNotFound
		return
	}
	result <- db.rows
}

func (db *DataBase) updateFieldsByID(id int, fieldsToUpdate map[string]interface{}, result chan<- interface{}) {
	index := sort.Search(len(db.rows), func(i int) bool {
		return db.rows[i].ID >= id
	})

	if index < len(db.rows) && db.rows[index].ID == id {
		row := &db.rows[index]

		for field, value := range fieldsToUpdate {
			switch field {
			case "brand":
				if v, ok := value.(string); ok {
					row.Brand = v
				}
			case "model":
				if v, ok := value.(string); ok {
					row.Model = v
				}
			case "mileage":
				if v, ok := value.(float64); ok {
					row.Mileage = int(v)
				}
			case "owners_count":
				if v, ok := value.(float64); ok {
					row.OwnersCount = int(v)
				}
			}
		}

		if err := db.saveToFile(db.fileName); err != nil {
			result <- err
			return
		}

		result <- *row
		return
	}

	result <- domain.ErrNotFound
}

func (db *DataBase) saveToFile(fileName string) error {
	data, err := json.MarshalIndent(db.rows, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(fileName, data, 0644); err != nil {
		return err
	}

	return nil
}

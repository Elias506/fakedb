package Fakedb

import (
	"sync"
	"strings"
	"fmt"
	_ "strconv"
	"Fakedb/helpers"
)

// in DB u cant use `"` in tableName

type DB struct {
	Tables map[string]*table
}

type table struct {
	sync.Mutex
	name string
	columnType map[string]string
	rows    []*rowType
}

func (t *table) Write() {
	for _, r := range t.rows {
		fmt.Println(r)
	}
}

type rowType map[string]interface{}

func (r *rowType) Write() {
	fmt.Println(r)
}

// CreateTable create a new table in DB
//
// For example:
// DB.CreateTable("tableName|colName1 = colType1, colName2 = colType2")

func (db *DB) CreateTable(s string) error {
	//...
	if db.Tables == nil {
		db.Tables = make(map[string]*table, 0)
	}
	var err error
	s, err = helpers.SpaceStringsBuilder(s)
	if err != nil {
		return err
	}
	globalSplit := strings.Split(s, "|")
	if len(globalSplit) != 2 {
		return fmt.Errorf(`incorrect input: "%s"`, s)
	}

	tableName := globalSplit[0]
	//if numberOfR(tableName, '"') > 0 {
	//	return fmt.Errorf(`expected banned symbol " in tableName: %v`, tableName)
	//}

	// Checking on unique table name
	if _, tableIsAlreadyExist := db.Tables[tableName]; tableIsAlreadyExist {
		return fmt.Errorf(`table "%s" already exist`, tableName)
	}
	couplesSplit := strings.Split(globalSplit[1], ",")

	t := &table{}
	t.rows = make([]*rowType, 0)
	t.columnType = make(map[string]string, 0)

	for _, couple := range couplesSplit {
		split := strings.Split(couple, "=")
		if len(split) != 2 || split[1] == "" {
			return fmt.Errorf(`incorrect params: "%s"`, couple)
		}
		if !helpers.IsType(split[1]) {
			return fmt.Errorf(`incorrect type: "%v"`, split[1])
		}
		if _, columnIsAlreadyExist := t.columnType[split[0]]; columnIsAlreadyExist {
			return fmt.Errorf(`repeated column name: "%s"`, split[1])
		}
		t.columnType[split[0]] = split[1]
	}
	db.Tables[tableName] = t
	return nil
}

func (db *DB) DeleteTable(tableName string) error {
	db.Tables[tableName].Lock()
	if _, tableIsExist := db.Tables[tableName]; !tableIsExist {
		return fmt.Errorf(`table "%v" does not exist`, tableName)
	}
	delete(db.Tables, tableName)
	return nil
}

//Insert data in table
//...
//For example:
//DB.Insert("tableName|colName1 = colValue1, colName2 = colValue2")
func (db *DB) Insert(s string) error {
	var err error
	s, err = helpers.SpaceStringsBuilder(s)
	if err != nil {
		return err
	}
	globalSplit := strings.Split(s, "|")
	if len(globalSplit) != 2 {
		return fmt.Errorf(`incorrect input: "%s"`, s)
	}
	//fmt.Println(tableName)
	t, tableIsExist := db.Tables[globalSplit[0]]
	if !tableIsExist {
		return fmt.Errorf(`table "%s" does not exist`, globalSplit[0])
	}
	t.Lock()
	defer t.Unlock()
	couplesSplit := strings.Split(globalSplit[1], ",")
	//rows := make([]interface{}, len(t.colName))
	newRow := make(rowType, len(couplesSplit))
	for _, couple := range couplesSplit {
		split := strings.Split(couple, "=")
		if len(split) != 2 || split[1] == "" {
			return fmt.Errorf(`incorrect input: "%s"`, s)
		}
		valueType, columnIsExist := t.columnType[split[0]]
		if !columnIsExist {
			return fmt.Errorf(`column "%s" does not exist`, split[0])
		}
		value, err := helpers.ParseType(split[1], valueType)
		if err != nil {
			return fmt.Errorf(`error ParseType: "%s"`, err)
		}
		newRow[split[0]] = value
	}
	t.rows = append(t.rows, &newRow)
	return nil
}

//func numberOfR(s string, r rune) int {
//	k := 0
//	for _, ch := range s {
//		if ch == r {
//			k++
//		}
//	}
//	return k
//}

// Select find and print wanted rowType in table
// DB.Select("tableName|colName1 = "sy", colName2 = 100")
func (db *DB) Select(s string) ([]*rowType, error) {

	var err error
	s, err = helpers.SpaceStringsBuilder(s)
	if err != nil {

		return nil, err
	}
	globalSplit := strings.Split(s, "|")
	if len(globalSplit) != 2 {
		return nil, fmt.Errorf(`incorrect input: "%s"`, s)
	}
	if globalSplit[1] == "" {
		return db.Tables[globalSplit[0]].rows, nil
	}
	//tableName := split[0]
	t, tableIsExist := db.Tables[globalSplit[0]]
	if !tableIsExist {
		return nil, fmt.Errorf(`table "%s" does not exist`, globalSplit)
	}
	t.Lock()
	defer t.Unlock()
	couplesSplit := strings.Split(globalSplit[1], ",")
	rows := make([]*rowType, 0)
	for _, row := range t.rows {
		isRowOk := true
		for _, couple := range couplesSplit {
			//Checking...
			split := strings.Split(couple, "=")
			if len(split) != 2 || split[1] == "" {
				return nil, fmt.Errorf(`incorrect input: "%s"`, s)
			}
			valueType, columnIsExist := t.columnType[split[0]]
			if !columnIsExist {
				return nil, fmt.Errorf(`column "%s" does not exist`, split[0])
			}
			value, err := helpers.ParseType(split[1], valueType)
			if err != nil {
				return nil, fmt.Errorf(`error ParseType: "%s"`, err)
			}

			//Working...
			if value != (*row)[split[0]] {
				isRowOk = false
				break
			}
		}
		if isRowOk {
			rows = append(rows, row)
		}
	}
	return rows, nil
}

//func main() {
//	DB := &DB{}
//	err := DB.CreateTable(`33|id = int, amount = string`)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(DB.Tables["33"])
//	err = DB.Insert(`33|id=?,amount="5"`, 1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(DB.Tables["33"])
//	err = DB.Insert(`33|id=2,amount="12"`)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(DB.Tables["33"])
//	s, err := DB.Select(`33|amount, id|id=?`, 1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(s)
//}


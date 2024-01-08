package database

import (
	"database/sql"
	"errors"
	"reflect"
)

func newSession(db *sql.DB) *Session {
	return &Session{db: db}
}

type Session struct {
	db *sql.DB
}

// Count 返回数量
func (s Session) Count(sql string, param ...any) (int, error) {
	var count int
	err := s.db.QueryRow(sql, param...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// ExecDDL 执行数据库变更语句
func (s Session) ExecDDL(sql string, param ...any) error {
	_, err := s.db.Exec(sql, param...)
	return err
}

func (s Session) Insert(sql string, param ...any) (sql.Result, error) {
	return s.db.Exec(sql, param...)
}

func (s Session) Select(sql string, param ...any) ([]map[string]interface{}, error) {
	rows, err := s.db.Query(sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []map[string]interface{}
	columns, err := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		err = rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}
		rowData := map[string]interface{}{}
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				rowData[colName] = string(b)
			} else {
				rowData[colName] = val
			}
		}
		list = append(list, rowData)
	}
	return list, nil
}

func (s Session) SelectOne(sql string, r interface{}, param ...any) (bool, error) {
	val := reflect.ValueOf(r)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return false, errors.New("Expected a pointer to a structs")
	}
	maps, err := s.Select(sql, param...)
	if err != nil {
		return false, err
	}
	if len(maps) == 0 {
		return false, nil
	}
	item := maps[0]
	sliceType := val.Type()
	for i := 0; i < sliceType.NumField(); i++ {
		field := sliceType.Field(i)
		mapVal := item[field.Name]
		val.FieldByName(field.Name).Set(reflect.ValueOf(mapVal))
	}
	return true, nil
}

func (s Session) SelectList(sql string, rList interface{}, param ...any) error {
	val := reflect.ValueOf(rList)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return errors.New("Expected a pointer to a slice")
	}
	sliceType := val.Elem().Type().Elem()
	if sliceType.Kind() != reflect.Struct {
		return errors.New("Expected a slice of structs")
	}

	maps, err := s.Select(sql, param...)
	if err != nil {
		return err
	}
	sliceVal := val.Elem()
	for _, item := range maps {
		instance := reflect.New(sliceType).Elem()
		for i := 0; i < sliceType.NumField(); i++ {
			field := sliceType.Field(i)
			mapVal := item[field.Name]
			instance.FieldByName(field.Name).Set(reflect.ValueOf(mapVal))
		}
		sliceVal.Set(reflect.Append(sliceVal, instance))
	}
	return nil
}

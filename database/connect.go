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

// Exec 执行数据库语句
func (s Session) Exec(sql string, param ...any) error {
	_, err := s.db.Exec(sql, param...)
	return err
}

func (s Session) Insert(sql string, param ...any) (sql.Result, error) {
	return s.db.Exec(sql, param...)
}

func (s Session) SelectListMap(sql string, param ...any) ([]map[string]interface{}, error) {
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

func (s Session) SelectOneForMap(sql string, param ...any) (map[string]interface{}, error) {
	list, err := s.SelectListMap(sql, param...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	} else {
		return nil, nil
	}
}

func rowScan(rows *sql.Rows, dest interface{}) error {
	destType := reflect.TypeOf(dest)
	if destType.Kind() != reflect.Ptr {
		return errors.New("Expected a pointer to a structs or slice")
	}
	if destType.Elem().Kind() == reflect.Slice {
		destValue := reflect.ValueOf(dest)
		elemType := destType.Elem()
		if elemType.Kind() != reflect.Struct {
			return errors.New("The slice element must be a structure")
		}
		numFields := elemType.NumField()
		for rows.Next() {
			structValue := reflect.New(elemType).Elem()
			fieldValues := make([]interface{}, numFields)
			for i := 0; i < numFields; i++ {
				fieldValues[i] = reflect.New(elemType.Field(i).Type).Interface()
			}
			err := rows.Scan(fieldValues...)
			if err != nil {
				return err
			}
			for i := 0; i < numFields; i++ {
				structValue.Field(i).Set(reflect.ValueOf(fieldValues[i]).Elem())
			}

			// 将结构体添加到切片中
			destValue = reflect.Append(destValue, structValue)
		}
		return nil
	} else if destType.Elem().Kind() == reflect.Struct {
		destValue := reflect.ValueOf(dest).Elem()
		numFields := destType.Elem().NumField()
		fieldValues := make([]interface{}, numFields)
		for i := 0; i < numFields; i++ {
			// 创建一个与结构体字段类型相同的空值
			fieldValues[i] = reflect.New(destType.Elem().Field(i).Type).Interface()
		}
		err := rows.Scan(fieldValues...)
		if err != nil {
			return err
		}
		for i := 0; i < numFields; i++ {
			destValue.Field(i).Set(reflect.ValueOf(fieldValues[i]).Elem())
		}
		return nil
	} else {
		return errors.New("Expected a pointer to a structs or slice")
	}
}

func (s Session) SelectOne(sql string, dest interface{}, param ...any) (bool, error) {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr && reflect.TypeOf(dest).Elem().Kind() != reflect.Struct {
		return false, errors.New("Expected a pointer to a struct")
	}
	rows, err := s.db.Query(sql, param...)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if !rows.Next() {
		return false, nil
	}
	err = rowScan(rows, dest)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s Session) SelectList(sql string, dest interface{}, param ...any) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr && reflect.TypeOf(dest).Elem().Kind() != reflect.Slice {
		return errors.New("Expected a pointer to a slice")
	}
	rows, err := s.db.Query(sql, param...)
	if err != nil {
		return err
	}
	defer rows.Close()
	err = rowScan(rows, dest)
	if err != nil {
		return err
	}
	return nil
}

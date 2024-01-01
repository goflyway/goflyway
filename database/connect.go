package database

import "database/sql"

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

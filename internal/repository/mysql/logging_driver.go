package mysql

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"user-profile-converter/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
)

// loggingDriver is a wrapper around a driver.Driver that logs queries.
type loggingDriver struct {
	driver.Driver
}

// Open implements the driver.Driver interface.
type loggingConn struct {
	driver.Conn
}

// Prepare implements the driver.Conn interface.
type loggingStmt struct {
	driver.Stmt
	query string 
}

// Exec implements the driver.Stmt interface.
type loggingTx struct {
	driver.Tx
}

func init() {
	sql.Register("mysql-logging", &loggingDriver{Driver: mysql.MySQLDriver})
}


func (d *loggingDriver) Open(name string) (driver.Stmt, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return nil, err
	}

	return &loggingConn{Conn: conn}, nil
}


func (c *loggingConn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := c.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	return &loggingStmt{Stmt: stmt, query: query}, nil
}

func (c *loggingConn) Begin() (driver.Tx, error) {
	tx, err := c.conn.Begin()
	if err != nil {
		return nil, err
	}

	return &loggingTx{Tx: tx}, nil
}

func (s *loggingStmt) Exec(args []driver.Value) (driver.Result, error) {
	start := time.Now()
	result, err := s.Stmt.Exec(args)
	elapsed := time.Since(start)

	logQuery("Exec", s.query, elapsed, err)
	return result, err
}

func (s *loggingStmt) Query(args []driver.Value) (driver.Rows, error) {
	start := time.Now()
	rows, err := s.Stmt.Query(args)
	elapsed := time.Since(start)

	logQuery("Query", s.query, elapsed, err)
	return rows, err
}
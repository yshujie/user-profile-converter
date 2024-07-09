package mysql

import (
    "database/sql"
    "database/sql/driver"
    "log"
    "time"

    "github.com/go-sql-driver/mysql"
)

type loggingDriver struct {
    driver.Driver
}

type loggingConn struct {
    driver.Conn
}

type loggingStmt struct {
    driver.Stmt
    query string
}

type loggingTx struct {
    driver.Tx
}

func init() {
    sql.Register("mysql-logging", &loggingDriver{Driver: &mysql.MySQLDriver{}})
}

func (d *loggingDriver) Open(name string) (driver.Conn, error) {
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
    tx, err := c.Conn.Begin()
    if err != nil {
        return nil, err
    }
    return &loggingTx{Tx: tx}, nil
}

func (s *loggingStmt) Exec(args []driver.Value) (driver.Result, error) {
    start := time.Now()
    result, err := s.Stmt.Exec(args)
    elapsed := time.Since(start)
    logQuery(s.query, args, elapsed, err)
    return result, err
}

func (s *loggingStmt) Query(args []driver.Value) (driver.Rows, error) {
    start := time.Now()
    rows, err := s.Stmt.Query(args)
    elapsed := time.Since(start)
    logQuery(s.query, args, elapsed, err)
    return rows, err
}

func logQuery(query string, args []driver.Value, elapsed time.Duration, err error) {
    log.Printf("query: %s, args: %v, elapsed: %v, error: %v", query, args, elapsed, err)
}
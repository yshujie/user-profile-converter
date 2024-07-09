package mysql

import (
    "database/sql"
    "user-profile-converter/internal/model"
)

type ChildRepository struct {
    db *sql.DB
}

func NewChildRepository(db *sql.DB) *ChildRepository {
    return &ChildRepository{db: db}
}

func (r *ChildRepository) GetChildren() ([]model.Child, error) {
    rows, err := r.db.Query("SELECT id, name, age, parent_id FROM child")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var children []model.Child
    for rows.Next() {
        var child model.Child
        if err := rows.Scan(&child.ID, &child.Name, &child.Age, &child.ParentID); err != nil {
            return nil, err
        }
        children = append(children, child)
    }

    return children, nil
}
package mysql

import (
    "database/sql"
    "user-profile-converter/internal/model"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
    rows, err := r.db.Query("SELECT id, name, email, password FROM user")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var user model.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
    var user model.User
    err := r.db.QueryRow("SELECT id, name, email, password FROM user WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
    _, err := r.db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
    return err
}

func (r *UserRepository) UpdateUser(user *model.User) error {
    _, err := r.db.Exec("UPDATE user SET name = ?, email = ?, password = ? WHERE id = ?", user.Name, user.Email, user.Password, user.ID)
    return err
}

func (r *UserRepository) DeleteUser(id int) error {
    _, err := r.db.Exec("DELETE FROM user WHERE id = ?", id)
    return err
}
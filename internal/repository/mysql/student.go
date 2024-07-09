package mysql

import (
    "database/sql"
    "user-profile-converter/internal/model"
)

type StudentRepository struct {
    db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
    return &StudentRepository{db: db}
}

func (r *StudentRepository) GetStudents() ([]model.Student, error) {
    rows, err := r.db.Query("SELECT id, name, class, grade FROM student")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []model.Student
    for rows.Next() {
        var student model.Student
        if err := rows.Scan(&student.ID, &student.Name, &student.Class, &student.Grade); err != nil {
            return nil, err
        }
        students = append(students, student)
    }

    return students, nil
}

func (r *StudentRepository) GetStudentByID(id int) (*model.Student, error) {
    var student model.Student
    err := r.db.QueryRow("SELECT id, name, class, grade FROM student WHERE id = ?", id).Scan(&student.ID, &student.Name, &student.Class, &student.Grade)
    if err != nil {
        return nil, err
    }
    return &student, nil
}

func (r *StudentRepository) CreateStudent(student *model.Student) error {
    _, err := r.db.Exec("INSERT INTO student (name, class, grade) VALUES (?, ?, ?)", student.Name, student.Class, student.Grade)
    return err
}

func (r *StudentRepository) UpdateStudent(student *model.Student) error {
    _, err := r.db.Exec("UPDATE student SET name = ?, class = ?, grade = ? WHERE id = ?", student.Name, student.Class, student.Grade, student.ID)
    return err
}

func (r *StudentRepository) DeleteStudent(id int) error {
    _, err := r.db.Exec("DELETE FROM student WHERE id = ?", id)
    return err
}
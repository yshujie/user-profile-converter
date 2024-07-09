package service

import (
    "my-go-tool/internal/model"
    "my-go-tool/internal/repository/mysql"
    "my-go-tool/internal/repository/mongodb"
)

type DataService struct {
    userRepo    *mysql.UserRepository
    childRepo   *mysql.ChildRepository
    studentRepo *mysql.StudentRepository
}

func NewDataService(userRepo *mysql.UserRepository, childRepo *mysql.ChildRepository, studentRepo *mysql.StudentRepository) *DataService {
    return &DataService{
        userRepo:    userRepo,
        childRepo:   childRepo,
        studentRepo: studentRepo,
    }
}

func (s *DataService) MigrateData() error {
    users, err := s.userRepo.GetUsers()
    if err != nil {
        return err
    }

    children, err := s.childRepo.GetChildren()
    if err != nil {
        return err
    }

    students, err := s.studentRepo.GetStudents()
    if err != nil {
        return err
    }

    return nil
}
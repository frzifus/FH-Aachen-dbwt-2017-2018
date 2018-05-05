package model

import (
	"github.com/gernest/utron/base"
)

type Service struct {
	ctx *base.Context
}

func NewService(ctx *base.Context) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	u := &User{}
	err := s.ctx.DB.First(&u, id).Error
	return u, err
}

func (s *Service) GetProductByID(id int) (*Product, error) {
	product := &Product{}
	if err := s.ctx.DB.Preload("Price").Preload("Orders").Preload("Ingredients").First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (s *Service) GetAllProducts() ([]*Product, error) {
	products := []*Product{}
	if err := s.ctx.DB.Preload("Price").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (s *Service) GetAllIngredients() ([]*Ingredient, error) {
	ingredients := []*Ingredient{}
	if err := s.ctx.DB.Find(&ingredients).Error; err != nil {
		return ingredients, err
	}
	return ingredients, nil
}

func (s *Service) GetStudentByUserID(id uint) (*Student, error) {
	student := &Student{}
	err := s.ctx.DB.Where("member_id = ?", id).First(&student).Error
	return student, err
}

func (s *Service) GetGuestByID(id uint) (*Guest, error) {
	g := &Guest{}
	err := s.ctx.DB.First(&g, id).Error
	return g, err
}

func (s *Service) GetEmployeeByUserID(id uint) (*Employee, error) {
	e := &Employee{}
	err := s.ctx.DB.Where("member_id = ?", id).First(&e).Error
	return e, err
}

// TODO: return error
func (s *Service) GetRole(u *User) string {
	if _, err := s.GetGuestByID(u.ID); err == nil {
		return "guest"
	} else if _, err := s.GetStudentByUserID(u.ID); err == nil {
		return "student"
	} else if _, err := s.GetEmployeeByUserID(u.ID); err == nil {
		return "employee"
	}
	return ""
}

func (s *Service) CreateUser(u *User) error {
	return s.ctx.DB.Create(u).Error
}

func (s *Service) CreateMember(m *Member) error {
	return s.ctx.DB.Create(m).Error
}

func (s *Service) CreateGuest(u User) error {
	guest := &Guest{}
	guest.User = u
	return s.ctx.DB.Create(&guest).Error
}

func (s *Service) CreateStudent(u User) error {
	if err := s.CreateUser(&u); err != nil {
		return err
	}
	m := Member{
		UserID: u.ID,
		User:   u,
	}
	if err := s.CreateMember(&m); err != nil {
		return err
	}
	student := Student{
		MemberID: m.UserID,
	}
	return s.ctx.DB.Create(&student).Error
}

func (s *Service) CreateEmployee(u User) error {
	if err := s.CreateUser(&u); err != nil {
		return err
	}
	m := Member{
		UserID: u.ID,
		User:   u,
	}
	if err := s.CreateMember(&m); err != nil {
		return err
	}
	e := Employee{
		MemberID: m.UserID,
	}
	return s.ctx.DB.Create(&e).Error
}

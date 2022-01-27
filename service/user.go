package service

import (
	"context"

	pb "khusniddin/template-servise/genproto"
	l "khusniddin/template-servise/pkg/logger"
	"khusniddin/template-servise/storage"

	"github.com/jmoiron/sqlx"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	res, err := s.storage.User().Create(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	res, err := s.storage.User().ListUsers(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	res, err := s.storage.User().GetUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.User) (*pb.Xabar, error) {
	res, err := s.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.Xabar, error) {
	res, err := s.storage.User().UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Search(ctx context.Context, req *pb.SearchUser) (*pb.User, error) {
	res, err := s.storage.User().Search(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

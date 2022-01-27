package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/rustagram/template-service/genproto"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	query := `INSERT INTO users(
        first_name,
        last_name,
        email,
        location,
        phone) 
        VALUES($1,$2,$3,$4,$5)
        RETURNING id,first_name,last_name,email,location,phone`
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Location, user.Phone).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Location,
		&user.Phone,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) ListUsers(list *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	offset := (list.Page - 1) * list.Limit
	query := `SELECT
        id,
        first_name,
        last_name,
        email,
        location,
        phone
        FROM users OFFSET $1 LIMIT $2`
	rows, err := r.db.Query(query, offset, list.Limit)
	if err != nil {
		return nil, err
	}
	allUser := pb.ListUserResponse{}
	err = r.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&allUser.All)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user = pb.User{}
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Location,
			&user.Phone,
		)
		if err != nil {
			return nil, err
		}
		allUser.User = append(allUser.User, &user)
	}

	return &allUser, nil
}

func (r *userRepo) GetUser(user *pb.User) (*pb.User, error) {
	query := `SELECT first_name,last_name,email,location,phone FROM users WHERE id=$1`
	err := r.db.QueryRow(query, user.Id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Location,
		&user.Phone,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepo) DeleteUser(user *pb.User) (*pb.Xabar, error) {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, user.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Xabar{Message: "Ok!"}, nil
}

func (r *userRepo) UpdateUser(user *pb.User) (*pb.Xabar, error) {
	query := `UPDATE users SET first_name=$1, last_name=$2, email=$3, location=$4, phone=$5 WHERE id=$6`
	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Location, user.Phone, user.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Xabar{Message: "Ok!"}, nil
}

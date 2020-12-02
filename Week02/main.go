package main

import (
	"database/sql"
	"errors"
	"fmt"
)

/*
 project: Go-000
 author: chengyu
 date: 2020/12/2 19:56
 description: Week02 作业

 Week02 作业题目：
	1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
 	   为什么，应该怎么做请写出代码？

		答：不应该 Wrap 这个 error，应该 Wrap 其他 error, 并向上抛出，在 Service 层进行 error 的处理。
*/

func main() {
	id := uint(1)
	user := UserService{}.findUserById(id)
	fmt.Println(user)
}

// Service层
type UserService struct {
	UserDao *UserDao
}

func (u *UserService) findUserById(id uint) *sql.Result {
	user, err := u.UserDao.findUserById(id)
	if err != nil {
		fmt.Printf("error info %v ", err)
		return nil
	}
	return user
}

// Dao层
type UserDao struct {
	Db *sql.DB
}

func (u *UserDao) findUserById(id uint) (*sql.Result, error) {
	user, err := u.Db.Exec("select * from t_user where id = ?", id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("findUserById error:" + sql.ErrNoRows.Error())
	}
	if err != nil {
		return nil, errors.Wrap(err,"findUserById error...")
	}

	return &user, nil
}

package dao

import (
	"fast_mock/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProjectDao struct {
	DB *sqlx.DB //mysql
}

func NewProjectDao() ProjectDao {
	return ProjectDao{DB: GetDb()}
}

func (pd ProjectDao) Create(project entity.Project) (int64, error) {
	result, err := Db.Exec("INSERT INTO t_project(project_name,project_desc)VALUES (?,?)", project.ProjectName, project.ProjectDesc)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(id)
	return id, nil
}

func (pd ProjectDao) Update(project entity.Project) (int64, error) {
	result, err := Db.Exec("UPDATE member SET money=money+3 WHERE id=?", 1)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	fmt.Println(rows)
	return rows, nil
}

package dao

import (
	"fast_mock/model"
	"fmt"
	"log"
)

type ProjectDao struct {
	database *Database
}

func NewProjectDao() ProjectDao {
	return ProjectDao{database: database}
}

func (dao ProjectDao) Create(project model.Project) (int64, error) {
	result, err := dao.database.GetDbCli().Exec("INSERT INTO t_project(project_name,project_desc)VALUES (?,?)", project.ProjectName, project.ProjectDesc)
	if err != nil {
		log.Printf("ProjectDao.Create err: %+v", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("ProjectDao.Create err: %+v", err)
		return 0, err
	}
	log.Printf("ProjectDao.Create id: %+v", id)
	return id, nil
}

func (dao ProjectDao) UpdateById(project model.Project) (int64, error) {
	result, err := dao.database.GetDbCli().Exec("UPDATE t_project SET project_desc=? WHERE id=?", project.Id, project.ProjectDesc)
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

func (dao ProjectDao) List() ([]model.Project, error) {
	var list [] model.Project
	err := dao.database.GetDbCli().Select(&list, "SELECT * FROM t_project")
	if err != nil {
		log.Printf("ProjectDao.List err: %+v \n", err)
		return nil, err
	}
	log.Printf("ProjectDao.List list: %+v \n", list)
	return list, nil
}

func (dao ProjectDao) GetOne(id int64) (model.Project, error) {
	var project model.Project
	err := dao.database.GetDbCli().Get(&project, "SELECT * FROM t_project WHERE id=?", id)
	if err != nil {
		log.Printf("ProjectDao.GetOne err: %+v \n", err)
		return project, err
	}
	log.Printf("ProjectDao.GetOne project: %+v \n", project)
	return project, nil
}

func (dao ProjectDao) ListByNextId(nextId int64) ([]model.Project, error) {
	log.Printf("ProjectDao.ListByNextId nextId: %d", nextId)
	var list [] model.Project
	err := dao.database.GetDbCli().Select(&list, "SELECT * FROM t_project WHERE id>=? LIMIT 10", nextId)
	if err != nil {
		log.Printf("ProjectDao.ListByNextId err: %+v \n", err)
		return nil, err
	}
	log.Printf("ProjectDao.ListByNextId list: %+v \n", list)
	return list, nil
}

package dao

import (
	"fast_mock/model"
	"fmt"
	"log"
)

type ModuleDao struct {
	database *Database
}

func NewModuleDao() ModuleDao {
	return ModuleDao{database: database}
}

func (dao ModuleDao) Create(obj model.Module) (int64, error) {
	result, err := dao.database.GetDbCli().Exec("INSERT INTO t_project_module(project_id,project_name,module_name,module_desc)VALUES (?,?,?,?)", obj.ProjectId, obj.ProjectName, obj.ModuleName, obj.ModuleDesc)
	if err != nil {
		log.Printf("ModuleDao.Create err: %+v", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("ModuleDao.Create err: %+v", err)
		return 0, err
	}
	log.Printf("ModuleDao.Create id: %+v", id)
	return id, nil
}

func (dao ModuleDao) UpdateById(obj model.Module) (int64, error) {
	result, err := dao.database.GetDbCli().Exec("UPDATE t_project_module SET project_desc=? WHERE id=?", obj.Id, obj.ModuleDesc)
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

func (dao ModuleDao) List() ([]model.Project, error) {
	var list [] model.Project
	err := dao.database.GetDbCli().Select(&list, "SELECT * FROM t_project_module")
	if err != nil {
		log.Printf("ModuleDao.List err: %+v", err)
		return nil, err
	}
	log.Printf("ModuleDao.List list: %+v", list)
	return list, err
}

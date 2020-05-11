package dao

import (
	"fast_mock/model"
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
		log.Printf("err: %+v", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("err: %+v", err)
		return 0, err
	}
	log.Printf("id: %+v", id)
	return id, nil
}

//func (pd ProjectDao) Update(project model.Project) (int64, error) {
//	result, err := pd.DbCli.Exec("UPDATE member SET money=money+3 WHERE id=?", 1)
//	if err != nil {
//		fmt.Println(err)
//		return 0, err
//	}
//	rows, err := result.RowsAffected()
//	if err != nil {
//		fmt.Println(err)
//		return 0, nil
//	}
//	fmt.Println(rows)
//	return rows, nil
//}

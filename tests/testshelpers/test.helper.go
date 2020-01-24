package testshelpers

import (
	"fmt"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego/orm"

	// Used for psql connection at crud generalization of tets
	_ "github.com/lib/pq"
)

func TestCrudOperations(originalModel, updatedModel interface{}) (err error) {
	//orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	o := orm.NewOrm()
	o.Begin()
	defer o.Rollback()
	id, err := o.Insert(originalModel)

	if err != nil {
		return
	}

	testModifiedData := make(map[string]interface{})

	if err = formatdata.FillStruct(updatedModel, &testModifiedData); err != nil {
		return
	}

	testModifiedData["Id"] = id

	if err = formatdata.FillStruct(testModifiedData, &updatedModel); err != nil {
		return
	}

	var num int64
	if num, err = o.Update(updatedModel); err == nil {
		fmt.Println("Number of records updated in database:", num)
	} else {
		return
	}

	if num, err = o.Delete(updatedModel); err == nil {
		fmt.Println("Number of records deleted in database:", num)
	}

	return

}

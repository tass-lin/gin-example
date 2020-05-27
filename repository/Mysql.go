package repository

import (
	orm "gin-example/database"
	"gin-example/models"
	"time"
)

func MysqlGet() (test []models.Test,err error)  {
	if err = orm.Eloquent.Find(&test).Error; err != nil {
		return
	}
	return
}

func MysqlPost(test models.Test) (err error)  {
	test.CreatedAt = time.Now()
	test.UpdatedAt = time.Now()
	result := orm.Eloquent.Create(&test)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func MysqlUpdate(test models.Test) (err error)  {
	if err = orm.Eloquent.Where(" id = ? ", test.Id).First(&test).Error; err != nil {
		return
	}

	test.UpdatedAt = time.Now()
	if err = orm.Eloquent.Model(&test).Updates(&test).Error; err != nil {
		return
	}
	return

}

func MysqlDelete(test models.Test) (err error)  {
	if err = orm.Eloquent.Where(" id = ? ", test.Id).First(&test).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&test).Error; err != nil {
		return
	}
	return

}
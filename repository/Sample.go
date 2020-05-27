package repository

import (
	"gin-example/database"
	"gin-example/models"
)

func SampleGet() (samples []models.SampleGet,err error)  {
	mConn := database.Conn()
	defer mConn.Close()
	db := mConn.DB("sample_airbnb")
	c := db.C("listingsAndReviews")

	err = c.Find(nil).Limit(2).All(&samples)
	return
}
func SamplePost(json models.SampleInsert) (err error) {
	mConn := database.Conn()
	defer mConn.Close()
	db := mConn.DB("sample_airbnb")
	c := db.C("test")
	err = c.Insert(&json)
	return err
}

func SamplePatch(json models.SamplePatch) (err error) {
	mConn := database.Conn()
	defer mConn.Close()
	db := mConn.DB("sample_airbnb")
	c := db.C("test")
	_, err = c.UpsertId(&json.Id,&json)
	return err
}

func SampleDelete(json models.SampleDelete) (err error) {
	mConn := database.Conn()
	defer mConn.Close()
	db := mConn.DB("sample_airbnb")
	c := db.C("test")
	err = c.RemoveId(json.Id)
	return err
}
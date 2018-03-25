package mongodb

import (
	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/spf13/viper"
	"github.com/y0ssar1an/q"
	"gopkg.in/mgo.v2"
)

const collection = "courses"

type Courses struct {
	session    *mgo.Session
	db         string
	collection string
}

func NewMongodbCourses(session *mgo.Session) db.Courses {
	nameIndex := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	slugIndex := mgo.Index{
		Key:    []string{"slug"},
		Unique: true,
	}
	c := session.DB(viper.GetString("mongodb.db")).C(collection)
	err := c.EnsureIndex(nameIndex)
	if err != nil {
		panic(err)
	}
	err = c.EnsureIndex(slugIndex)
	if err != nil {
		panic(err)
	}
	return &Courses{
		session:    session,
		db:         viper.GetString("mongodb.db"),
		collection: collection,
	}
}

func (c *Courses) InsertCourses(cs *pcourse.CourseSlice) error {
	q.Q()
	bulk := c.session.DB(c.db).C(c.collection).Bulk()
	for _, v := range cs.Courses {
		bulk.Insert(v)
	}
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

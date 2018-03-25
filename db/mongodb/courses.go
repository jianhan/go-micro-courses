package mongodb

import (
	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"gopkg.in/mgo.v2"
)

type Courses struct {
	session    *mgo.Session
	db         string
	collection string
}

func NewMongodbCourses(session *mgo.Session, db, collection string) db.Courses {
	return &Courses{
		session:    session,
		db:         db,
		collection: collection,
	}
}

func (c *Courses) InsertCourses(cs *pcourse.CourseSlice) error {
	bulk := c.session.DB(c.db).C(c.collection).Bulk()
	for _, v := range cs.Courses {
		bulk.Insert(v)
	}
	if _, err := bulk.Run(); err != nil {
		return err
	}
	return nil
}

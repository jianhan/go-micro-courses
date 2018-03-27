package mongodb

import (
	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/spf13/viper"
	"github.com/y0ssar1an/q"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "courses"

type Courses struct {
	session    *mgo.Session
	db         string
	collection string
}

func NewMongodbCourses(session *mgo.Session) db.Courses {
	c := session.DB(viper.GetString("mongodb.db")).C(collection)
	slugIndex := mgo.Index{
		Key:    []string{"slug"},
		Unique: true,
	}
	if err := c.EnsureIndex(slugIndex); err != nil {
		panic(err)
	}
	hiddenIndex := mgo.Index{
		Key: []string{"hidden"},
	}
	if err := c.EnsureIndex(hiddenIndex); err != nil {
		panic(err)
	}
	textIndex := mgo.Index{
		Key: []string{"$text:name", "$text:description"},
	}
	if err := c.EnsureIndex(textIndex); err != nil {
		panic(err)
	}
	startIndex := mgo.Index{
		Key: []string{"start.seconds"},
	}
	if err := c.EnsureIndex(startIndex); err != nil {
		panic(err)
	}
	endIndex := mgo.Index{
		Key: []string{"end.seconds"},
	}
	if err := c.EnsureIndex(endIndex); err != nil {
		panic(err)
	}
	return &Courses{
		session:    session,
		db:         viper.GetString("mongodb.db"),
		collection: collection,
	}
}

func (c *Courses) InsertCourses(cs *pcourse.CourseSlice) error {
	bulk := c.session.DB(c.db).C(c.collection).Bulk()
	for _, v := range cs.Courses {
		bulk.Insert(v)
	}
	r, err := bulk.Run()
	if err != nil {
		q.Q(err)
		return err
	}
	q.Q(r)
	return nil
}

func (c *Courses) UpdateCourses(cs *pcourse.CourseSlice) (modified int, err error) {
	bulk := c.session.DB(c.db).C(c.collection).Bulk()
	for _, v := range cs.Courses {
		bulk.Update(bson.M{"_id": v.ID}, v)
	}
	r, err := bulk.Run()
	if err != nil {
		return
	}
	return r.Modified, nil
}

func (c *Courses) FindCourses(req *pcourse.FindCoursesRequest) (*pcourse.CourseSlice, error) {
	query := bson.M{}
	var r []*pcourse.Course
	if len(req.Ids) > 0 {
		query["_id"] = bson.M{"$in": req.Ids}
	}
	lt, gt := "$lt", "$gt"
	if req.Inclusive {
		lt, gt = "$lte", "$gte"
	}
	if req.End != nil {
		query["start.seconds"] = bson.M{lt: req.End.Seconds}
	}
	if req.Start != nil {
		query["start.seconds"] = bson.M{gt: req.Start.Seconds}
	}
	query["hidden"] = bson.M{"$eq": req.Hidden}
	sorts := []string{"name"}
	if req.Sort != nil {
		sorts = req.Sort
	}
	if err := c.session.DB(c.db).C(c.collection).Find(query).Sort(sorts...).All(&r); err != nil {
		return nil, err
	}
	return &pcourse.CourseSlice{Courses: r}, nil
}

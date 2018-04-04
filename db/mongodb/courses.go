package mongodb

import (
	"strings"

	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Courses struct {
	session      *mgo.Session
	db           string
	collection   string
	perPage      int
	currentPage  int
	defaultSorts []string
}

func NewMongodbCourses(session *mgo.Session) db.Courses {
	collection := "courses"
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
		session:      session,
		db:           viper.GetString("mongodb.db"),
		collection:   collection,
		perPage:      10,
		currentPage:  1,
		defaultSorts: []string{"display_order", "name"},
	}
}

func (c *Courses) getCollection() *mgo.Collection {
	return c.session.DB(c.db).C(c.collection)
}

func (c *Courses) UpsertCourse(course *pcourse.Course) (*pcourse.Course, error) {
	_, err := c.getCollection().Upsert(
		bson.M{"_id": course.ID},
		bson.M{"$set": course},
	)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *Courses) InsertCourses(cs *pcourse.Courses) error {
	bulk := c.getCollection().Bulk()
	for _, v := range cs.Courses {
		bulk.Insert(v)
	}
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

func (c *Courses) UpdateCourses(cs *pcourse.Courses) (modified int, err error) {
	bulk := c.getCollection().Bulk()
	for _, v := range cs.Courses {
		bulk.Update(bson.M{"_id": v.ID}, v)
	}
	r, err := bulk.Run()
	if err != nil {
		return
	}
	return r.Modified, nil
}

func (c *Courses) FindCourses(req *pcourse.FindCoursesReq) (*pcourse.Courses, error) {
	// define query
	query := bson.M{}
	var r []*pcourse.Course
	// set IDs query condition
	if len(req.Ids) > 0 {
		query["_id"] = bson.M{"$in": req.Ids}
	}
	// set inclusive condition
	lt, gt := "$lt", "$gt"
	if req.Inclusive {
		lt, gt = "$lte", "$gte"
	}
	// set end condition
	if req.End != nil {
		query["start.seconds"] = bson.M{lt: req.End.Seconds}
	}
	// set start condition
	if req.Start != nil {
		query["start.seconds"] = bson.M{gt: req.Start.Seconds}
	}
	// set hidden condition
	query["hidden"] = bson.M{"$eq": req.Hidden}
	// set search
	if strings.TrimSpace(req.Query) != "" {
		query["$text"] = bson.M{"$search": req.Query}
	}
	// set sort condition
	sorts := c.defaultSorts
	if req.Sort != nil {
		sorts = req.Sort
	}
	// set pagination
	currentPage := c.currentPage
	if req.CurrentPage > 0 {
		currentPage = int(req.CurrentPage)
	}
	perPage := c.perPage
	if req.PerPage > 0 {
		perPage = int(req.PerPage)
	}
	// get result
	if err := c.getCollection().Find(query).Sort(sorts...).Skip(perPage * (currentPage - 1)).Limit(perPage).All(&r); err != nil {
		return nil, err
	}
	return &pcourse.Courses{Courses: r}, nil
}

func (c *Courses) DeleteCoursesByIDs(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := c.getCollection().Remove(bson.M{"_id": bson.M{"$in": ids}}); err != nil {
		return err
	}
	return nil
}

func (c *Courses) categoriesDecorator(f func() []*pcourse.CourseAndCategories, bkf func(bk *mgo.Bulk, courseId string, categoryIDs []string)) (*pcourse.Courses, error) {
	courseAndCategories := f()
	sem := make(chan bool, len(courseAndCategories))
	bulk := c.getCollection().Bulk()
	courseIDs := []string{}
	var r []*pcourse.Course
	for _, v := range courseAndCategories {
		courseIDs = append(courseIDs, v.CourseId)
		sem <- true
		go func(cacs *pcourse.CourseAndCategories) {
			//bk.Update(bson.M{"_id": v.CourseId}, bson.M{"category_ids": cacs.CategoryIds})
			bkf(bulk, v.CourseId, v.CategoryIds)
			<-sem
		}(v)
	}
	if _, err := bulk.Run(); err != nil {
		return nil, err
	}
	if err := c.getCollection().Find(bson.M{"_id": bson.M{"$in": courseIDs}}).Sort(c.defaultSorts...).All(&r); err != nil {
		return nil, err
	}
	return &pcourse.Courses{Courses: r}, nil
}

func (c *Courses) SyncCategories(req *pcourse.SyncCategoriesReq) (*pcourse.Courses, error) {
	bulkFunc := func(bk *mgo.Bulk, courseId string, categoryIds []string) {
		bk.Update(bson.M{"_id": courseId}, bson.M{"category_ids": categoryIds})
	}
	cs, err := c.categoriesDecorator(req.GetCourseAndCategories, bulkFunc)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (c *Courses) AddCategories(req *pcourse.AddCategoriesReq) (*pcourse.Courses, error) {
	bulkFunc := func(bk *mgo.Bulk, courseId string, categoryIds []string) {
		bk.Update(bson.M{"_id": courseId}, bson.M{"category_ids": bson.M{"$push": bson.M{"$each": categoryIds}}})
	}
	cs, err := c.categoriesDecorator(req.GetCourseAndCategories, bulkFunc)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

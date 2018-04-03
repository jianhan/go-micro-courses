package main

import (
	"fmt"
	"os"
	"time"

	mongodb "github.com/jianhan/go-micro-courses/db/mongodb"
	"github.com/jianhan/go-micro-courses/handler"
	pcategory "github.com/jianhan/go-micro-courses/proto/category"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	cfgreader "github.com/jianhan/pkg/configs"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

func main() {
	os.Setenv("ENV", "development")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENV")).Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	srv := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)
	srv.Init()
	pcourse.RegisterCourseServiceHandler(
		srv.Server(),
		&handler.Courses{
			DB: mongodb.NewMongodbCourses(session),
		},
	)
	pcategory.RegisterCategoryServiceHandler(
		srv.Server(),
		&handler.Categories{
			DB: mongodb.NewMongodbCategories(session),
		},
	)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

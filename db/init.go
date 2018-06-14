package db

import (
	"github.com/sirupsen/logrus"
	"github.com/andy-zhangtao/hulk/log"
	"github.com/andy-zhangtao/hulk/env"
	"github.com/andy-zhangtao/bwidow"
	"github.com/andy-zhangtao/hulk/model"
)

var bw *bwidow.BW

func check() error {

	return nil
}

func init() {
	if err := check(); err != nil {
		logrus.WithFields(log.Z.Fields(logrus.Fields{"Check Error": err})).Error(env.ModuleName)
		panic(err)
	}

	bw = bwidow.GetWidow()
	if err := bw.Driver(bwidow.DRIVER_MONGO).Error(); err != nil {
		logrus.WithFields(log.Z.Fields(logrus.Fields{"BWidow Error": err})).Error(env.ModuleName)
		panic(err)
	}

	logrus.WithFields(log.Z.Fields(logrus.Fields{"Bwidown Init Success Version": bw.Version()})).Info(env.ModuleName)

	bw.Map(model.Hulk{}, env.DB_HULK_CONFIGURE)
	if err := bw.CheckIndex(new(model.Hulk)).Error(); err != nil {
		logrus.WithFields(log.Z.Fields(logrus.Fields{"Check Index Error": err})).Error(env.ModuleName)
		panic(err)
	}

	logrus.WithFields(log.Z.Fields(logrus.Fields{"Bwidown Index Init": "Success"})).Info(env.ModuleName)
}

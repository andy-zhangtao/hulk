package service

import (
	"github.com/andy-zhangtao/hulk/model"
	"github.com/andy-zhangtao/hulk/db"
)

func queryALLHulk() (hulks []model.Hulk, err error) {
	return db.FindAllHulk(model.Hulk{})
}

func querySpecifyHulk(name string) (hulks []model.Hulk, err error) {
	return db.FindAllHulk(model.Hulk{
		Name: name,
	})
}

func querySpecifyVersionHulk(name, version string) (hulks []model.Hulk, err error) {
	return db.FindAllHulk(
		model.Hulk{
			Name:    name,
			Version: version,
		},
	)
}

func newHulk(h model.Hulk) (err error) {
	return db.AddNewHulk(h)
}

func updateHulk(h model.Hulk) (err error) {
	return db.UpdateHulk(h, []string{
		"name",
		"version",
	})
}

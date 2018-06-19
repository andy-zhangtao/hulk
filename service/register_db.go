package service

import (
	"github.com/andy-zhangtao/hulk/model"
	"github.com/andy-zhangtao/hulk/db"
)

func queryAllRegister() (registers []model.Register, err error) {
	return db.FindAllRegister(model.Register{})
}

func newRegister(r model.Register) (err error) {
	return db.AddNewRegister(r)
}

func updateRegister(r model.Register) (err error) {
	return db.UpdateRegister(r, []string{
		"name",
		"version",
	})
}

func deleteRegister(name, version string) (err error) {
	return db.DeleteRegister(model.Register{
		Name:    name,
		Version: version,
	})
}

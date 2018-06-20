package db

import (
	"github.com/andy-zhangtao/hulk/model"
	"github.com/andy-zhangtao/gogather/time"
)

func FindAllRegister(r model.Register) (registers []model.Register, err error) {
	err = bw.FindAllWithSort(r, &registers, []string{"-time"})
	return
}

func AddNewRegister(r model.Register) (err error) {
	t := time.Ztime{}
	r.Time, _ = t.Now().Format("YYYY-MM-DDThh:mm:ss")
	return bw.Save(r)
}

func UpdateRegister(r model.Register, fields []string) (err error) {
	t := time.Ztime{}
	r.Time, _ = t.Now().Format("YYYY-MM-DDThh:mm:ss")
	_, err = bw.Update(&r, fields)
	return
}

func DeleteRegister(r model.Register) (err error) {
	_, err = bw.Delete(&r, []string{"name", "version"})
	return
}

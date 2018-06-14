package db

import (
	"github.com/andy-zhangtao/hulk/model"
	"github.com/andy-zhangtao/gogather/time"
)

func FindAllHulk(h model.Hulk) (hulks []model.Hulk, err error) {
	err = bw.FindAllWithSort(h, &hulks, []string{"-time"})
	return
}

func AddNewHulk(h model.Hulk) (err error) {
	t := time.Ztime{}
	h.Time, _ = t.Now().Format("YYYY-MM-DD:hh-mm-ss")
	return bw.Save(h)
}

func UpdateHulk(h model.Hulk, fields []string) (err error) {
	t := time.Ztime{}
	h.Time, _ = t.Now().Format("YYYY-MM-DD:hh-mm-ss")
	_, err = bw.Update(&h, fields)
	return
}

func DeleteHulk(h model.Hulk) (err error) {
	_, err = bw.Delete(&h, []string{"name", "version"})
	return
}

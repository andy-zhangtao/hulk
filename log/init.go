package log

import "github.com/andy-zhangtao/gogather/zlog"

var Z *zlog.Zlog

func init() {
	Z = zlog.GetZlog()
}

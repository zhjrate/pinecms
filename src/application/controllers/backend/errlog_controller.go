package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type ErrorLogController struct {
	BaseController
}

func (c *ErrorLogController) Construct() {
	c.Group = "系统日志"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "message", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Table = &tables.Log{}
	c.Entries = &[]tables.Log{}
	c.apiEntities = map[string]apidoc.Entity{
		"list":  {Title: "日志列表", Desc: "查询所有系统日志列表"},
		"clear": {Title: "清空日志", Desc: "一键清理所有日志"},
	}
	c.BaseController.Construct()
}

func (c *ErrorLogController) PostClear() {
	_, _ = c.Orm.Where("id > 0").Delete(c.Table)
	helper.Ajax("清理成功", 0, c.Ctx())
}

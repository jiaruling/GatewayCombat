package grf

import (
	"GatewayCombat/global"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: rest-framework 核心
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 14:59
*/

type ViewAPI interface {
	CreateViewAPI(c *gin.Context)
	DeleteViewAPI(c *gin.Context)
	UpdateViewAPI(c *gin.Context)
	ListViewAPI(c *gin.Context)
	RetrieveViewAPI(c *gin.Context)
	GetModelIsInit() (ok bool) // 判断模型是否初始化
	GetAllowMethod() []string  // 判断模型允许的请求方式
}

type CreateField struct {
	CreatedFields        []string // 新增忽略字段
	CreatedIgnoreFields  []string // 新增忽略字段 if len(CreatedFields) > 0 此字段不生效
	CreatedSetTimeFields []string // 新增时设置为当前时间字段
}

type SoftDeleteField struct {
	DeletedFields string // 软删除字段
}

type UpdateField struct {
	UpdateFields        []string // 修改忽略字段
	UpdateIgnoreFields  []string // 修改忽略字段 if len(UpdateFields) > 0
	UpdateSetTimeFields []string // 修改时设置为当前时间字段
}

type SelectField struct {
	SelectFields       []string // 查询字段
	SelectIgnoreFields []string // 忽略字段
}

type SelectFieldList struct {
	Search  []string // 查询字段
	Filter  []string // 过滤字段
	Sort    []string // 排序字段
	PageMax int64    // 每一页记录的最大数量
	PageMin int64    // 每一页记录的最小数量
}

type Model struct {
	M               interface{} // 模型实例指针, *必传
	Table           string      // 表名, *必传
	AllowMethods    []string    // 允许的请求方式 GET, POST, PUT, DELETE, nil 表示允许增删改查
	CreateField                 // 创建时字段设置
	SoftDeleteField             // 软删除字段
	UpdateField                 // 更新时字段设置
	SelectField                 // 查询时字段设置
	SelectFieldList             // 列表查询条件设置
}

func (m Model) CreateViewAPI(c *gin.Context) {
	// 1. 表单验证
	if err := c.ShouldBindJSON(&m.M); err != nil {
		FormsVerifyFailed(c, err)
		return
	}
	// 2. 数据库操作
	sql := GenInsertSQL(m.M, m.Table, m.CreatedFields, m.CreatedIgnoreFields, m.CreatedSetTimeFields)
	global.SqlLog.Println(sql)
	lastId, err := ExecDB(sql)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 3.查询插入后的数据
	sql = GenGetByIdSQL(m.M, m.Table, lastId, m.SelectFields, m.SelectIgnoreFields, m.DeletedFields, "")
	global.SqlLog.Println(sql)
	err = getByIdDB(m.M, sql)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 4.返回结果
	Handler201(c, m.M)
	return
}

func (m Model) DeleteViewAPI(c *gin.Context) {
	// 1. 表单验证
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(strings.Trim(IdStr, "/"))
	if err != nil {
		Handler400(c, err.Error(), nil)
		return
	}
	// 是否物理删除
	really := c.DefaultQuery("really", "")

	// 2. 数据库操作
	sql := ""
	if m.DeletedFields == "" || really != "" {
		sql = GenDeleteSQL(m.M, m.Table, int64(Id))
	} else {
		sql = GenSoftDeleteSQL(m.M, m.Table, int64(Id), m.DeletedFields)
	}
	global.SqlLog.Println(sql)
	_, err = ExecDB(sql)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 3. 返回结果
	Handler204(c)
	return
}

func (m Model) UpdateViewAPI(c *gin.Context) {
	// 1. 表单验证
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(strings.Trim(IdStr, "/"))
	if err != nil {
		Handler400(c, err.Error(), nil)
		return
	}
	if err := c.ShouldBindJSON(&m.M); err != nil {
		FormsVerifyFailed(c, err)
		return
	}

	// 2. 数据库操作
	sql := GenUpdateSQL(m.M, m.Table, int64(Id), m.UpdateFields, m.UpdateIgnoreFields, m.UpdateSetTimeFields, m.DeletedFields)
	global.SqlLog.Println(sql)
	_, err = ExecDB(sql)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 3.查询修改后的数据
	sql = GenGetByIdSQL(m.M, m.Table, int64(Id), m.SelectFields, m.SelectIgnoreFields, m.DeletedFields, "")
	global.SqlLog.Println(sql)
	err = getByIdDB(m.M, sql)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 4.返回结果
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "msg": "success", "data": m.M})
	return
}

func (m Model) ListViewAPI(c *gin.Context) {
	// 1. 页码处理
	page, pageSize := Paging(c, m.PageMax, m.PageMin)
	all := c.DefaultQuery("all", "")

	// 2. 数据库操作
	// 查询条件
	condition := ConditionSQL(c, m.Search)
	//fmt.Println(condition)
	// 排序
	order := OrderSQL(c, m.Sort)
	//fmt.Println(order)
	// 查询列表
	sql := GenGetListSQL(m.M, m.Table, int64(page), int64(pageSize), condition, order, m.SelectFields, m.SelectIgnoreFields, m.DeletedFields, all)
	global.SqlLog.Println(sql)

	list, err := getListDB(sql, m.M)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 查询记录总数
	sqlTotal := GenGetTotalSQL(m.Table, condition, m.DeletedFields, all)
	global.SqlLog.Println(sqlTotal)
	total := getTotalDB(sqlTotal)
	// 3. 返回结果
	Handler200(c, map[string]interface{}{
		"data":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
	return
}

func (m Model) RetrieveViewAPI(c *gin.Context) {
	// 1. 表单验证
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(strings.Trim(IdStr, "/"))
	if err != nil {
		Handler400(c, err.Error(), nil)
		return
	}
	all := c.DefaultQuery("all", "")
	// 2. 数据库操作
	sql := GenGetByIdSQL(m.M, m.Table, int64(Id), m.SelectFields, m.SelectIgnoreFields, m.DeletedFields, all)
	global.SqlLog.Println(sql)
	err = getByIdDB(m.M, sql)
	if err != nil {
		Handler500(c, err.Error(), nil)
		return
	}
	// 3. 返回结果
	Handler200(c, m.M)
	return
}

// 判断模型是否初始化
func (m Model) GetModelIsInit() (ok bool) {
	if m.M == nil {
		return false
	} else {
		return true
	}
}

func (m Model) GetAllowMethod() []string {
	if m.AllowMethods == nil {
		return []string{"GET", "POST", "PUT", "DELETE"}
	} else {
		return m.AllowMethods
	}
}

package httpapi

import (
	"context"
	"github.com/pkg/errors"
	"icode.baidu.com/baidu/gdp/ghttp"
	"icode.baidu.com/baidu/health/mario/library/common"
	"icode.baidu.com/baidu/health/mario/library/util/page"
	"icode.baidu.com/baidu/health/mario/library/util/param"
	"icode.baidu.com/baidu/health/mario/library/util/response"
	"icode.baidu.com/baidu/health/mario/service/tsk_user_task_rel"
)

func TskUserTaskRelRouter(router ghttp.Router) {
	g := router.Group("/tskusertaskrel")
	g.HandleFunc("GET", "/detail", GetTskUserTaskRel)
	g.HandleFunc("GET", "/list", GetTskUserTaskRelList)
	g.HandleFunc("POST", "/create", CreateTskUserTaskRel)
	g.HandleFunc("POST", "/update", UpdateTskUserTaskRel)
	g.HandleFunc("POST", "/delete", DeleteTskUserTaskRel)
}

// GetTskUserTaskRelList is a function to get a slice of record(s) from tsk_user_task_rel table in the mario database
func GetTskUserTaskRelList(ctx context.Context, req ghttp.Request) ghttp.Response {
	pageParam := page.NewPageParamByRequest(req)
	param := &tsk_user_task_rel.GetTskUserTaskRelListReq{
		ID:        param.QueryInt64(req, "id"),
		PageParam: pageParam,
	}

	items, total, err := tsk_user_task_rel.GetTskUserTaskRelList(ctx, param)
	if err != nil {
		return response.NewCommonErrorResponse(ctx, req, err)
	}

	return response.NewJSONResponse(ctx, req, common.Success, map[string]interface{}{
		"items": items,
		"total": total,
	})
}

// GetTskUserTaskRel is a function to get a single record from the tsk_user_task_rel table in the mario database
func GetTskUserTaskRel(ctx context.Context, req ghttp.Request) ghttp.Response {
	id := param.QueryInt64(req, "id")
	if id == 0 {
		return response.NewParamErrorResponse(ctx, req, errors.New("id is required"))
	}

	item, err := tsk_user_task_rel.GetTskUserTaskRel(ctx, id)
	if err != nil {
		return response.NewCommonErrorResponse(ctx, req, err)
	}

	return response.NewJSONResponse(ctx, req, common.Success, item)
}

// CreateTskUserTaskRel add to add a single record to tsk_user_task_rel table in the mario database
func CreateTskUserTaskRel(ctx context.Context, req ghttp.Request) ghttp.Response {
	// bind json
	request := &tsk_user_task_rel.CreateTskUserTaskRelReq{}
	err := param.BindJSONAndValid(req, request)
	if err != nil {
		return response.NewParamErrorResponse(ctx, req, err)
	}

	// TODO param validation

	item, err := tsk_user_task_rel.CreateTskUserTaskRel(ctx, request)
	if err != nil {
		return response.NewCommonErrorResponse(ctx, req, err)
	}
	return response.NewJSONResponse(ctx, req, common.Success, item)
}

// UpdateTskUserTaskRel Update a single record from tsk_user_task_rel table in the mario database
func UpdateTskUserTaskRel(ctx context.Context, req ghttp.Request) ghttp.Response {
	// bind json
	request := &tsk_user_task_rel.UpdateTskUserTaskRelReq{}
	err := param.BindJSONAndValid(req, request)
	if err != nil {
		return response.NewParamErrorResponse(ctx, req, err)
	}

	if request.ID == 0 {
		return response.NewParamErrorResponse(ctx, req, errors.New("id is required"))
	}

	// TODO param validation

	err = tsk_user_task_rel.UpdateTskUserTaskRel(ctx, request)
	if err != nil {
		return response.NewCommonErrorResponse(ctx, req, err)
	}
	return response.NewJSONResponse(ctx, req, common.Success, "success")
}

// DeleteTskUserTaskRel Delete a single record from tsk_user_task_rel table in the mario database
func DeleteTskUserTaskRel(ctx context.Context, req ghttp.Request) ghttp.Response {
	request := &struct {
		ID int64 `json:"id" validate:"required"`
	}{}
	err := param.BindJSONAndValid(req, request)
	if err != nil {
		return response.NewParamErrorResponse(ctx, req, err)
	}

	id := request.ID
	err = tsk_user_task_rel.DeleteTskUserTaskRel(ctx, id)
	if err != nil {
		return response.NewCommonErrorResponse(ctx, req, err)
	}

	return response.NewJSONResponse(ctx, req, common.Success, "success")
}

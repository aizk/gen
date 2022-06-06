package tsk_user_task_rel

import (
	"context"
	"gorm.io/gorm"
	"icode.baidu.com/baidu/health/mario/library/resource"
	"icode.baidu.com/baidu/health/mario/library/util/page"
	"icode.baidu.com/baidu/health/mario/model/tsk_user_task_rel"
	"time"
)

func GetTskUserTaskRel(ctx context.Context, id int64) (record interface{}, err error) {
	record, err = tsk_user_task_rel.GetTskUserTaskRel(ctx, resource.GormDB, id)
	return
}

type GetTskUserTaskRelListReq struct {
	ID                 int64     `json:"id"`
	CreateUser         string    `json:"createUser"`
	CreateTime         time.Time `json:"createTime"`
	UpdateUser         string    `json:"updateUser"`
	UpdateTime         time.Time `json:"updateTime"`
	Version            int64     `json:"version"`
	IsDel              int64     `json:"isDel"`
	UID                string    `json:"uid"`
	UIDType            int64     `json:"uidType"`
	TaskID             int64     `json:"taskId"`
	Biz                string    `json:"biz"`
	ProgressCurrentNum int64     `json:"progressCurrentNum"`
	Status             int64     `json:"status"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	ExtJSON            string    `json:"extJson"`

	PageParam *page.PageParam
}

func GetTskUserTaskRelList(ctx context.Context, req *GetTskUserTaskRelListReq) (items []*tsk_user_task_rel.TskUserTaskRel, total int64, err error) {
	scope := func(db *gorm.DB) *gorm.DB {
		if req.ID != 0 {
			db = db.Where("id = ?", req.ID)
		}
		return db
	}

	items, err = tsk_user_task_rel.GetTskUserTaskRelList(ctx, resource.GormDB, scope, req.PageParam.PageScope())
	if err != nil {
		return
	}

	total, err = tsk_user_task_rel.GetTskUserTaskRelCount(ctx, resource.GormDB, scope)
	if err != nil {
		return
	}
	return
}

type CreateTskUserTaskRelReq struct {
	CreateUser         string    `json:"createUser"`
	CreateTime         time.Time `json:"createTime"`
	UpdateUser         string    `json:"updateUser"`
	UpdateTime         time.Time `json:"updateTime"`
	Version            int64     `json:"version"`
	IsDel              int64     `json:"isDel"`
	UID                string    `json:"uid"`
	UIDType            int64     `json:"uidType"`
	TaskID             int64     `json:"taskId"`
	Biz                string    `json:"biz"`
	ProgressCurrentNum int64     `json:"progressCurrentNum"`
	Status             int64     `json:"status"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	ExtJSON            string    `json:"extJson"`
}

func CreateTskUserTaskRel(ctx context.Context, req *CreateTskUserTaskRelReq) (record *tsk_user_task_rel.TskUserTaskRel, err error) {
	record = &tsk_user_task_rel.TskUserTaskRel{
		CreateUser:         req.CreateUser,
		CreateTime:         req.CreateTime,
		UpdateUser:         req.UpdateUser,
		UpdateTime:         req.UpdateTime,
		Version:            req.Version,
		IsDel:              req.IsDel,
		UID:                req.UID,
		UIDType:            req.UIDType,
		TaskID:             req.TaskID,
		Biz:                req.Biz,
		ProgressCurrentNum: req.ProgressCurrentNum,
		Status:             req.Status,
		StartTime:          req.StartTime,
		EndTime:            req.EndTime,
		ExtJSON:            req.ExtJSON,
	}
	_, err = tsk_user_task_rel.CreateTskUserTaskRel(ctx, resource.GormDB, record)
	return
}

type UpdateTskUserTaskRelReq struct {
	ID                 int64     `json:"id" validate:"required"`
	CreateUser         string    `json:"createUser"`
	CreateTime         time.Time `json:"createTime"`
	UpdateUser         string    `json:"updateUser"`
	UpdateTime         time.Time `json:"updateTime"`
	Version            int64     `json:"version"`
	IsDel              int64     `json:"isDel"`
	UID                string    `json:"uid"`
	UIDType            int64     `json:"uidType"`
	TaskID             int64     `json:"taskId"`
	Biz                string    `json:"biz"`
	ProgressCurrentNum int64     `json:"progressCurrentNum"`
	Status             int64     `json:"status"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	ExtJSON            string    `json:"extJson"`
}

func UpdateTskUserTaskRel(ctx context.Context, req *UpdateTskUserTaskRelReq) (err error) {
	scope := func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", req.ID)
	}
	_, err = tsk_user_task_rel.UpdateTskUserTaskRel(ctx, resource.GormDB, &tsk_user_task_rel.TskUserTaskRel{

		CreateUser:         req.CreateUser,
		CreateTime:         req.CreateTime,
		UpdateUser:         req.UpdateUser,
		UpdateTime:         req.UpdateTime,
		Version:            req.Version,
		IsDel:              req.IsDel,
		UID:                req.UID,
		UIDType:            req.UIDType,
		TaskID:             req.TaskID,
		Biz:                req.Biz,
		ProgressCurrentNum: req.ProgressCurrentNum,
		Status:             req.Status,
		StartTime:          req.StartTime,
		EndTime:            req.EndTime,
		ExtJSON:            req.ExtJSON,
	}, scope)
	return
}

func DeleteTskUserTaskRel(ctx context.Context, id int64) (err error) {
	_, err = tsk_user_task_rel.DeleteTskUserTaskRel(ctx, resource.GormDB, id)
	return
}

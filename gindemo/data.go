/*
 * @Author: tj
 * @Date: 2022-11-07 14:08:27
 * @LastEditors: tj
 * @LastEditTime: 2022-11-07 14:32:12
 * @FilePath: \createApiMarkdown\gindemo\data.go
 */
package gindemo

type Result struct {
	Code int         `json:"code" validate:"required,0-成功 1-失败"`
	Data interface{} `json:"data" validate:"required"`
	Msg  string      `json:"msg" validate:"required,错误信息"`
}

type ActivityListRequest struct {
	Page   int `json:"page" validate:"required,页码：从1开始"`
	Status int `json:"status" validate:"required,活动状态: -1:表示查全部;1-准备;2-关闭,oneof=-1 1 2"`
}

type ActivityListDetail struct {
	Aid       uint64 `json:"aid" validate:"required,活动id"`
	TitleName string `json:"titleName" validate:"required,活动标题"`
	StartTime string `json:"startTime" validate:"required,活动开始时间"`
	EndTime   string `json:"endTime" validate:"required,活动结束时间"`
	Status    uint8  `json:"status" validate:"required,活动状态"`
}

type GoodName struct {
	Gid       uint64 `json:"gid" validate:"required,商品id"`
	GoodsName string `json:"goodsName" validate:"required,商品名称"`
}

type ActivityListResponse struct {
	TotalCount int64                `json:"totalCount" validate:"required,总数"`
	Details    []ActivityListDetail `json:"details" validate:"required,活动详情列表"`
	Goods      []GoodName           `json:"goods" validate:"required,商品详情列表"`
}

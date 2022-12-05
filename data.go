/*
 * @Author: tj
 * @Date: 2022-11-01 19:21:07
 * @LastEditors: tj
 * @LastEditTime: 2022-12-05 10:17:20
 * @FilePath: \createApiMarkdown\data.go
 */
package main

type UpdateActivityRequest struct {
	Aid        int    `json:"aid" validate:"required,活动id"`
	IsTop      int    `json:"isTop" validate:"required,是否置顶,oneof=0 1"`
	IsValid    int    `json:"isValid" validate:"required,是否有效,oneof=0 1"`
	IsExtend   int    `json:"isExtend" validate:"required,是否延长时间(0:否;1:是;),oneof=0 1"`
	ManualTime int    `json:"manualTime" validate:"手动延迟时间(单位小时)"`
	AiNum      int    `json:"aiNum" validate:"required,Ai人数"`
	Picture    string `json:"picture" validate:"图片存放的url"`
	StartTime  string `json:"startTime" validate:"required,开始时间"`
	EndTime    string `json:"endTime" validate:"required,结束时间"`
}

type Result struct {
	Code int         `json:"code" validate:"required,0-成功 1-失败"`
	Data interface{} `json:"data" validate:"required"`
	Msg  string      `json:"msg" validate:"required,错误信息"`
}

type ActivityListResponse struct {
	TotalCount int64                 `json:"totalCount" validate:"required,总数"`
	Counts     []int64               `json:"counts" validate:"required,数量"`
	Names      []string              `json:"names" validate:"required,名称"`
	Details    []*ActivityListDetail `json:"details" validate:"required,详细信息"`
	Goods      []*GoodName           `json:"goods" validate:"required,有库存的商品信息"`
}

type ActivityListDetail struct {
	Aid       uint64   `json:"aid" validate:"required,活动id"`
	TitleName []string `json:"titleName" validate:"required,活动标题"`
	StartTime string   `json:"startTime" validate:"required,活动开始时间"`
	EndTime   string   `json:"endTime" validate:"required,活动结束时间"`
	AiNum     uint     `json:"aiNum" validate:"required,Ai人数"`
	Status    uint8    `json:"status" validate:"required,活动状态"`
}

type GoodName struct {
	Gid        uint64  `json:"gid"  validate:"required,商品id"`
	GoodsName  string  `json:"goodsName"  validate:"required,商品名称"`
	GoodsPrice float64 `json:"goodsPrice"  validate:"required,商品价格"`
}

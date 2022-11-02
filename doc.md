# 接口文档
> 版本号：1.0<br>
> BaseUrl: api/


## 1. 获取列表
> 作者：1260

### 请求说明
> 请求方式：POST<br>
请求URL ：[activity/list](#)

#### 请求参数

|字段      |字段类型       |必填     |字段说明    |
|:--:|:--:|:--:|:--:|
|aid   |int        |是      |活动id    |
|isTop   |int        |是      |是否置顶,oneof=0 1    |
|isValid   |int        |是      |是否有效,oneof=0 1    |
|isExtend   |int        |是      |是否延长时间(0:否;1:是;),oneof=0 1    |
|manualTime   |int        |否      |手动延迟时间(单位小时)    |
|aiNum   |int        |是      |Ai人数    |
|picture   |string        |否      |图片存放的url    |
|startTime   |string        |是      |开始时间    |
|endTime   |string        |是      |结束时间    |


### 返回结果
```json
 {
	"code": 0,
	"data": {
		"totalCount": 100,
		"counts": [
			123,
			456,
			789
		],
		"names": [
			"123",
			"456",
			"789"
		],
		"details": [
			{
				"aid": 1,
				"titleName": [
					"手机"
				],
				"startTime": "2022-11-1 17:31:10",
				"endTime": "2022-11-1 17:31:17",
				"aiNum": 0,
				"status": 1
			}
		],
		"goods": [
			{
				"gid": 1,
				"goodsName": "GoodsName1",
				"goodsPrice": 125.5
			}
		]
	},
	"msg": ""
} 
```
### 返回参数

|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|code   |int        |0-成功 1-失败    |
|data   |interface        | [rsp-data](#rsp-1.1.data)    |
|msg   |string        |错误信息    |

<a id="rsp-1.1.data"></a> 
##### rsp-data 
 
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|totalCount   |int64        |总数    |
|counts   |slice[int64]        |数量    |
|names   |slice[string]        |名称    |
|details   |slice[struct]        |详细信息 [rsp-details](#rsp-1.1.details)    |
|goods   |slice[struct]        |有库存的商品信息 [rsp-goods](#rsp-1.1.goods)    |
 
<a id="rsp-1.1.details"></a> 
##### rsp-details 
 
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|aid   |uint64        |活动id    |
|titleName   |slice[string]        |活动标题    |
|startTime   |string        |活动开始时间    |
|endTime   |string        |活动结束时间    |
|aiNum   |uint        |Ai人数    |
|status   |uint8        |活动状态    |
 
<a id="rsp-1.1.goods"></a> 
##### rsp-goods 
 
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|gid   |uint64        |商品id    |
|goodsName   |string        |商品名称    |
|goodsPrice   |float64        |商品价格    |
 



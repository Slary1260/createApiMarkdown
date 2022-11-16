# 接口文档
> 版本号：1.0<br>
> BaseUrl: <br>
> ApiList: 
[1.activity-add](#1.activity-add)&emsp;[2.activity-close](#2.activity-close)&emsp;[3.activity-info](#3.activity-info)&emsp;[4.activity-list](#4.activity-list)&emsp;[5.activity-update](#5.activity-update)&emsp;
[6.api-gethtml](#6.api-gethtml)<br>

<a id="1.activity-add"></a>
## 1. activity-add
> 作者：

### 请求说明
> 请求方式：POST<br>
请求URL ：[/activity/add](#)

#### 请求参数


### 返回结果
```json
 {
	"code": 0,
	"data": null,
	"msg": ""
} 
```
### 返回参数

|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|code   |int        |0-成功 1-失败    |
|data   |interface        |    |
|msg   |string        |错误信息    |



<a id="2.activity-close"></a>
## 2. activity-close
> 作者：

### 请求说明
> 请求方式：POST<br>
请求URL ：[/activity/close](#)

#### 请求参数


### 返回结果
```json
 {
	"code": 0,
	"data": null,
	"msg": ""
} 
```
### 返回参数

|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|code   |int        |0-成功 1-失败    |
|data   |interface        |    |
|msg   |string        |错误信息    |



<a id="3.activity-info"></a>
## 3. activity-info
> 作者：

### 请求说明
> 请求方式：POST<br>
请求URL ：[/activity/info](#)

#### 请求参数


### 返回结果
```json
 {
	"code": 0,
	"data": {},
	"msg": ""
} 
```
### 返回参数

|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|code   |int        |0-成功 1-失败    |
|data   |interface        |    |
|msg   |string        |错误信息    |



<a id="4.activity-list"></a>
## 4. activity-list
> 作者：

### 请求说明
> 请求方式：POST<br>
请求URL ：[/activity/list](#)

#### 请求参数

|字段      |字段类型       |必填     |字段说明    |
|:--:|:--:|:--:|:--:|
|details   |slice[struct]        |是      |活动详情列表 [req-details](#req-4.1.details)    |
|page   |int        |是      |页码：从1开始    |
|status   |int        |是      |活动状态: -1:表示查全部;1-准备;2-关闭,oneof=-1 1 2    |

<a id="req-4.1.details"></a> 
##### req-details 
 
|字段      |字段类型       |必填     |字段说明    |
|:--:|:--:|:--:|:--:|
|aid   |uint64        |是      |活动id    |
|titleName   |string        |是      |活动标题    |
|startTime   |string        |是      |活动开始时间    |
|endTime   |string        |是      |活动结束时间    |
|status   |uint8        |是      |活动状态    |
 

### 返回结果
```json
 {
	"code": 0,
	"data": {
		"totalCount": 0,
		"details": [
			{
				"aid": 0,
				"titleName": "",
				"startTime": "",
				"endTime": "",
				"status": 0
			}
		],
		"goods": [
			{
				"gid": 0,
				"goodsName": ""
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
|data   |interface        | [rsp-data](#rsp-4.1.data)    |
|msg   |string        |错误信息    |

<a id="rsp-4.1.data"></a> 
##### rsp-data 
 
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|totalCount   |int64        |总数    |
|details   |slice[struct]        |活动详情列表 [rsp-details](#rsp-4.1.details)    |
|goods   |slice[struct]        |商品详情列表 [rsp-goods](#rsp-4.1.goods)    |
 
<a id="rsp-4.1.details"></a> 
##### rsp-details 
 
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|aid   |uint64        |活动id    |
|titleName   |string        |活动标题    |
|startTime   |string        |活动开始时间    |
|endTime   |string        |活动结束时间    |
|status   |uint8        |活动状态    |
 
<a id="rsp-4.1.goods"></a> 
##### rsp-goods 
 
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|gid   |uint64        |商品id    |
|goodsName   |string        |商品名称    |
 


<a id="5.activity-update"></a>
## 5. activity-update
> 作者：

### 请求说明
> 请求方式：POST<br>
请求URL ：[/activity/update](#)

#### 请求参数


### 返回结果
```json
 {
	"code": 0,
	"data": null,
	"msg": ""
} 
```
### 返回参数

|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
|code   |int        |0-成功 1-失败    |
|data   |interface        |    |
|msg   |string        |错误信息    |



<a id="6.api-gethtml"></a>
## 6. api-gethtml
> 作者：

### 请求说明
> 请求方式：GET<br>
请求URL ：[/api/gethtml](#)

#### 请求参数


### 返回结果

### 返回参数




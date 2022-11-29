<!--
 * @Author: tj
 * @Date: 2022-11-02 12:02:53
 * @LastEditors: tj
 * @LastEditTime: 2022-11-29 14:19:00
 * @FilePath: \createApiMarkdown\README.md
-->
# API接口markdown文档生成工具
通过api的请求对象、返回对象，自动生成对应的API接口文档。
使得开发人员接口开发完成时，或接口维护后，无需再更新api文档。

## 编写工具的原因
目前使用比较多的是swagger，但是swagger依旧有一些缺点。
来自网友的总结：
```
如果是与服务端代码集成，直接嵌入到工程代码中，侵入性比较大。将文档参数和应用参数杂糅在一起，不易阅读，而且比较依赖于项目，无法独立部署。项目挂掉，文档也无法访问。给后期代码维护增加难度。
如果直接编辑json文档，则难度比较大。即使是官网的在线编辑，功能也比较弱，提示功能差劲。很多时候在编辑预览中没问题，导出来部署就显示不正常。而且不支持多人编辑，只能一次一个人改，部署相当不方便。
用户体验，无论请求还是响应，无法方便的输入自定义json格式。特别是多层嵌套，异常繁琐。
```

我个人觉得swagger生成的页面展示不够直观。需要点击相应的按钮，才能查看接口的参数与返回值。

因此在日常工作中，需要一个简单、便捷的API接口生成工具。根据接口的定义，自动生成对应的接口文档，且后续维护不再需要手动更新接口文档。于是开始编写该工具。

## 1.使用
```
doc := document.NewDocument("api/")

// item 接口对象
err := doc.AddDocItem(item)
if err != nil {
    log.Panic(err)
}

md := markdown.New(doc, markdown.WithMd2Html(true))
err = md.Generate("doc.md")
if err != nil {
    log.Panic(err)
}

log.Infoln("success")
```

## 2.特性
### 2.1.嵌套层数
请求对象和返回对象目前只支持三层嵌套(暂不支持指针对象嵌套):Result->ActivityListResponse->ActivityListDetail
```
type Result struct {
	Data interface{} `json:"data" validate:"required"`
}

type ActivityListResponse struct {
	Details    []ActivityListDetail `json:"details" validate:"required,详细信息"`
}

type ActivityListDetail struct {
}
```

### 2.2.支持类型
请求对象、返回对象：指针、结构体和切片(slice)。
对象字段：结构体、匿名结构体、有符号整型、无符号整型、字符串、bool，浮点型(float32,float64)

### 2.3.可自定义markdown tag 标识(默认值为validate)
例：
修改前：
```
    TotalCount int64 `json:"totalCount" validate:"required,总数"`
```
修改后：
```
    TotalCount int64 `json:"totalCount" doc:"required,总数"`
```
代码设置
```
    doc := document.NewDocument("接口文档", "1.0", "api/")
	doc.SetMdKey("doc")
```

### 2.4.支持markdown转html
```
md := markdown.New(doc, markdown.WithMd2Html(true))
```

### 2.5.支持传入解析好的API结构体数据，避免再次解析结构体
```
doc := document.NewDocument("api/", document.WithParseReq(false), document.WithParseRsq(false))

reqFields, err := doc.ParseReqOrRsp(req)
if err != nil {
    log.Panic(err)
}

rspFields, err := doc.ParseReqOrRsp(rsp)
if err != nil {
    log.Panic(err)
}

item.ReqFields = reqFields
item.RspFields = rspFields
err = doc.AddDocItem(item)
if err != nil {
    log.Panic(err)
}
```

## 3.gin自动注册路由
根据gin自动注册的路由参数，自动生成API接口文档

## 4.参考项目
https://github.com/w3liu/gendoc 

<!--
 * @Author: tj
 * @Date: 2022-11-02 12:02:53
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 13:05:02
 * @FilePath: \createApiMarkdown\README.md
-->
# createApiMarkdown
通过api的请求对象、返回对象，自动生成对应的API接口文档。 使得开发人员专注业务开发，节省了开发人员的api文档修改时间。

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
请求对象和返回对象目前只支持三层嵌套:Result->ActivityListResponse->ActivityListDetail
```
type Result struct {
	Data interface{} `json:"data" validate:"required"`
}

type ActivityListResponse struct {
	Details    []*ActivityListDetail `json:"details" validate:"required,详细信息"`
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

## 3.TODO
自动获取gin路由生成接口文档

## 4.参考项目
https://github.com/w3liu/gendoc 

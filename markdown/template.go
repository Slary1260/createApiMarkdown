/*
 * @Author: tj
 * @Date: 2022-10-26 18:07:40
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 12:17:54
 * @FilePath: \createApiMarkdown\markdown\template.go
 */
package markdown

const TplPage = `# {title}
> 版本号：{version}<br>
> BaseUrl: {Url}
{body}
`

const TplBody = `

## {id}. {name}
> 作者：{author}

### 请求说明
> 请求方式：{method}<br>
请求URL ：[{url}](#)

#### 请求参数
{reqTable}

### 返回结果
{respParam}
### 返回参数
{respTable}

`

const TplReqTable = `
|字段      |字段类型       |必填     |字段说明    |
|:--:|:--:|:--:|:--:|
{params}`

const TplReqParam = `|{name}   |{kind}        |{required}      |{description}    |
`

const TplRespTable = `
|字段      |字段类型       |字段说明    |
|:--:|:--:|:--:|
{params}`

const TplRespParam = `|{name}   |{kind}        |{description}    |
`

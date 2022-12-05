/*
 * @Author: tj
 * @Date: 2022-11-02 11:05:05
 * @LastEditors: tj
 * @LastEditTime: 2022-12-05 10:14:05
 * @FilePath: \createApiMarkdown\markdown\option.go
 */
package markdown

type Option func(*Markdown)

func WithMd2Html(isMd2Html bool) Option {
	return func(m *Markdown) {
		m.is2html = isMd2Html
	}
}

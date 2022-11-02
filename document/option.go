/*
 * @Author: tj
 * @Date: 2022-11-02 11:04:39
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 12:18:10
 * @FilePath: \createApiMarkdown\document\option.go
 */
package document

type Option func(*Document)

func WithTitle(title string) Option {
	return func(d *Document) {
		d.Title = title
	}
}

func WithVersion(version string) Option {
	return func(d *Document) {
		d.Version = version
	}
}

// 设置markdown鉴别标识
func WithMdKey(key string) Option {
	return func(d *Document) {
		d.mdKey = key
	}
}

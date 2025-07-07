package model

// StringArray 用于处理 JSON 数组类型的字符串
// 实现了 driver.Valuer 和 sql.Scanner 接口
type Wallpaper struct {
	BaseModel              // ✅ 匿名嵌套，自动带入字段
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Type      string   `json:"type"`      // "single" 或 "group"
	ImageUrls StringArray `json:"imageUrls"` // 一个或多个图
	Author    string   `json:"author"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`      // 标签列表
	Thumbnail string   `json:"thumbnail"` // 缩略图 URL
}

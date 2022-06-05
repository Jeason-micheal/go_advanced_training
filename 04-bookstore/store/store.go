package store

type Book struct {
	Id      string   `json:"id"` // 图书ISBN ID
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"` //出版社
}

type Store interface {
	Create(*Book) error      // 创建一个新图书条目
	Update(*Book) error      // 更新某图书条目
	GetAll() ([]Book, error) // 获取所有图书信息
}

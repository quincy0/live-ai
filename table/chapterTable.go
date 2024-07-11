package table

type ChapterTable struct {
	ChapterId    int64  `json:"chapter_id" gorm:"PRIMARY_KEY;column:chapter_id"`
	ChapterTitle string `json:"chapter_title" gorm:"column:chapter_title"`
	ScriptTag    string `json:"script_tag" gorm:"column:script_tag"`
	ProductTag   string `json:"product_tag" gorm:"column:product_tag"`
	Summary      string `json:"summary" gorm:"column:summary"`
}

func (p *ChapterTable) TableName() string {
	return "chapter"
}

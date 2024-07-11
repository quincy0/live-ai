package table

type ParagraphTable struct {
	ParagraphId    int64  `json:"paragraph_id" gorm:"PRIMARY_KEY;column:paragraph_id"`
	ChapterId      int64  `json:"chapter_id" gorm:"column:chapter_id"`
	ParagraphTitle string `json:"paragraph_title" gorm:"column:paragraph_title"`
	Content        string `json:"content" gorm:"column:content"`
}

func (t *ParagraphTable) TableName() string {
	return "paragraph"
}

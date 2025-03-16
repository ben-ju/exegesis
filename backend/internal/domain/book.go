package domain

type BibleVersion struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Version  string `json:"version"`
	Language string `json:"language"`
	FilePath string `json:"-"`
}

type Verse struct {
	OsisID  string `json:"osis_id"`
	Content string `json:"content"`
}

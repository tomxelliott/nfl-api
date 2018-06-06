package logic

type Team struct {
	Id   string
	Name string
	Info TeamData
}

type TeamData struct {
	Coach      string
	NoOfTitles int
	Stadium    string
}

type PageData struct {
	PageTitle string
	Teams     []Team
}

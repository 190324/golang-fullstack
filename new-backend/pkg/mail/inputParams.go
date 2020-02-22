package mail

type InputParams struct {
	Name        string // 收件人姓名
	Email       string // 收件人 email
	Subject     string // email 主旨
	Template    string // email 版型
	ContentData Data   // email 變數
}

type Data struct {
	Title   string
	Content string
}

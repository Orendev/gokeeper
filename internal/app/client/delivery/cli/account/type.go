package account

type CreateAccountArgs struct {
	Title    string `json:"title"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

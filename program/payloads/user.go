package payloads

type User struct {
	UserId    int64      `json:"id"`
	Username  string     `json:"username"`
	Posts     []Post     `json:"posts"`
	Followed  []UserLike `json:"followed"`
	Followers []UserLike `json:"followers"`
}

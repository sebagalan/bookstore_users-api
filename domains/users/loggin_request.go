package users

type LogginRequest struct {
	Email    string "json:email"
	Password string "json:password"
}

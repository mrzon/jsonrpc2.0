package user_common

type User struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Password string `json:"password"`
}
type RegisterUserSpec struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUserSpec struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterUserResult struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type LoginUserResult struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type UserService struct {
	Register func(spec RegisterUserSpec) RegisterUserResult;
	Login    func(spec LoginUserSpec) LoginUserResult;
}


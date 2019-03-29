package user_server

import (
	"github.com/mrzon/jsonrpc2.0/example/user/user-common"
)

type UserServiceImpl struct {
	user_common.UserService
	registeredUserMap map[string]*user_common.User
	loggedInUserMap   map[string]*user_common.User
}

func NewUserServiceImpl() UserServiceImpl {
	impl := &UserServiceImpl{
		registeredUserMap: make(map[string]*user_common.User),
		loggedInUserMap:   make(map[string]*user_common.User),
	}
	impl.Register = func(spec user_common.RegisterUserSpec) user_common.RegisterUserResult {
		len := len(impl.registeredUserMap)
		newId := len + 1;
		user := &user_common.User{
			Name:     spec.Name,
			Password: spec.Password,
			Id:       newId,
		}
		impl.registeredUserMap[spec.Name] = user

		return user_common.RegisterUserResult{
			Status: "Success",
			Code:   200,
		}
	}

	impl.Login = func(spec user_common.LoginUserSpec) user_common.LoginUserResult {
		user := impl.registeredUserMap[spec.Name]
		if user == nil {
			return user_common.LoginUserResult{
				Status: "User not found",
				Code:   404,
			}
		}
		if user.Password != spec.Password {
			return user_common.LoginUserResult{
				Status: "You have wrong credential",
				Code:   403,
			}
		}
		impl.loggedInUserMap[spec.Name] = user
		return user_common.LoginUserResult{
			Status: "Success Logged in",
			Code:   200,
		}

	}
	return *impl
}

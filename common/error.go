package common

type HttpStatus interface {
	Status() int
}

type UsernameExist struct {
}

func (u UsernameExist) Error() string {
	return "用户名已存在"
}

func (u UsernameExist) Status() int {
	return 400
}

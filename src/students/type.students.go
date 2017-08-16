package students

import "github.com/fadhilimamk/ganeca/src/global"

// Student struct represent a student in campus
type Student struct {
	fullname string
	nim      int
	major    int
}

func (p Student) ToString() string {
	return string(global.ObjectToJSON(p))
}

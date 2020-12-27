package model

import "fmt"

type Calculator struct {
	Expression string `json:"expression,omitEmpty"`
	Result     string `json:"result,omitEmpty"`
}

func (cal Calculator) String() string {
	return fmt.Sprintf(`%s = %s`, cal.Expression, cal.Result)
}

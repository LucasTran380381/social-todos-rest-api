package shared

type Paging struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
	Total int64
}

func (paging *Paging) Format() {
	if paging.Page < 1 {
		paging.Page = 1
	}

	if paging.Limit < 3 || paging.Limit > 10 {
		paging.Limit = 3
	}
}

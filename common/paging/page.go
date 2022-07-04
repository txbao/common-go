package paging

type PageStruct struct {
	Count int64 `json:"count"`
}

type PageResStruct struct {
	Count     int64       `json:"count"`
	TotalPage int64       `json:"totalPage"`
	Limit     int64       `json:"limit"`
	Page      int64       `json:"page"`
	List      interface{} `json:"list"`
}

// 获取分页信息
func GetPaging(limit int64, page int64) (int64, int64, int64) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	//总页数
	offset := (page - 1) * limit
	return limit, offset, page
}

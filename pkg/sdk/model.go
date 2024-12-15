package sdk

type SegmentResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data SegmentData `json:"data"`
}

type SegmentData struct {
	MinID int64 `json:"min_id"`
	MaxID int64 `json:"max_id"`
	Step  int64 `json:"step"`
}

type SnowflakeResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data SnowflakeData `json:"data"`
}

type SnowflakeData struct {
	Total int64   `json:"total"`
	List  []int64 `json:"list"`
}

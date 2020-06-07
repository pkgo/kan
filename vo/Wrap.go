package vo

type Wrap struct {
	Err    bool        `json:"err"`
	ErrMsg string      `json:"errMsg"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(data interface{}) (vw Wrap) {
	vw.Err = false
	vw.Data = data
	vw.ErrMsg = ""
	vw.Status = "success"
	return vw
}

func Error(errMsg string) (vw Wrap) {
	vw.Err = true
	vw.Data = nil
	vw.ErrMsg = errMsg
	vw.Status = "error"
	return vw
}

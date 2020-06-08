package vo

// Wrap is a vo interface
type Wrap struct {
	Ok       bool        `json:"ok"`
	Err      bool        `json:"err"`
	ErrMsg   string      `json:"errMsg"`
	Status   string      `json:"status"`
	Duration int64       `json:"duration"`
	Data     interface{} `json:"data"`
}

func (vw *Wrap) SetDuration(duration int64) Wrap {
	vw.Duration = duration
	return *vw
}

// Success will new a Wrap struct with success
func Success(data interface{}) (vw Wrap) {
	vw.Ok = true
	vw.Err = false
	vw.Data = data
	vw.ErrMsg = ""
	vw.Status = "success"
	return vw
}

// Error will new a Wrap struct with error
func Error(errMsg string) (vw Wrap) {
	vw.Ok = false
	vw.Err = true
	vw.Data = nil
	vw.ErrMsg = errMsg
	vw.Status = "error"
	return vw
}

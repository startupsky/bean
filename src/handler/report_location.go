package handler

import "net/http"
import "storage"
import "time"
import "util"
import "model"
import "encoding/json"
import "def"
import "strconv"
import "runtime/debug"

//import "io/ioutil"
type playerReportResponse struct {
	ErrNo   int             `json:"errno"`
	ErrMsg  string          `json:"errmsg"`
	Players []*model.Player `json:"players"`
	Beans   []*model.Bean   `json:"beans"`
}

func PlayerReport(rw http.ResponseWriter, req *http.Request) {
	errNo := 0
	errMsg := ""
	startTick := time.Now()

	session := util.Session()

	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case *def.BusinessException:
				errNo = x.(*def.BusinessException).ErrNo
				errMsg = x.(*def.BusinessException).ErrMsg
			default:
				errMsg = def.UnExpectedErrMsg
				errNo = def.UnExpectedErrNo
				logger.Error("session=[%s] unexpected exception=[%v] stack=[%s]", session, x, string(debug.Stack()))
			}
		}
		response := playerReportResponse{errNo, errMsg, storage.Ele.Players, storage.Ele.Beans}

		encoder := json.NewEncoder(rw)
		encoder.Encode(response)

		endTick := time.Now()
		logger.Info("Access ReportLocation session=[%s] errno=[%d] errmsg=[%s] duration=[%d] id=[%s] longitude=[%s] latitude=[%s]",
			session,
			errNo,
			errMsg,
			int64(endTick.Sub(startTick)/time.Millisecond),
			req.FormValue("id"),
			req.FormValue("longitude"),
			req.FormValue("latitude"))
	}()

	id, err1 := strconv.ParseUint(req.FormValue("id"), 10, 64)
	longitude, err2 := strconv.ParseInt(req.FormValue("longitude"), 10, 64)
	latitude, err3 := strconv.ParseInt(req.FormValue("latitude"), 10, 64)

	if err1 == nil && err2 == nil && err3 == nil {
		storage.Ele.PlayerReport(id, longitude, latitude)

	} else {
		panic(def.ParamException)
	}
}

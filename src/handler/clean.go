package handler

import "net/http"
import "storage"
import "time"
import "util"
import "encoding/json"
import "def"
import "strconv"
import "runtime/debug"

//import "io/ioutil"
type cleanResponse struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
}

func Clean(rw http.ResponseWriter, req *http.Request) {
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
		response := cleanResponse{errNo, errMsg}

		encoder := json.NewEncoder(rw)
		encoder.Encode(response)

		endTick := time.Now()
		logger.Info("Access Clean session=[%s] errno=[%d] errmsg=[%s] duration=[%d] type=[%s]",
			session,
			errNo,
			errMsg,
			int64(endTick.Sub(startTick)/time.Millisecond),
			req.FormValue("type"))
	}()

	typ, err := strconv.ParseUint(req.FormValue("type"), 10, 8)

	if err == nil {
		if typ&1 > 0 {
			storage.Ele.CleanPlayers()
		}
		if typ&2 > 0 {
			storage.Ele.CleanBeans()
		}

	} else {
		panic(def.ParamException)
	}
}

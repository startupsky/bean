package handler

import "github.com/nporsche/np-golang-logging"
import "net/http"
import "storage"
import "time"
import "util"
import "model"
import "encoding/json"
import "def"
import "strconv"
import "runtime/debug"

var logger = logging.MustGetLogger("handler")

type beanManipulateResponse struct {
	ErrNo   int             `json:"errno"`
	ErrMsg  string          `json:"errmsg"`
	Players []*model.Player `json:"players"`
	Beans   []*model.Bean   `json:"beans"`
}

func BeanManipulate(rw http.ResponseWriter, req *http.Request) {
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
		response := beanManipulateResponse{errNo, errMsg, storage.Ele.Players, storage.Ele.Beans}

		encoder := json.NewEncoder(rw)
		encoder.Encode(response)

		endTick := time.Now()
		logger.Info("Access BeanManipulate session=[%s] errno=[%d] errmsg=[%s] duration=[%d] id=[%s] state=[%s] longitude=[%s] latitude=[%s]",
			session,
			errNo,
			errMsg,
			int64(endTick.Sub(startTick)/time.Millisecond),
			req.FormValue("id"),
			req.FormValue("state"),
			req.FormValue("longitude"),
			req.FormValue("latitude"))
	}()

	id, err1 := strconv.ParseUint(req.FormValue("id"), 10, 64)
	state, err2 := strconv.ParseUint(req.FormValue("state"), 10, 8)
	longitude, err3 := strconv.ParseInt(req.FormValue("longitude"), 10, 64)
	latitude, err4 := strconv.ParseInt(req.FormValue("latitude"), 10, 64)

	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		storage.Ele.BeanManipulate(id, uint8(state), longitude, latitude)

	} else {
		panic(def.ParamException)
	}
}

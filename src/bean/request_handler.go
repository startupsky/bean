package main

import (
	"domains"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type RequestHandler struct {
	domainMgr *domains.DomainManager
}

/*
parameters:
	1. domain
	2. guest ip

return:
	1: yes
	0: no
*/
func NewRequestHandler(path string) *RequestHandler {
	this := new(RequestHandler)
	this.domainMgr = domains.NewDomainManager(path)

	return this
}

var InvalidParameterException = errors.New("InvalidParameter")
var DomainNotExistedException = errors.New("Domain does not existed")
var IPNotExistedException = errors.New("IP does not existed in this domain")
var NoBackendServersException = errors.New("No Backend servers")

func (this *RequestHandler) ListAuthenticator(args []string) (data []string, err error) {
	if len(args) != 1 {
		return nil, InvalidParameterException
	}
	domain := strings.ToLower(args[0])
	dobj := this.domainMgr.GetDomainObj(domain)
	if dobj == nil {
		return nil, DomainNotExistedException
	}
	data = dobj.WhiteList
	log.Debug("args=[%v] data=[%v]", args, data)
	return dobj.WhiteList, nil
}

/*
parameters:
	1. Host

return:
	1: IP_address
*/
func (this *RequestHandler) IP(args []string) (data []string, err error) {
	if len(args) != 1 {
		return nil, InvalidParameterException
	}
	domain := strings.ToLower(args[0])

	dobj := this.domainMgr.GetDomainObj(domain)
	if dobj == nil {
		return nil, DomainNotExistedException
	}
	if len(dobj.Servers) == 0 {
		return nil, NoBackendServersException
	}
	var ip string
	rate := int32(-65535)
	i := 1
	for _, svr := range dobj.Servers {
		svrRate := svr.Weight + svr.Score
		if svrRate > rate {
			ip = svr.IP
			rate = svrRate
			i = 1
		} else if svrRate == rate {
			// 1/i possibility to choose this one
			i++
			if rand.Intn(i) == 0 {
				ip = svr.IP
			}
		}
	}
	data = []string{ip}
	log.Debug("args=[%v] data=[%v]", args, data)
	return
}

/*
parameters:
	1. Host
	2. IP
	3. ScoreDelta

return:
	1: IP
	2: Score
*/
func (this *RequestHandler) ScoreIP(args []string) (data []string, err error) {
	if len(args) != 3 {
		return nil, InvalidParameterException
	}
	domain := strings.ToLower(args[0])
	ip := args[1]
	weightDelta, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		return nil, InvalidParameterException
	}

	dobj := this.domainMgr.GetDomainObj(domain)
	if dobj == nil {
		return nil, DomainNotExistedException
	}
	for _, svr := range dobj.Servers {
		if svr.IP == ip {
			score := atomic.AddInt32(&svr.Score, int32(weightDelta))
			svr.ScoreTm = time.Now()
			data = []string{ip, fmt.Sprint(score)}
			break
		}
	}
	if data == nil {
		return nil, IPNotExistedException
	}

	log.Debug("args=[%v] data=[%v]", args, data)
	return
}

/*
parameters:
1. Host

return:
1: Connect
2: Read
3: Write
*/
func (this *RequestHandler) Timeout(args []string) (data []string, err error) {
	if len(args) != 1 {
		return nil, InvalidParameterException
	}

	domain := strings.ToLower(args[0])
	dobj := this.domainMgr.GetDomainObj(domain)
	if dobj == nil {
		return nil, DomainNotExistedException
	}

	data = []string{dobj.Timeout.Connect, dobj.Timeout.Read, dobj.Timeout.Write}
	log.Debug("args=[%v] data=[%v]", args, data)
	return
}

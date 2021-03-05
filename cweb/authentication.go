package cweb

import (
	"github.com/gin-gonic/gin"
	http2 "github.com/waittttting/cRPC-common/http"
)

const (
	callerService = "service"
	callerUser = "user"
)

func getHeaderFromCtx(ctx *gin.Context) (*http2.RequestHeader, bool, string) {

	header := ctx.Request.Header
	reqHeader := &http2.RequestHeader{}
	var hasParamsErr = false
	var paramsErrInfo = ""


	if header[http2.HttpHeaderStringDomainId] == nil {
		hasParamsErr = true
		paramsErrInfo = paramsErrInfo + " " + "not find domain id in header"
	} else {
		reqHeader.DomainID = header[http2.HttpHeaderStringDomainId][0]
	}
	if header[http2.HttpHeaderStringAppId] == nil {
		hasParamsErr = true
		paramsErrInfo = paramsErrInfo + " " + "not find app id in header"
	} else {
		reqHeader.AppID = header[http2.HttpHeaderStringAppId][0]
	}

	if header[http2.HttpHeadrStringCaller] == nil {
		hasParamsErr = true
		paramsErrInfo = paramsErrInfo + " " + "not find caller in header"
	} else {
		reqHeader.Caller = header[http2.HttpHeadrStringCaller][0]
	}

	switch reqHeader.Caller {
	case callerService:
		if header[http2.HttpHeadrStringSecretKey] == nil {
			hasParamsErr = true
			paramsErrInfo = paramsErrInfo + " " + "not find secret key id in header"
		} else {
			reqHeader.SecretKey = header[http2.HttpHeadrStringSecretKey][0]
		}

		reqHeader.Uid = "0" // grpc 会强转这个字段，导致报错，所以设置为0
	case callerUser:
		if header[http2.HttpHeaderStringUid] == nil {
			hasParamsErr = true
			paramsErrInfo = paramsErrInfo + " " + "not find user id in header"
		} else {
			reqHeader.Uid = header[http2.HttpHeaderStringUid][0]
		}

		if header[http2.HttpHeadrStringAccessToken] == nil {
			hasParamsErr = true
			paramsErrInfo = paramsErrInfo + " " + "not find access token id in header"
		} else {
			reqHeader.AccessToken = header[http2.HttpHeadrStringAccessToken][0]
		}
	default:
		hasParamsErr = true
		return nil, true, "Caller was wrong"
	}

	if header[http2.HttpHeadrStringCommandId] == nil {
		hasParamsErr = true
		paramsErrInfo = paramsErrInfo + " " + "not find command id id in header"
	} else {
		reqHeader.CommandID = header[http2.HttpHeadrStringCommandId][0]
	}

	if header[http2.HttpHeadrStringCommandVersion] == nil {
		hasParamsErr = true
		paramsErrInfo = paramsErrInfo + " " + "not find command version id in header"
	} else {
		reqHeader.CommandVersion = header[http2.HttpHeadrStringCommandVersion][0]
	}

	if hasParamsErr {
		return nil, hasParamsErr, paramsErrInfo
	}

	return reqHeader, false, ""
}

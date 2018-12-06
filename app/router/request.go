package router

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-requestbin/app/constant"
	"github.com/isayme/go-requestbin/app/model"
	"github.com/isayme/go-requestbin/app/schema"
)

var globalRequest = &model.Request{}

// RecordRequest record a request
func RecordRequest(c *gin.Context) {
	request := c.Request

	contentLength, _ := strconv.Atoi(c.GetHeader(constant.HeaderContentLength))

	requestInfo := &schema.RequestInfo{
		Method:        request.Method,
		Path:          request.RequestURI,
		IP:            c.ClientIP(),
		ContentType:   c.ContentType(),
		ContentLength: contentLength,
	}

	// header
	headers := schema.M{}
	for k := range request.Header {
		headers[k] = request.Header.Get(k)
	}
	requestInfo.Headers = headers

	// query
	query := schema.M{}
	for k, v := range request.URL.Query() {
		if len(v) > 1 {
			query[k] = v
		} else {
			query[k] = v[0]
		}
	}
	requestInfo.Query = query

	// body
	rawBody, _ := ioutil.ReadAll(request.Body)
	requestInfo.Data = string(rawBody)

	contentType := request.Header.Get(constant.HeaderContentType)
	mediaType, _, _ := mime.ParseMediaType(contentType)
	if mediaType == constant.MIMEApplicationJSON { // json
		var body schema.M
		json.Unmarshal(rawBody, &body)
		requestInfo.Body = body
	} else if mediaType == constant.MIMEApplicationForm { // form
		form := schema.M{}
		formValues, _ := url.ParseQuery(string(rawBody))
		for k, v := range formValues {
			if len(v) > 1 {
				form[k] = v
			} else {
				form[k] = v[0]
			}
		}
		requestInfo.Form = form
	}

	slug := c.Param("slug")

	r, err := globalRequest.Create(c, slug, requestInfo)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, r)
}

// ListResponse list latest requests response
type ListResponse struct {
	Result []*schema.Request `json:"result"`
}

// ListRequests list requests
func ListRequests(c *gin.Context) {
	slug := c.Param("slug")

	requests, err := globalRequest.List(c, slug)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, ListResponse{
		Result: requests,
	})
}

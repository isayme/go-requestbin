package router

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/isayme/go-requestbin/app/constant"
	"github.com/isayme/go-requestbin/app/model"
	"github.com/isayme/go-requestbin/app/schema"
	"github.com/r3labs/sse/v2"
)

type Request struct {
	model     *model.Request
	sseServer *sse.Server
}

func NewRequest(model *model.Request, sseServer *sse.Server) *Request {
	return &Request{
		model:     model,
		sseServer: sseServer,
	}
}

// RecordRequest record a request
func (req *Request) RecordRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request := r
	reqHeaders := r.Header

	vars := mux.Vars(r)
	slug := vars["slug"]

	// contentLength, _ := strconv.Atoi(reqHeaders.Get(constant.HeaderContentLength))

	requestInfo := &schema.RequestInfo{
		Method:        request.Method,
		Path:          request.RequestURI,
		IP:            r.RemoteAddr,
		ContentType:   reqHeaders.Get("content-type"),
		ContentLength: r.ContentLength,
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

	record, err := req.model.Create(ctx, slug, requestInfo)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}

	req.sseServer.Publish("requests", &sse.Event{
		ID:   []byte(record.ID.Hex()),
		Data: data,
	})

	w.Write([]byte("ok"))
}

// ListResponse list latest requests response
type ListResponse struct {
	Result []*schema.Request `json:"result"`
}

// ListRequests list requests
func (req *Request) ListRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	slug := vars["slug"]

	requests, err := req.model.List(ctx, slug)
	if err != nil {
		panic(err)
	}

	if r.URL.Query().Has("pretty") {
		writeJson(w, ListResponse{
			Result: requests,
		})
	} else {
		writeJson(w, ListResponse{
			Result: requests,
		})
	}
}

func writeJson(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Write(data)
}

package constant

// CollectionRequest collection name for request
const CollectionRequest = "requests"

// HeaderUserAgent http header: User-Agent
const HeaderUserAgent = "User-Agent"

// HeaderContentType http header: Content-Type
const HeaderContentType = "Content-Type"

// HeaderContentLength http header: Content-Length
const HeaderContentLength = "Content-Length"

// MIMEApplicationJSON content type: json
const MIMEApplicationJSON = "application/json"

// MIMEApplicationForm content type: form
const MIMEApplicationForm = "application/x-www-form-urlencoded"

// SessionKey ctx 中保存 mongo.Session 的 key
type SessionKey struct{}

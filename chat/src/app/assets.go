package app

import (
	"encoding/base64"
	"strings"
)

// Content Type Map
var ctypes = map[string]string{
	"html": "text/html",
	"js":   "application/javascript",
	"css":  "text/css",
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"ico":  "image/x-icon",
}

// only images are treated as finaries
func isBinary(ctype string) bool {
	return strings.HasPrefix(ctype, "image/")
}

// Asset extract a file from the box with its content type
// returns either a content-type with "/" or an error code
func Asset(path string) (string, string) {
	// identify the content type
	splits := strings.Split(path, ".")
	ext := splits[len(splits)-1]
	ctype, ok := ctypes[ext]
	if !ok {
		ctype = "application/octet-stream"
	}
	// extract data
	var str string
	var bytes []byte
	var err error
	if isBinary(ctype) {
		// encode binaries in base64
		bytes, err = box.MustBytes(path)
		if err == nil {
			str = base64.StdEncoding.EncodeToString(bytes)
		}
	} else {
		str, err = box.MustString(path)
	}
	if err != nil {
		return err.Error(), "404"
	}
	return str, ctype
}

// WebResponse returns a full response
// suitable for a Web Action
func WebResponse(path string) map[string]interface{} {
	// interpret as an asset
	body, ctype := Asset(path)
	// prepare the answer
	res := make(map[string]interface{})
	res["body"] = body
	if strings.Index(ctype, "/") != -1 {
		// asset found
		res["headers"] = map[string]string{
			"Content-Type": ctype,
		}
		res["statusCode"] = "200"
	} else {
		// asset not found
		res["statusCode"] = ctype
		res["headers"] = map[string]string{}
	}
	return res
}

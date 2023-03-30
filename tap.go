package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func Tap(x any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
	}
	y := map[string]any{
		"value": x,
		"type":  fmt.Sprintf("%T", x),
		"file":  file,
		"line":  line,
		"time":  time.Now().UnixNano(),
	}
	payload, err := json.Marshal(y)
	if err != nil {
		panic(err)
	}

	tap(payload)
}

func tap(x []byte) {
	resp, err := http.Post("http://127.0.0.1:5678/submit", "application/json", bytes.NewReader(x))
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}

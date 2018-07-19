package tripsgo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	ai "github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

//
var (
	Info  *log.Logger
	Debug *log.Logger
	Fatal *log.Logger
)

// Client - Pretty declaration of our telemetry client application insights
var (
	client = ai.NewTelemetryClient("91c2e8a3-5944-4ce4-bc6c-e5ee730cb607")
)

// InitLogging - Initialize logging for trips api
func InitLogging(
	infoHandle io.Writer,
	debugHandle io.Writer,
	fatalHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Fatal = log.New(fatalHandle,
		"FATAL: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// Logger - basic console logger that writes request info to stdout
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		Info.Println(fmt.Sprintf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))

		request := ai.NewRequestTelemetry(r.Method, r.RequestURI, time.Since(start), fmt.Sprintf("%s", r.Response.StatusCode))
		client.Track(request)

	})
}

func LogMessage(msg string) {
	Info.Println(msg)
	client.TrackEvent(msg)
}

func LogError(err error, msg string) {
	Info.Println(msg)
	client.TrackException(msg)
	Debug.Println(err.Error())

}

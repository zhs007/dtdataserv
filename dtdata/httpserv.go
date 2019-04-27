package dtdata

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/zhs007/dtdataserv/proto"
	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

func replyDTReport(w http.ResponseWriter, report *dtdatapb.DTReport) {
	jsonBytes, err := json.Marshal(report)
	if err != nil {
		jarvisbase.Warn("replyDTReport:Marshal", zap.Error(err))

		return
	}

	w.Write(jsonBytes)
}

// HTTPServer -
type HTTPServer struct {
	addr string
	serv *http.Server
	db   *dtDataDB
}

func (s *HTTPServer) onGetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	token := r.URL.Query().Get("token")
	cache, err := s.db.getCache(r.Context(), token)
	if err != nil {
		replyDTReport(w, newErrorDTReport(err))

		return
	}

	replyDTReport(w, cache)
}

// HTTPServer -
func newHTTPServer(addr string, db *dtDataDB) (*HTTPServer, error) {
	s := &HTTPServer{
		addr: addr,
		serv: nil,
		db:   db,
	}

	return s, nil
}

func (s *HTTPServer) start(ctx context.Context) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/getdata", func(w http.ResponseWriter, r *http.Request) {
		s.onGetData(w, r)
	})

	// fsh := http.FileServer(http.Dir("./www/static"))
	// mux.Handle("/", http.StripPrefix("/", fsh))

	server := &http.Server{
		Addr:         s.addr,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      mux,
	}

	s.serv = server

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *HTTPServer) stop() {
	if s.serv != nil {
		s.serv.Close()
	}

	return
}

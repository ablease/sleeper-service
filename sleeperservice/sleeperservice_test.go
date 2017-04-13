package sleeperservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/ablease/sleeper-service/sleeperservice"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SleeperService", func() {
	var (
		server *httptest.Server
	)

	BeforeEach(func() {
		router := mux.NewRouter()
		sleeperservice.AttachRoutes(router)
		server = httptest.NewServer(router)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("emitting communications", func() {
		var commsResp *http.Response
		var err error

		BeforeEach(func() {
			commsResp, err = http.Get(fmt.Sprintf("%s/comms", server.URL))
		})

		It("responsds with a good status code", func() {
			Expect(commsResp.StatusCode).To(Equal(http.StatusOK))
		})

	})
})

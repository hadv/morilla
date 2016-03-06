package api_test

import (
	. "github.com/hadv/morilla/api"

	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

/*
Convert JSON data into a map.
*/
func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

var _ = Describe("Server", func() {
	var server *httptest.Server
	var request *http.Request
	var usersUrl string

	BeforeEach(func() {
		server = httptest.NewServer(NewMovieServer())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GET /api/v1/movies", func() {

		// Set up a new GET request before every test
		// in this describe block.
		BeforeEach(func() {
			usersUrl = fmt.Sprintf("%s/api/v1/movies", server.URL)
			request, _ = http.NewRequest("GET", usersUrl, nil)
		})

		Context("when exist movies", func() {
			It("returns a status code of 200", func() {
				res, _ := http.DefaultClient.Do(request)
				Expect(res.StatusCode).To(Equal(200))
			})

			It("returns a status of [200 OK]", func() {
				res, _ := http.DefaultClient.Do(request)
				Expect(res.Status).To(Equal("200 OK"))
			})

			It("returns a body with a list of 2 movies", func() {
				res, _ := http.DefaultClient.Do(request)
				b, _ := ioutil.ReadAll(res.Body)
				moviesJSON := mapFromJSON(b)
				Expect(len(moviesJSON)).To(Equal(2))

				movieJSON_1, _ := moviesJSON["tt0082971"].(map[string]interface{})
				Expect(movieJSON_1["rating"]).To(Equal("8.6"))
				Expect(movieJSON_1["year"]).To(Equal("1981"))
				Expect(movieJSON_1["title"]).To(Equal("Indiana Jones: Raiders of the Lost Ark"))

				movieJSON_2, _ := moviesJSON["tt0076759"].(map[string]interface{})
				Expect(movieJSON_2["rating"]).To(Equal("8.7"))
				Expect(movieJSON_2["year"]).To(Equal("1977"))
				Expect(movieJSON_2["title"]).To(Equal("Star Wars: A New Hope"))
			})
		})
	})

	Describe("GET /api/v1/movie", func() {

		// Set up a new GET request before every test
		// in this describe block.
		BeforeEach(func() {
			usersUrl = fmt.Sprintf("%s/api/v1/movie/tt0076759", server.URL)
			request, _ = http.NewRequest("GET", usersUrl, nil)
		})

		Context("when exist movie", func() {
			It("returns a status code of 200", func() {
				res, _ := http.DefaultClient.Do(request)
				Expect(res.StatusCode).To(Equal(200))
			})

			It("returns a status of [200 OK]", func() {
				res, _ := http.DefaultClient.Do(request)
				Expect(res.Status).To(Equal("200 OK"))
			})

			It("returns a body with a movie", func() {
				res, _ := http.DefaultClient.Do(request)
				b, _ := ioutil.ReadAll(res.Body)
				movieJSON := mapFromJSON(b)

				Expect(movieJSON["rating"]).To(Equal("8.7"))
				Expect(movieJSON["year"]).To(Equal("1977"))
				Expect(movieJSON["title"]).To(Equal("Star Wars: A New Hope"))
			})
		})
	})

	Describe("GET /api/v1/movie", func() {

		// Set up a new GET request before every test
		// in this describe block.
		BeforeEach(func() {
			usersUrl = fmt.Sprintf("%s/api/v1/movie/aaaa1232423", server.URL)
			request, _ = http.NewRequest("GET", usersUrl, nil)
		})

		Context("Not found", func() {
			It("returns a status code of 404", func() {
				res, _ := http.DefaultClient.Do(request)
				Expect(res.StatusCode).To(Equal(404))
			})
		})
	})
})

package rss

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Мок для HTTP запросов
func MockHTTPServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/mock/rss", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<rss>
				<channel>
					<title>Title</title>
					<item>
						<title>News Title 1</title>
						<description>News Description 1</description>
						<pubDate>Mon, 02 Jan 2006 15:04:05 -0700</pubDate>
						<link>http://mock-link1.com</link>
					</item>
				</channel>
			</rss>
		`))
	})
	return httptest.NewServer(mux)
}

func TestParse(t *testing.T) {
	server := MockHTTPServer()
	defer server.Close()

	feed, err := Parse(server.URL + "/mock/rss")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(feed)) // Проверяем, что возвращается 1 пост
	assert.Equal(t, "News Title 1", feed[0].Title)
}

package http

import (
	"fmt"
	hackmiddleware "github.com/jensneuse/graphql-go-tools/hack/middleware"
	"github.com/jensneuse/graphql-go-tools/pkg/proxy"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestProxyHandler(t *testing.T) {
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != assetOutput {
			t.Fatalf("Expected:\n%s\ngot\n%s", assetOutput, body)
		}
	}))
	defer es.Close()

	schema := []byte(assetSchema)
	backendURL, err := url.Parse(es.URL)
	if err != nil {
		t.Fatal(err)
	}

	requestConfigProvider := proxy.NewStaticRequestConfigProvider(proxy.RequestConfig{
		Schema:     &schema,
		BackendURL: *backendURL,
	})

	ph := NewDefaultProxy(requestConfigProvider, &hackmiddleware.AssetUrlMiddleware{})
	ts := httptest.NewServer(ph)
	defer ts.Close()

	t.Run("Test proxy handler", func(t *testing.T) {
		_, err := http.Post(ts.URL, "application/graphql", strings.NewReader(assetInput))
		if err != nil {
			t.Error(err)
		}
	})
}

func BenchmarkProxyHandler(b *testing.B) {

	//go printMemUsage()

	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			b.Error(err)
		}
		if string(body) != assetOutput {
			b.Errorf("Expected %s, got %s", assetOutput, body)
		}
	}))
	defer es.Close()

	schema := []byte(assetSchema)
	backendURL, err := url.Parse(es.URL)
	if err != nil {
		b.Fatal(err)
	}

	requestConfigProvider := proxy.NewStaticRequestConfigProvider(proxy.RequestConfig{
		Schema:     &schema,
		BackendURL: *backendURL,
	})

	ph := NewDefaultProxy(requestConfigProvider, &hackmiddleware.AssetUrlMiddleware{})
	ts := httptest.NewServer(ph)
	defer ts.Close()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := http.Post(ts.URL, "application/graphql", strings.NewReader(assetInput))
		if err != nil {
			b.Error(err)
		}
	}
}

// nolint
func printMemUsage() {
	for {
		time.Sleep(time.Millisecond * time.Duration(1000))
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		// For info on each, see: https://golang.org/pkg/runtime/#MemStats
		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
		fmt.Printf("\tNumGC = %v\n", m.NumGC)
	}
}

// nolint
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

const assetSchema = `
schema {
    query: Query
}

type Query {
    assets(first: Int): [Asset]
}

type Asset implements Node {
    status: Status!
    updatedAt: DateTime!
    createdAt: DateTime!
    id: ID!
    handle: String!
    fileName: String!
    height: Float
    width: Float
    size: Float
    mimeType: String
    url: String!
}`

const assetInput = `{"query":"query testQueryWithoutHandle {assets(first: 1) { id fileName url(transformation: {image: {resize: {width: 100, height: 100}}})}}"}`

const assetOutput = `{"query":"query testQueryWithoutHandle {assets(first:1) {id fileName handle}}"}
`
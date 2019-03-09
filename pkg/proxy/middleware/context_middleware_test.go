package middleware

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jensneuse/graphql-go-tools/pkg/proxy/handler"
)

func TestContextMiddleware(t *testing.T) {
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set context that would usually be set in other application middleware
		userCtx := context.WithValue(r.Context(), "user", "jsmith@example.org")
		r = r.WithContext(userCtx)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		if string(body) != privateQuery {
			t.Errorf("Expected %s, got %s", privateQuery, body)
		}
	}))
	defer es.Close()

	ph := handler.NewProxyHandler([]byte(publicSchema), es.URL, &ContextMiddleware{})
	ts := httptest.NewServer(ph)
	defer ts.Close()

	t.Run("Test context middleware", func(t *testing.T) {
		r, err := http.NewRequest("POST", ts.URL, strings.NewReader(publicQuery))
		ctx := context.WithValue(context.Background(), "user", "jsmith@example.org")
		r = r.WithContext(ctx)
		if err != nil {
			t.Error(err)
		}
		client := http.DefaultClient
		_, err = client.Do(r)
		if err != nil {
			t.Error(err)
		}
	})
}


const publicSchema = `
schema {
	query: Query
}

type Query {
	documents: [Document] @context_insert(user: "user")
}

type Document implements Node {
	owner: String
	sensitiveInformation: String
}
`

// This schema is unused, left for reference
const privateSchema = `
schema {
	query: Query
}

type Query {
	documents(user: String!): [Document]
}

type Document implements Node {
	owner: String
	sensitiveInformation: String
}
`

const publicQuery = `
query myDocuments {
	documents {
		sensitiveInformation
	}
}
`

const privateQuery = `
query myDocuments {
	documents(user: "jsmith@example.org") {
		sensitiveInformation
	}
}
`
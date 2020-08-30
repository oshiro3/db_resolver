package find

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGitbucketResolver(t *testing.T) {
	t.Run("resolve http path to encoded body", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "http://githost:8080/api/v3/schema.sql",
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(200, "CREATE TABLE `test_table` (\n `id` int(20) NOT NULL, `name` varchar(20) COLLATE utf8_bin NOT NULL,\n `created_at` datetime DEFAULT NULL,\n `updated_at` datetime DEFAULT NULL,\n `dev` int(11) DEFAULT NULL\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;")
				return res, nil
			},
		)
		g := gitbucketResolver{path: "http://githost:8080/api/v3/schema.sql"}
		v, _ := g.Resolve()
		assert.Equal(t, "CREATE TABLE `test_table` (\n `id` int(20) NOT NULL, `name` varchar(20) COLLATE utf8_bin NOT NULL,\n `created_at` datetime DEFAULT NULL,\n `updated_at` datetime DEFAULT NULL,\n `dev` int(11) DEFAULT NULL\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;", v)
	})

	t.Run("fail get request when path is blank", func(t *testing.T) {
		g := gitbucketResolver{path: ""}
		_, e := g.Resolve()
		assert.EqualError(t, e, "Get \"\": unsupported protocol scheme \"\"")
	})

	t.Run("fail get request when path is wrong", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "http://githost:8080/api/v3/repos/root/test/raw/master/wrongPath",
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(404, "")
				return res, nil
			},
		)
		g := gitbucketResolver{path: "http://githost:8080/api/v3/repos/root/test/raw/master/wrongPath"}
		_, e := g.Resolve()
		assert.EqualError(t, e, "Invalid Path")
	})
}

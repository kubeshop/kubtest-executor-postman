package newman

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/stretchr/testify/assert"
)

// TestRun runs newman instance on top of example collection
// creates temporary server and check if call to the server was done from newman
func TestRun(t *testing.T) {
	// given
	runner := NewNewmanRunner()

	// and test server for getting newman responses
	requestCompleted := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCompleted = true
	}))

	defer ts.Close()

	parts := strings.Split(ts.URL, ":")
	port := parts[2]

	execution := testkube.Execution{
		Content: testkube.NewStringTestContent(fmt.Sprintf(exampleCollection, port, port)),
	}

	// when
	result, err := runner.Run(execution)

	// then
	assert.NoError(t, err)
	assert.Empty(t, result.ErrorMessage)
	assert.Contains(t, result.Output, "Successful GET request")
	assert.Equal(t, requestCompleted, true)
}

const exampleCollection = `
{
	"info": {
		"_postman_id": "3d9a6be2-bd3e-4cf7-89ca-354103aab4a7",
		"name": "testkube",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Test",
			"event": [
				{
					"listen": "test",
					"script": {
						"exec": [
							"    pm.test(\"Successful GET request\", function () {",
							"        pm.expect(pm.response.code).to.be.oneOf([200, 201, 202]);",
							"    });"
						],
						"type": "text/javascript"
					}
				}
			],
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://127.0.0.1:%s",
					"protocol": "http",
					"host": [
						"127",
						"0",
						"0",
						"1"
					],
					"port": "%s"
	
				},
				"host": ["localhost"]
			},
			"response": []
		}
	]
}
`

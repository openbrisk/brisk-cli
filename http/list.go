package http

import (
	"net/http"
	"strings"

	"github.com/openbrisk/brisk-cli/requests"
)

// ExecuteListFunctions lists the deployed functions.
func ExecuteListFunctions(clusterURL string) ([]requests.Function, error) {
	var results []requests.Function

	clusterURL = strings.TrimRight(clusterURL, "/")

	request, err := http.NewRequest(http.MethodGet, clusterURL+"/functions", nil)
	SetAuth(request, clusterURL)

	return results, nil
}

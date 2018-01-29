package http

// CreateHTTPClient creates a http client with some sane defaults.
/*func CreateHTTPClient(timeout *time.Duration) http.Client {

	if timeout != nil {
		return http.Client{
			Timeout: *timeout,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout: *timeout,
				}).DialContext,
				IdleConnTimeout:       120 * time.Millisecond,
				ExpectContinueTimeout: 1500 * time.Millisecond,
			},
		}
	}

	// THisd should be used for brisk invoke.
	return http.Client{}
}*/

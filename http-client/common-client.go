//All the calls make externally are defined in this package.
package httpclient

import (
	"context"
	"errors"
	"net/http"
	"time"
)

//makes a call to Http Client. Necessary options like timeouts have been added. 
//This method only supports get Http Call as Posts call are not part of requirement
func MakeHttpCall(method string, body *interface{}, url string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10)) //create context with timeout
	defer cancel() //free up all the context resources
	if method == http.MethodGet { 
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil) // Form a http request
		if err != nil {
			return nil, err
		}
		c := &http.Client{}
		return c.Do(req) 

	}
	return nil, errors.New("method to be called is not http.MethodGet") // If method passed is not GET return error

	//TODO :: handle post http but currently it is not required.
}

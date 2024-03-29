/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package dedicatedshortcode_pkg


import(
	"fmt"
	"github.com/apimatic/unirest-go"
	"ytelapi_lib/apihelper_pkg"
	"ytelapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type DEDICATEDSHORTCODE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Update a dedicated shortcode.
 * @param    string         shortcode          parameter: Required
 * @param    *string        friendlyName       parameter: Optional
 * @param    *string        callbackMethod     parameter: Optional
 * @param    *string        callbackUrl        parameter: Optional
 * @param    *string        fallbackMethod     parameter: Optional
 * @param    *string        fallbackUrl        parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) UpdateShortcode (
            shortcode string,
            friendlyName *string,
            callbackMethod *string,
            callbackUrl *string,
            fallbackMethod *string,
            fallbackUrl *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/updateshortcode.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : shortcode,
        "FriendlyName" : friendlyName,
        "CallbackMethod" : callbackMethod,
        "CallbackUrl" : callbackUrl,
        "FallbackMethod" : fallbackMethod,
        "FallbackUrl" : fallbackUrl,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of Short Code assignment associated with your Ytel account.
 * @param    *string        shortcode     parameter: Optional
 * @param    *string        page          parameter: Optional
 * @param    *string        pagesize      parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) CreateListShortcodes (
            shortcode *string,
            page *string,
            pagesize *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/listshortcode.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : shortcode,
        "page" : page,
        "pagesize" : pagesize,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrive a list of inbound Sms Short Code messages associated with your Ytel account.
 * @param    *int64         page            parameter: Optional
 * @param    *int64         pagesize        parameter: Optional
 * @param    *string        from            parameter: Optional
 * @param    *string        shortcode       parameter: Optional
 * @param    *string        datecreated     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) CreateListInboundSMS (
            page *int64,
            pagesize *int64,
            from *string,
            shortcode *string,
            datecreated *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/getinboundsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "page" : page,
        "pagesize" : pagesize,
        "From" : from,
        "Shortcode" : shortcode,
        "Datecreated" : datecreated,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a single Short Code object by its shortcode assignment.
 * @param    string        shortcode     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) CreateViewSMS (
            shortcode string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/viewshortcode.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : shortcode,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of Short Code messages.
 * @param    *string        shortcode     parameter: Optional
 * @param    *string        to            parameter: Optional
 * @param    *string        dateSent      parameter: Optional
 * @param    *int64         page          parameter: Optional
 * @param    *int64         pageSize      parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) CreateListSMS (
            shortcode *string,
            to *string,
            dateSent *string,
            page *int64,
            pageSize *int64) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/listsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Shortcode" : shortcode,
        "To" : to,
        "DateSent" : dateSent,
        "Page" : page,
        "PageSize" : pageSize,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Send Dedicated Shortcode
 * @param    int64          shortcode                 parameter: Required
 * @param    float64        to                        parameter: Required
 * @param    string         body                      parameter: Required
 * @param    *string        method                    parameter: Optional
 * @param    *string        messagestatuscallback     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) CreateSendSMS (
            shortcode int64,
            to float64,
            body string,
            method *string,
            messagestatuscallback *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/dedicatedshortcode/sendsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "shortcode" : shortcode,
        "to" : to,
        "body" : body,
        "method" : method,
        "messagestatuscallback" : messagestatuscallback,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * View a single Sms Short Code message.
 * @param    string        messageSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *DEDICATEDSHORTCODE_IMPL) CreateViewSMS (
            messageSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/shortcode/viewsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "MessageSid" : messageSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}


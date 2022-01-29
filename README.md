# ms-url-shortner
Micro-service used to short the url.

# Technical Specification
    1. If length of URL is 20 or less then 20 then no need to shorten the URL. For example "https://www.test.com", the length of
       this url is 20 so we don't need to short this URL. Same URL will be return in JSON response body.
    
    2. If the URL is empty or doesn't contain "https://" OR "https://" as prefix, then the URL will be treated as invalid URL.
       For example "infracloud.io" will be treated as invalid, because it doesn't contain "https://" OR "https://" as prefix.

    3. If URL is valid then a new shorted url will be generated with a 8 digit unique id.
       For example "https://www.infracloud.io" will be converted to "https://shorturl.com/ebs123le. ebs123le is a 8 digit unique id.

    4. Once the shorted URL is generated successfully, then it will be save into memory(map), so every time when request is 
    made then it will first check in map, and if the URL exists then it will directly return the same shorted URL from map.

# Error Codes for JSON Response- 
    -----------------------------------------------------------------------------------------
    | Error Codes               |  Description                                              |
    -----------------------------------------------------------------------------------------
    | 400(Bad Request)          |  If JSON request is not able to bind with object          |
    -----------------------------------------------------------------------------------------
    | 422(Unprocessable entity) |  If data inside JSON request is empty or is not valid     |
    -----------------------------------------------------------------------------------------
    | 200(status ok)            |  URL is successfully shorted                              |
    -----------------------------------------------------------------------------------------

# go run main.go -> this command will start the application locally

# Application Running Port - 4242

# Base URL - http://localhost:4242/ms-url-shortner

# EndPoints - 
    ---------------------------------------------------------------------------------------------------------------------------------
    | EndPoints    | Request Type | Description                                 | URL                                               |
    ---------------------------------------------------------------------------------------------------------------------------------
    | /ping        |    Get       | test whether application is running or not  | http://localhost:4242/ms-url-shortner/ping        |
    ---------------------------------------------------------------------------------------------------------------------------------
    | /getshorturl |    POST      | get the shorted URL                         | http://localhost:4242/ms-url-shortner/getshorturl |
    ---------------------------------------------------------------------------------------------------------------------------------

# Json Request Object for endpoint -> /getshorturl 
    {
        url: string
    }
    Example - 
    {
        "url":"https://infracloud.io"
    }
    
# Json Response object for endpoint - > /getshorturl 
        {
            data  interface{} 
            error {
                code int
                message string
            }      
        }

    Example
    1. If there is no error it means response is successful
        {
            "data": {
                "url": "https://shorturl.com/ecfd35c4"
            },
            "error": null
        }
    2. If there is any error
            {
                "data": null,
                "error": {
                    "code": 422,
                    "message": "invalid url"
                }
            }

# Docker image - https://hub.docker.com/r/shubhambansal96/msurlshortner
    docker image for this app is available on docker hub, you can run below mentioned command to run image on your machine. I have
    created this image as public for as of now, so no credentials are required to pull this image.

    Image Name - shubhambansal96/msurlshortner

    1. To pull and run this image -> docker run -td -p 4246:4242 shubhambansal96/msurlshortner
                
    NOTE -> Container will be running on port 4242, make sure you map host port to container port on 4242
        example -> docker run -td -p 4246:4242 shubhambansal96/msurlshortner, 4246 is host port, you can give any port number in place of 4246.
                                            
# END
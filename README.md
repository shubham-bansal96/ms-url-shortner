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
    400(Bad Request) - If JSON request is not able to bind with object
    422(Unprocessable entity) - If data inside JSON request is empty or is not valid
    200(status ok) - URL is successfully shorted

# go run main.go -> this command will start the application locally

# Application Running Port - 4242

# URL - http://localhost:4242/ms-url-shortner

# EndPoints - 
    /ping - GET -> test whether application is running or not
    /getshorturl - POST -> get the shorted URL

# Json Request Object for endpoint -> /getshorturl 
    {
        url: string
    }
    example - 
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

    example
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
--------------------------------------------------------------------------------------------------------------
# Docker image - https://hub.docker.com/r/shubhambansal96/msurlshortner
    docker image for this app is available on docker hub(public image), you can run below mentioned command to run image on your machine.

    1. To get the image -> docker pull shubhambansal96/msurlshortner
                
    2. Container will be running on port 4242, make sure you map host port to container port on 4242
        example -> docker run -td -p 4246:4242 shubhambansal96/msurlshortner
                                            
# END
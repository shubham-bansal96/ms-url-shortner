Application Running Port - 4242

Error Codes - 
    400(Bad Request) - If JSON request is not able to bind with object
    422(Unprocessable entity) - If data inside JSON request is empty or is not valid
    200(status ok) - URL is successfully shorten

EndPoints - 
    ms-url-shortner/ping - GET -> test whether application is running or not
    ms-url-shortner/getshorturl - POST -> get the short URL
# Hike Safe (Under Construction)
##### An API service to assists hikers in setting their trails and getting important updates about the trail from other hikers.

## End points

#### Authentication
To get the API key, make a POST request containing username and password to the following endpont:

```
POST https://localhost:1323/auth
```
The post body would look like this

```JSON
{
    "username":"exampleuser",
    "password":"examplepassword"
}

```

If the username is not already registered the API will response with a key in the following way

```JSON
{
    "key":"randomcharactersequencekey"
}
```

You can also check if they key you have is authentic by using the following endpoint

```
GET https://localhost:1323/check
?key=API_KEY
```

If the key is authentic you would get the following response

```JSON
{
    "message":"Key is authenticated"
}
```

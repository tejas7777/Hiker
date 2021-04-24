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

#### Getting trails
You can get list of all availaible trails to subscribe to by using the following endpoint:

```
GET https://localhost:1323/trails
?key=API_KEY
```

If the request is successfull you'll get the following response:

```JSON
{
    "trails":{
        "60829121098b2fc8eee3ab13": "Salt Marsh Nature Trail",
        "60829121098b2fc8eee3ab14": "Lullwater",
        "60829121098b2fc8eee3ab15": "Midwood",
        "60829121098b2fc8eee3ab16": "Peninsula",
        "60829121098b2fc8eee3ab17": "Waterfall",
        "60829121098b2fc8eee3ab18": "Alley Pond Trails",
        "60829121098b2fc8eee3ab19": "Blue Trail",
        "60829121098b2fc8eee3ab1a": "Orange Trail",
        "60829121098b2fc8eee3ab1b": "Yellow Trail",
        "60829121098b2fc8eee3ab1c": "South Preserve Trail",
        "60829121098b2fc8eee3ab1d": "Greenbelt Blue Trail (Southern Trailhead)",
        "60829121098b2fc8eee3ab1e": "Greenbelt Nature Center Trail"
        \.
        \.
        \.
    }
}
```

# Forecast API
A simple api which uses [National Weather Service API Web Service](https://www.weather.gov/documentation/services-web-api). Simply, the program gets latitude and longitude, then, returns the *short forecast for that area for Today (“Partly Cloudy” etc)* and *a characterization of whether the temperature is “hot”, “cold”, or “moderate”*. I'll explain more about it in the following.

## How It Works
Fortunately, ([National Weather Service API Web Service](https://www.weather.gov/documentation/services-web-api)) supports latitude and longitude to give you the forecast of the area. You may want to read [the document](https://www.weather.gov/documentation/services-web-api#/default/point). Simply, the```https://api.weather.gov/points/{point}``` endpoint, which is a GET endpoint, returns metadata about a given latitude/longitude point. Here you can see how to try it:

``` bash
curl -X GET "https://api.weather.gov/points/your_latitude,your_longitude" -H "accept: application/ld+json"  

```
In response, it returns an endpoint("forecast": "string") that contains the forecast or the area. Thus, we'll need to send another request to get the forecast of the area.

``` bash
curl -X GET "https://api.weather.gov/gridpoints/HNX/change,change/forecast" -H "accept: application/ld+json"
```

On the above, you can see how my program works to return the forecast.


# How to Run
1. Clone the program:
``` bash
git clone https://github.com/hilton-james/forecast

```
2. Run:
``` bash
make run
```
 After running the program with ```make run``` it gives you and endpoint on *http://localhost:5001/forecase* which receives *lat* (latitude) and *long* (longitude). You can send request with ```curl```

``` bash
curl -v 'localhost:5001/forecast?lat=change&long=change'
```
In response, you should receive  *lat* (latitude) and *long* (longitude) if they are right.

``` json
{"forecast":"Partly Cloudy then Areas Of Fog","temperature":"moderate"}
```
# API geo-sensor simulator

This is a simple API to simulate geo-sensor data.

## Installation

*Require GO install on your computer. Install it <a href="https://go.dev/">here</a>*


```bash
$ git clone
```
1- Create your database<br>
2- Create a .env file from the .env.example file. Fill in the database credentials and the port.

## Usage
```bash
$ go build
$ ./api_sensor
```



## Endpoints

- POST <a>/api/auth</a> to get an authentication token
- GET <a>/api/air</a> to get the 60 latest records from the air-sensor
- GET <a>/api/air/latest</a> to get the latest record from the air-sensor




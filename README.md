# Unit Converter App

A simple Go web service that converts kilometers to miles.

- GET /health

- POST /convert


Input JSON :

{
  "kilometers": 10
}


Output JSON : 

{
  "kilometers": 10,
  "miles": 6.21371
}


## Run locally

go run main.go


## Test 

go test ./...



## Once running open new terminal and run : 

curl -X POST http://localhost:8080/convert \
  -H "Content-Type: application/json" \
  -d '{"kilometers":10}'



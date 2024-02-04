# list humans
curl http://localhost:8080/human -X GET -H "Content-Type: application/json" 

# crate human data.
curl http://localhost:8080/human -X POST -H "Content-Type: application/json" -d '{"id": 4,"name": "yuki","birthday": "1995-11-11","address": {"country": "Japan","state": "Tokyo"}}'

# update human data.
curl http://localhost:8080/human/4 -X PUT -H "Content-Type: application/json" -d '{"id":4,"name": "yuki","birthday":"1995-11-11","address":{"country":"Japan","state":"Okinawa"}}'

# get human data.
curl http://localhost:8080/human/4 -X GET -H "Content-Type: application/json" 

# delete human data.
curl http://localhost:8080/human/4 -X DELETE -H "Content-Type: application/json"







# Kitties

## Queries


Create a Kitty

    curl -X POST -d '  {
      "name": "Michel",
      "breed": "British",
      "birthDate": "2016-07-05"
    }' "http://localhost:8080/api/v1/kitties"


Show all kitties


    curl http://localhost:8080/api/v1/kitties


Kill a kitty


    curl -X DELETE "http://localhost:8080/api/v1/kitties/1"

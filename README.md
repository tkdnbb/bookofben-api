# bookofben-api
The API service for the book of Jachanan Ben Kathryn.

## docker commands
docker run -d \
  --name bookofben-mongo \
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=example \
  mongo:latest

docker exec -it bookofben-mongo mongosh -u root -p example

## MongoDB shell commands
use bible_api
db.books.find()

## .env
Please create a .env file.
```
MONGO_CONNECTION=mongodb+srv://example
```

## API testing
curl -XGET http://localhost:8080/john%203:16
curl -XGET http://localhost:8080/the%20book%20of%20jachanan%20ben%20kathryn%201
curl -XGET "http://localhost:8080/api/search?q=the%20LORD%20would"
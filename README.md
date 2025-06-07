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

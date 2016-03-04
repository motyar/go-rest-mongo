# Go JSON REST API with mongoDb database

Demo https://rest-golangapi.rhcloud.com/pandas/

# Using client.go

### List all the Pandas
``
go run client.go -url=https://rest-golangapi.rhcloud.com/pandas -method=get
``
### Create new Panda
``
go run client.go -url=https://rest-golangapi.rhcloud.com/pandas/ -method=post -data='{"Name":"Name of Panda"}'
``
### Delete existing Panda
``
go run client.go -url=https://rest-golangapi.rhcloud.com/pandas/{Id} -method=delete
``
### Update existing Panda

`` go run client.go -url=https://rest-golangapi.rhcloud.com/pandas/{Id} -method=put -data='{"Name":"New name"}'
``

# Nothing but cute Pandas

GET /pandas/ to get all the pandas.
    or /pandas/:id to get on with id

POST /pandas/ to add new panda {"name":"Name of new panda"}

DELETE /pandas/:id to remove that one panda.

PUT /pandas/ to update details {"name":"Name of new panda"}


users-service
=============

A microservice for experimenting with Docker. To run the image from Dockerhub, run `$ docker run -d -p 8080:8080 jmoorhead1/users-service:latest`.

API
---

`GET:/users/all` - returns an array of all users
`GET:/users/:username` - returns the user specified by `:username`, or 404 if not found
`POST:/users/add` - add a new user specified in the request body, returns 409 if username already exists
`PUT:/users/:username/update` - update the user specified by `:username`, returns 404 if not found

Questions
---------

Please reach out to jmoorhead@performline.com with any questions.

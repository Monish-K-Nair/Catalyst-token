Instructions to run

Install docker 

docker-compose build
docker-compose up 

FOR JWT TOKEN 
1. Create a row inside the table admins in the database catalyst { username catalyst-user}
2. /admin/login will generate the token 

FOR ALL REQUESTS EXCEPT /api/vi/invite-token/validate JWT TOKEN IS REQUIRED 

1. ADMIN LOGIN FOR TOKEN
Sample Request
-------------
curl --location --request POST 'http://localhost:8080/admin/login/' --header 'Content-Type: application/json' --data-raw '{"Email": "sample@sample.sample", "Password":"1234"}'

Sample Response 
----------------
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemF0aW9uIjp0cnVlLCJlbWFpbCI6InNhbXBsZUBzYW1wbGUuc2FtcGxlIiwiZXhwIjoxNjYxMjE1MzU3LCJpZCI6IjEifQ.p3JoBGprrCDRhMp56aiJ8ddQlSqWTz4ARGqeUt3G1m0"

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


2. GENERATING INVITE TOKEN WITHOUT JWT KEY 

curl --location --request POST 'http://localhost:8080/api/v1/invite-token' \
--data-raw ''

{
    "status": "Forbidden",
    "code": 403,
    "message": "Auth required"
}

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

3. GENERATING INVITE TOKEN WITH JWT TOKEN 

curl --location --request POST 'http://localhost:8080/api/v1/invite-token' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemF0aW9uIjp0cnVlLCJlbWFpbCI6InNhbXBsZUBzYW1wbGUuc2FtcGxlIiwiZXhwIjoxNjYxMjE1MzU3LCJpZCI6IjEifQ.p3JoBGprrCDRhMp56aiJ8ddQlSqWTz4ARGqeUt3G1m0'

4. LIST ALL TOKENS WITH JWT TOKEN

curl --location --request GET 'http://localhost:8080/api/v1/invite-token' \
--header 'Authorization: <Bearer Token>'

{
    "data": [
        {
            "ID": "6d8f9362-7c78-400a-a6a6-77c524948ab1",
            "Token": "a6e798d3b80c",
            "Name": "",
            "Active": false,
            "CreatedAt": "2022-08-15T23:43:46.737887Z",
            "UpdatedAt": "2022-08-15T23:49:52.791973Z"
        },
        {
            "ID": "ca288263-801b-4a49-bdec-6740ea371c94",
            "Token": "bddaba3b46a1",
            "Name": "",
            "Active": false,
            "CreatedAt": "2022-08-15T23:52:10.497553Z",
            "UpdatedAt": "2022-08-16T00:19:52.388403Z"
        },
        {
            "ID": "dda7fbfc-b1db-4fdc-937b-92f6a8f68df7",
            "Token": "92edf35c044d",
            "Name": "",
            "Active": true,
            "CreatedAt": "2022-08-16T00:51:31.178632Z",
            "UpdatedAt": "2022-08-16T00:51:31.178639Z"
        }
    ],
    "message": "List of All Tokens"
}


5. VALIDATE TOKEN FOR FALSE INPUT

curl --location --request POST 'http://localhost:8080/api/v1/invite-token/validate' \
--header 'Content-Type: application/json' --data-raw '{"Token": "1234"}

{
    "data": "1234",
    "message": "Token Invalid"
}

6. VALIDATE TOKEN FOR valid INPUT

curl --location --request POST 'http://localhost:8080/api/v1/invite-token/validate' \
--header 'Content-Type: application/json' --data-raw '{"Token": "92edf35c044d"}'

{
    "data": "92edf35c044d",
    "message": "Token Valid"
}


7. REVOKE TOKEN

curl --location --request PUT 'http://localhost:8080/api/v1/invite-token/' \
--header 'Authorization: Bearer token'
--header 'Content-Type: application/json'
--data-raw '{"Token": "92edf35c044d"}'

{
    "data": "92edf35c044d",
    "message": "Revoked the token"
}

8. DELETE TOKEN
curl --location --request DELETE 'http://localhost:8080/api/v1/invite-token/' \
--header 'Authorization: Bearer Token 
--header 'Content-Type: application/json' \
--data-raw '{"Token": "92edf35c044d"}'

{
    "data": "92edf35c044d",
    "message": "Deleted Token"
}
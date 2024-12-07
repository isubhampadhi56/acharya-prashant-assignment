### How to start Application - 
The main file is under cmd directory. Just run it using go run cmd/main.go
It uses sqlite DB to store users data.
All Configs has been already present in pkg/config 

### 1. User Signup
The signup request should have a JSON object in it's body which should have email and password attributes.
Note that email should be in valid format and password should be atleast 8 character in length. Otherwise input validation will fail and user will not be created.


curl --location 'http://localhost:8080/api/v1/auth/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"user1@mail.com",
    "password": "123456789"
}'

### 2. User Login
The login request should have a JSON object in it's body which should have email and password attributes.
After Successful User Login it should return accessToken and refreshToken in response in JSON format.

curl --location 'http://localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"user1@mail.com",
    "password": "123456789"
}'

### 3. Authorization Token and Protected Routes
There are 2 protected Routes which requires a valid AccessToken to access.
For accessing protected routes Authorization should be present on Request Header. 

 - GET /api/v1/me which checks if user's access token is still valid or not.
 - GET /api/v1/user which returns the details of LoggedIn user
 - PATCH /api/v1/user/deactivate

curl --location 'http://localhost:8080/api/v1/me' \
--header 'Authorization: <Replace with your accessToken after login>'

curl --location 'http://localhost:8080/api/v1/user' \
--header 'Authorization: <Replace with your accessToken after login>'

### 4. Revokation of Token
PATCH /api/v1/deactivate API will deactivate the loggedIn User
- After expiration of accesstoken user can't renew it using refreshtoken.
- User login also won't generate accesstoken and refreshtoken after disabling user account.
- The existing Authorization token will expire within 5 mins.

curl --location --request PATCH 'http://localhost:8080/api/v1/deactivate' \
--header 'Authorization: <Replace with your accessToken after login>'

The other approach is to store current Authorization token in some centralized cache and blacklist this token while checking for authorization in middleware.

### 5. Refresh Access Token

As per current config AccessToken expires in 5 minutes and refresh Token Expires in 72 Hours (defined in config). 
To renew AccessToken a GET request has to be sent to api/v1/auth/token route.
If the RefreshToken Header on Request is valid and user is active it will generate a valid access token and send it in Response Body.
If the user's account has been deactivated or deleted or updated after Refresh Token got generated then the Refresh Token request will fail and prompt user to relogin.

curl --location 'http://localhost:8080/api/v1/auth/token' \
--header 'RefreshToken: <Replace with your refreshToken after login>'


### Time Tracker
This is a small full stack web app, that can help a freelancer track their time.

### Approach

I used Golang and MongoDb on the backend and React with Javascript on the frontend. 

This test is not complete - please see the To-Do section below for my priorities in how I would continue to work on this 
project if I gave it more time. 

#### State management
Initially considered handling state all on backend i.e. sending a POST for starting and stopping as well as saving a 
session. The benefit would have been that you have persistence cross browser window - so if someone closed window it 
would not lose session. However for simplicity I went with client side state management and only added an endpoint for 
saving a session. I avoided Redux as the standard React state management was adequate for the requirements and 
complexity of the app at this stage. 

#### Users
I built a concept of a user login using a UURI. Users can link browser sessions by saving their initial UUID and 
re-entering it if they wish to associate a new session with old ones. 

The downside of this is that it is not secure and users have to store a long uri somewhere but it is a simple 
first step to user persistence. 

The implementation does not yet handle edge cases such as users entering an id that is not a valid UUID whcih could 
cause conflicts with other users. 

#### Styling
Styling is very basic as I decided to focus on the backend and delivering the required features instead.

### AC's
* As a user, I want to be able to start a time tracking session

* As a user, I want to be able to stop a time tracking session

* As a user, I want to be able to name my time tracking session

* As a user, I want to be able to save my time tracking session when I am done with it

`POST /v1/users/:userId/session`
{ "sessionName": Optional[""], "start": utc-timestamp, "end": utc-timestamp, "duration": unix-timestamp }

* As a user, I want an overview of my sessions for the day, week and month

`GET /v1/users/:userId/sessions`

* As a user, I want to be able to close my browser and shut down my computer and still have my sessions visible to me when I power it up again.

### To Run:


#### Start a local db:
```shell script
docker run -p 27017:27017 mongo:4.4
```

#### Run the frontend:
```shell script
cd web
yarn install
yarn start
```
Navigate to http://localhost:3000

#### Run the server:
```shell script
go build
go run .
```

### To Do

1. Unit tests!
2. Integration tests of the server and Cypress tests for frontend
3. Improve styling
4. Improve User Id edge case handling

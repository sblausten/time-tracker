### Pento tech challenge
The challenge is to build a small full stack web app, that can help a freelancer track their time.

### Approach

#### State management
Initially considered handling state all on backend i.e. sending a POST for starting and stopping as well as saving a 
session. The benefit would have been that you have persistence cross browser window - so if someone closed window it 
would not lose session. However this would only work after persistent user management was added (see below), so for 
simplicity I went with client side state management and only added an endpoint for saving a session. I avoided Redux 
as the standard React state management was adequate for the requirements and complexity of the app at this stage. 

#### Users
I built a concept of a user id using a UURI. Users can link browser sessions by saving their initial UUID and 
re-entering it if they wish to associate a new session with old ones. 

The downside of this is that it is not completely secure and users have to store a long uri somewhere but it is a simple 
first step to user persistence. 

The implementation does not yet handle edge cases such as users entering an id that is not a valid UUID whcih could 
cause conflicts with other users. 

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


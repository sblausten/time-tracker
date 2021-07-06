### Pento tech challenge
The challenge is to build a small full stack web app, that can help a freelancer track their time.

### Approach
Initially considered handling state all on backend, so sending a POST for starting and stopping as well as saving a 
session. Benefit would be that you have persistance cross browser window - so if someone closed window it would not lose session.

For simplicity and first MVP went with client side state management and only added endpoint for saving. 

To avoid implementing user management and login, I opted for a uri based user persistance mechanism where users must save 
their initial uri and re-enter it if they wish to associate a new session with old ones. The downside of this is 
that it is not completely secure and users have to store a long uri somewhere. 

### AC's
It should satisfy these user stories:

* As a user, I want to be able to start a time tracking session

* As a user, I want to be able to stop a time tracking session

* As a user, I want to be able to name my time tracking session

* As a user, I want to be able to save my time tracking session when I am done with it

`POST /v1/users/:userId/session`
{ "sessionName": Option[""], "start": timestamp, "end": timestamp }

* As a user, I want an overview of my sessions for the day, week and month

`GET /v1/users/:userId/sessions?duration=<day|week\month>`

* As a user, I want to be able to close my browser and shut down my computer and still have my sessions visible to me when I power it up again.

#### Timing
Don't spend more than a days work on this challenge. We're not looking for perfection, rather try to show us something special and have reasons for your decisions.
Get back to us when you have a timeline for when you are done.

#### Notes
Please focus on code quality and showcasing your skills regarding the role you are applying to.

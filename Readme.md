# PetsAlone Technical Assesment

Picture the scene…you’ve tuned into your favourite show on Netflix ready to relax for the evening…and the phone rings…

It transpires that your friend, Bill, has lost his beloved pet – a rare breed Norwegian Forest Cat. In fact it’s been about 48 hours and Bill is rather distressed.

Thankfully Bill hasn’t asked you to go out searching on foot. However, knowing that you’re a keen technologist, he asks if you could help him out in a different way. He has had an idea of creating a web-based system to help owners raise awareness of their missing pets.

Bill has come up with a name (which he’s very proud of) ‘PetsAlone’, and he has also given you a head-start and made it available on GitHub, which you can clone and then modify to meet his requirements… You begrudgingly agree to help, knowing that Bill isn’t an accomplished developer, but he’s in a desperate situation and you want to help out. You manage to get some high-level requirements from Bill before he hangs up...


### User stories

As an anonymous user  
I want to see all missing pets  
So that I can help to reunite a missing pet with its owner.  

As an authenticated user  
I want to create a new missing pet on the system  
So that I can raise awareness of a missing pet.

### Assessment Acceptance criteria

1. The home page of the site will show all the missing pets, with the most recently added pets listed first
2. A user will be able to filter the pets that have gone missing by animal type (dog, cat, ferret etc)
3. A new page, only accessible by an authenticated user, which will be able to create a new missing pet. Bill has created a model to use in pkg/service/pets.go
4. A developer picking up the project after you should see some evidence of you testing the code, even if it only covers a few key scenarios
5. You should also include any recommendations for changes or problems found in the original application, to be given to a developer picking up the project

### To help save you some time...

-   The project found in this GitHub repo contains a service with some existing methods to get you started. It interacts with an in-memory database which can be used in this exercise in place of a ‘real’ database. 

-   This project also contains a mocked identity provider which can be used to simulate logging in. For an authenticated user you can log in with the following credentials:  
Username: elmyraduff  
Password: MontanaMax!!

-	You’ll find ‘boilerplate’ projects for React and Angular in the repo as well (look inside the 'ui' folder'). Feel free to use one of these to get you started on the frontend, but if you’d prefer to start from scratch (or use a different technology or framework) that’s also fine. Remember - Bill is not an experienced developer, so don't be shy about changing or removing anything that you don't need.

-	Bill needs this app live ASAP, so aim to spend only a couple of hours on this. 

### Running the app
Ensure you have golang > 1.13 and npm installed on your machine. A Makefile is imncluded with some basic targets.

### ======= Added by sajjan ========

## API documentation
[API](docs/openapi.yaml)

## Test Execution
  
```
make test
```
## GUI testing locally

```
    make build-react ; make run-react and access gui as http://localhost:8080
```
## Observations/Suggestions
- pet_type changed in db as string from integer for ease of use
- Logger library added - zap (go.uber.org/zap)
- API endpoint implementations are in api/service.go
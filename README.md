---
# weather-mini-project

a very simple project to play with go fiber package, we have three main parts:
- cmd which starts the program
- handler where it registers a route bound to specific Http method
- the real login and weight of the program is in sevice layer

- handler gets the HTTP input, calls service then converts output into HTTP response
- service deals with business logic, communicates with outside of the program, then it gets the info and hands it to handler 

Summery:
```
browser → main → handler → service → handler → browser

```
## further aims to improve is
- [ ] make it more privacy friendly; hiding the api key in `.env` file
- [ ] create a surface including type and response features

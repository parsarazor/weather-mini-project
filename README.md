---
# weather-mini-project

a very simple project to play with go fiber package, we have three main parts:
1. cmd which starts the program
2. handler where it registers a route bound to specific Http method
3. the real weight of the program is on the sevice layer

- handler gets the HTTP input, calls service then converts output into HTTP response
- service deals with business logic, communicates with outside of the program, then it gets the info and hands it to handler 

Summery:
```

browser → main → handler → service → handler → browser

```
## further aims to improve is
- [ ] make it more privacy friendly; hiding the api key in `.env` file
- [ ] create a surface including type and response features

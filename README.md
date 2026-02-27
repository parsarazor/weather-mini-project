# weather-mini-project
---
a very simple project to play with go fiber package, we have three main parts:
- main which starts the program
- handler where it registes a route bound to specific Http method
- the real login and weight of the program is in sevice

handler deals with 
service deals with business logic, call with outside
then it gets the info and hands it to handler 

## further aims to improve is
- [ ] make it more privacy friendly; hiding the api key in `.env` file
- [ ] create a surface including type and response features

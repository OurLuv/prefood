## Project Description
I have developed a REST API monolith using the Golang. With this API, you can manage data in your food ordering application. 

My Golang development serves as the foundation for next projects in food ordering. While it is just a pet project, the API ensures reliability and efficiency required for the smooth operation of an application.

## Swagger Documentation
Check out Swagger documentation [here](https://ourluv.github.io/#/) <-


## Used Technologies 
- Gorilla Mux - An URL router and dispatcher for creating flexible and efficient routes in your Golang application.
- Pgxpool - A PostgreSQL connection pool library for Golang, allowing management and utilization of database connections.
- Docker with Postgres Image - Docker provides containerization for your application and using a Postgres image allows for setup and management of your database environment.
- Golang JWT - A library for generating JWT (JSON Web Tokens) and handling authentication in the application.
- Gomock - A mocking framework for Golang for generating mock objects for unit testing, enabling easier and more effective testing of your code.


## Architecture
![](https://raw.githubusercontent.com/OurLuv/prefood/master/static/architecture.drawio.png?token=GHSAT0AAAAAACNNM3DHTLLAPPVJP5UIHDNSZNU2ILA)


## To launch the application
```bash
make build && make run
```

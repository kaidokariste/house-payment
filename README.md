# API structure and endpoints
The RESTful API application has been divided into the following packages:

* **common** - Implements some utility functions and provides initialization logic for the
application
* **controllers** - Implements the application’s application handlers
* **data** - Implements the persistence logic with the MongoDB database
* **models** - Describes the data model of the application
* **routers** - Implements the HTTP request routers for the RESTful API

## API endpoints

| URI   	|      HTTP Verb      	|  Functionality 	|
|----------	|:-------------:	|------:	|
| /user/register 	|  POST 	| Creates a new user 	|
| /users/login 	|    POST   |  User logs in to the system, which returns a JWT if the login is successful.|
| /cities 	| POST 	| Creates a new city	|
| /cities | PUT | Updates an existing city |
| /cities | GET | Get all cities
| /cities/{id} | GET | Gets single city for a given ID. The value of the comes from the route parameter

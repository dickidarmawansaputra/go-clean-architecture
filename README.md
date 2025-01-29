# Go Clean Architecture With Fiber

<p>Learn Go Clean Architecture with Fiber Framework</p>

## HTTP Methods for RESTful APIs

| HTTP Method | Description                         |
| ----------- | ----------------------------------- |
| GET         | Retrieve a resource.                |
| POST        | Create a new resource.              |
| PUT         | Update a resource (full update).    |
| PATCH       | Update a resource (partial update). |
| DELETE      | Remove a resource.                  |

## Error Handling and Standard HTTP Status Codes

| Status Code               | Description                                       |
| ------------------------- | ------------------------------------------------- |
| 200 OK                    | Successfully returned data.                       |
| 201 Created               | Successfully created a resource.                  |
| 204 No Content            | Successfully deleted a resource.                  |
| 400 Bad Request           | Invalid request (e.g., malformed JSON).           |
| 401 Unauthorized          | Authentication required.                          |
| 403 Forbidden             | Access denied.                                    |
| 404 Not Found             | Resource not found.                               |
| 409 Conflict              | Conflict with the current state of a resource.    |
| 422 Unprocessable Entity  | Validation or processing error.                   |
| 500 Internal Server Error | Server encountered an error.                      |
| 503 Service Unavailable   | Server temporarily unavailable.                   |

## Pattern

## Tech Stack & Tools

- Framework: Fiber
- Git Convention: Husky Hooks
- Database: PostgreSQL
- ORM: GORM
- Config: Viper
- Validator: Go Validator
- Log: Logrus
- Test: Testify
- APIs Docs: Swagger
- Container: Docker

## Author

Dicki D. Saputra

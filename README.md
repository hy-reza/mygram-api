# MyGram REST API

MyGram is an application that allows users to store photos and create comments on other people's photos. The application is equipped with CRUD processes on tables and flows as described below. Authentication is required to access data on the SocialMedia, Photo, and Comment tables using JsonWebToken. Authorization is required for modifying ownership data, such as update or delete.

## Endpoints

### Authentication

| Method | Endpoint                   | Description                       |
| ------ | -------------------------- | --------------------------------- |
| POST   | `/api/v1/authentication/register` | Register a new user             |
| POST   | `/api/v1/authentication/login`    | Login and retrieve access token |

### Users

| Method | Endpoint          | Description         |
| ------ | ----------------- | ------------------- |
| GET    | `/api/v1/users/list` | Get a list of users |

### Photos

| Method | Endpoint                     | Description                  |
| ------ | ---------------------------- | ---------------------------- |
| POST   | `/api/v1/photos/`            | Create a new photo           |
| GET    | `/api/v1/photos/`            | Get a list of photos         |
| GET    | `/api/v1/photos/:id`         | Get a specific photo         |
| PUT    | `/api/v1/photos/:id`         | Update a specific photo      |
| DELETE | `/api/v1/photos/:id`         | Delete a specific photo      |

### Comments

| Method | Endpoint                     | Description                   |
| ------ | ---------------------------- | ----------------------------- |
| POST   | `/api/v1/comments/`          | Create a new comment          |
| GET    | `/api/v1/comments/`          | Get a list of comments        |
| GET    | `/api/v1/comments/:id`       | Get a specific comment        |
| PUT    | `/api/v1/comments/:id`       | Update a specific comment     |
| DELETE | `/api/v1/comments/:id`       | Delete a specific comment     |

### Media

| Method | Endpoint                     | Description                   |
| ------ | ---------------------------- | ----------------------------- |
| POST   | `/api/v1/media/`             | Create a new media            |
| GET    | `/api/v1/media/`             | Get a list of media           |
| GET    | `/api/v1/media/:id`          | Get a specific media          |
| PUT    | `/api/v1/media/:id`          | Update a specific media       |
| DELETE | `/api/v1/media/:id`          | Delete a specific media       |

## Access Documentation

- Swagger: https://mygram-api-production-09d0.up.railway.app/swagger/index.html#/
- Postman: https://documenter.getpostman.com/view/15041975/2s93Xx1k31

To run the project locally, follow these steps:

1. Clone the repository.
2. Install the required dependencies.
3. Set up the database and environment variables.
4. Run the application using `go run main.go`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you find this project useful, please consider giving it a ⭐️ on [GitHub](https://github.com/hy-reza/learn-go-jwt). Your support is greatly appreciated! 😄





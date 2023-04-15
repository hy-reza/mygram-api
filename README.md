# MyGram REST API

MyGram adalah aplikasi yang memungkinkan pengguna untuk menyimpan foto dan membuat komentar pada foto orang lain. Aplikasi ini dilengkapi dengan proses CRUD pada tabel dan alur yang dijelaskan di bawah ini. Autentikasi diperlukan untuk mengakses data pada tabel SocialMedia, Photo, dan Comment menggunakan JsonWebToken. Proses otorisasi diperlukan untuk memodifikasi data kepemilikan seperti update atau delete.

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

## Dokumentasi Akses

- Swagger: https://mygram-api-production-09d0.up.railway.app/swagger/index.html#/
- Postman: https://documenter.getpostman.com/view/15041975/2s93Xx1k31


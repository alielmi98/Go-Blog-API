# Go-Blog-API

## About
This project is a part of the [Blogging Platform API](https://roadmap.sh/projects/blogging-platform-api) project on roadmap.sh, which is a beginner-level backend project.

The implementation of this RESTful API using Go and Gin, with MongoDB for data storage and Docker for deployment, results in a scalable, maintainable, and easily deployable application.

## Features
- **CRUD Operations**: The API allows users to create, read, update, and delete blog posts.
- **Filtering**: Users can filter blog posts by search terms.
- **MongoDB Integration**: The application uses MongoDB as the database for storing blog posts.
- **Docker Deployment**: The application can be easily deployed using Docker.

## Technologies Used
- **Go**: The programming language used for the API implementation.
- **Gin Web Framework**: A popular Go web framework used for building the API.
- **MongoDB**: The NoSQL database used for storing blog post data.
- **Docker**: Used for containerizing the application for easy deployment.

## Getting Started
To set up the project, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/alielmi98/Go-Blog-API.git
```

2. Navigate to the project directory:

```bash
cd Go-Blog-API
```

3. Start the application using Docker Compose:

```bash
docker-compose up
```

This will start the API server and the MongoDB database.

4. The API can be accessed at `http://localhost:8080`.

## API Endpoints
The following API endpoints are available:

- `POST /posts`: Create a new blog post
- `PUT /posts/{id}`: Update an existing blog post
- `DELETE /posts/{id}`: Delete a blog post
- `GET /posts/{id}`: Get a single blog post
- `GET /posts`: Get all blog posts
- `GET /posts?term={searchTerm}`: Filter blog posts by search term

For detailed information about the API, please refer to the [Blogging Platform API](https://roadmap.sh/projects/blogging-platform-api) project on roadmap.sh.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the [MIT License](LICENSE).
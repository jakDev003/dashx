# Project Name: DashBoardX

DashBoardX is a dynamic dashboard web application that provides an interactive and customizable interface for viewing and managing various data and configurations. The project consists of a Go Web API backend and a React frontend, both containerized using Docker.

## Features
- Dynamic tiles on the main page that display various content.
- CRUD operations for records stored in MongoDB.
- Image upload and management.
- Configuration import and export.
- Health check for backend and database connection.
- Dynamic logo and text in the navigation bar.

## Requirements
- Docker
- Docker Compose

## Setup Instructions

### Clone the Repository
```sh
git clone <repository_url>
cd <repository_directory>
```

### Environment Variables
Create a `.env` file in the root directory and add the following variables:
```
# Backend
MONGO_INITDB_ROOT_USERNAME=root
MONGO_INITDB_ROOT_PASSWORD=example
MONGO_DB=dashboardx
LOGO_TEXT="DashBoardX"

# Frontend
REACT_APP_API_URL=http://localhost:8080/api
LOGO_TEXT="DashBoardX"
```

### Docker Compose
Use Docker Compose to build and run the application.
```sh
docker-compose up --build
```

### Access the Application
- Frontend: http://localhost:3000
- Backend: http://localhost:8080

## Endpoints

### Backend API
- `GET /api/health`: Health check endpoint.
- `POST /api/change-db`: Change the database connection string.
- `POST /api/record`: Create a new record.
- `GET /api/records`: Get all records.
- `DELETE /api/record/{guid}`: Delete a record by GUID.
- `PUT /api/record/{guid}`: Update a record by GUID.
- `POST /api/upload`: Upload an image.
- `GET /api/export`: Export all records as JSON.
- `POST /api/import`: Import records from a JSON file.

### Frontend
- **Home**: Main page with dynamic tiles.
- **Configuration**: Change database connection string.
- **Help**: Redirects to this README file.

## License
This project is licensed under the MIT License.
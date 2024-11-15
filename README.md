# Retail Pulse Backend Assignment (Kirana Club)

This project implements a backend service for handling and processing image-related job submissions for retail store visits. The service receives job submissions, validates the data, and processes images asynchronously. It provides endpoints for submitting jobs and checking the status of each job.

---

## Table of Contents
- [Project Overview](#project-overview)
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
  - [Submit a Job](#submit-a-job)
  - [Check Job Status](#check-job-status)
- [Directory Structure](#directory-structure)
- [API Endpoints](#api-endpoints)
  - [POST /api/submit](#post-apisubmit)
  - [GET /api/status](#get-apistatus)
- [Job Processing](#job-processing)
- [Docker Deployment](#docker-deployment)
- [Error Handling](#error-handling)
- [Assumptions](#assumptions)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## Project Overview

The Retail Pulse Backend service is designed to:
- Accept job submissions containing information about store visits, including store IDs and image URLs.
- Validate each job by checking that each store ID is present in a predefined store master data.
- Process each image URL for each store visit by downloading the image and calculating its perimeter based on width and height.
- Provide a mechanism for clients to query the status of each job.

---

## Features
- **Job Submission**: Submit a list of store visits with associated image URLs.
- **Job Status Checking**: Check the processing status of a specific job.
- **Concurrency**: Processes images concurrently with simulated delays.
- **Error Handling**: Returns appropriate error responses for invalid input or failed processing steps.
- **Docker Support**: Easily deployable using Docker.

---

## Architecture

The system follows a modular structure with components for handling HTTP requests, managing jobs, processing images, and validating store data. The main components include:
- **Job Manager**: Manages job submission, status tracking, and concurrency.
- **Store Validation**: Validates that each store ID in the job submission exists in the store master data.
- **Image Processor**: Downloads images and performs computations on each image.

---

## Tech Stack
- **Language**: Go (Golang)
- **HTTP Server**: `net/http` package
- **Data Encoding/Decoding**: `encoding/json`
- **Concurrency**: Goroutines and channels for job processing
- **Containerization**: Docker

---

## Installation

# Clone the Repository

```bash
git clone https://github.com/your-username/retail-pulse-backend.git
cd retail-pulse-backend
```

# Install Dependencies

```bash
go mod tidy
```

# Run the Application

```bash
go run main.go
```

# Run in Docker If Docker is installed, you can run the service in a Docker container.

## Usage

# Submit a Job

- Endpoint: POST /api/submit
1. Request Payload:
```bash
{
  "count": 2,
  "visits": [
    {
      "store_id": "store123",
      "image_url": ["https://example.com/image1.jpg", "https://example.com/image2.jpg"],
      "visit_time": "2024-10-10T15:00:00Z"
    },
    {
      "store_id": "store456",
      "image_url": ["https://example.com/image3.jpg"],
      "visit_time": "2024-10-11T12:00:00Z"
    }
  ]
}
```
2. Response:
```bash
{
  "job_id": "job-1"
}
```

# Check Job Status

- Endpoint: GET /api/status
- Query Parameter: jobid (e.g., /api/status?jobid=job-1)
- Response (example for completed job):
```bash
{
  "job_id": "job-1",
  "status": "completed",
  "error": []
}
```

## Directory Structure

```bash
retail-pulse-backend/
├── main.go                   // Entry point of the application
├── handlers/
│   ├── job_handler.go        // API handlers for submit and status endpoints
├── jobs/
│   ├── job_manager.go        // Logic for job creation, tracking, and processing
│   ├── image_processor.go    // Logic for downloading images and calculating perimeter
├── store/
│   ├── store_master.go       // Logic for loading and validating store master data
├── utils/
│   ├── response.go           // Utility for handling JSON responses
├── Dockerfile                // Docker setup
├── go.mod                    // Go modules configuration
├── go.sum                    // Go modules dependencies
└── README.md                 // Documentation and setup instructions
```

## API Endpoints

# POST /api/submit
- Description: Submits a job containing multiple store visits, each with store ID, image URLs, and visit time.
1. Request Body:
```bash
{
  "count": 2,
  "visits": [
    {
      "store_id": "store123",
      "image_url": ["https://example.com/image1.jpg", "https://example.com/image2.jpg"],
      "visit_time": "2024-10-10T15:00:00Z"
    },
    {
      "store_id": "store456",
      "image_url": ["https://example.com/image3.jpg"],
      "visit_time": "2024-10-11T12:00:00Z"
    }
  ]
}
```
2. Response:
```bash
{
  "job_id": "job-1"
}
```

- Error Responses:
1. 400 Bad Request for invalid JSON or if the count does not match the length of the visits array.
2. 400 Bad Request if any store_id is invalid.

# GET /api/status

- Description: Retrieves the status of a submitted job by jobid.
1. Query Parameter:
- jobid: The unique identifier of the job.
- Response (for completed job):
```bash
{
  "job_id": "job-1",
  "status": "completed",
  "error": []
}
```
2. Error Responses:
- 400 Bad Request if the jobid is invalid or not found.

## Job Processing

The job processing system works as follows:

- **Submit Job**: The client submits a job with multiple store visits.
- **Validate Data**: Each store_id is validated against a master list loaded at startup.
- **Image Processing**: For each valid store visit, the associated images are downloaded, and the perimeter is calculated.
- **Concurrency**: Job processing is handled asynchronously, allowing multiple jobs to run simultaneously.
- **Job Completion**: The job status is updated to "completed" or "failed" based on the result.

## Docker Deployment

1. Build the Docker Image
```bash
docker build -t retail-pulse-backend .
```

2. Run the Docker Container
```bash
docker run -p 8080:8080 retail-pulse-backend
```

3. Access the API
Submit jobs and check job statuses via http://localhost:8080/api/submit and http://localhost:8080/api/status.

## Error Handling

- **Invalid Store ID**: If a store ID in the job submission is not found in the store master data, the request is rejected.
- **Image Download Failure**: If any image fails to download or process, the job's status is updated to "failed" with an error message.
- **Invalid Job ID**: If an invalid jobid is requested, a 400 Bad Request is returned with an appropriate message.

## Assumptions

- Each job has a unique identifier and is processed independently.
- Store master data (store_master.json) is preloaded and remains constant during runtime.
- The perimeter calculation simulates image processing; it does not involve complex image recognition or analysis.
- Jobs are processed sequentially, but individual images within a job are processed concurrently.
- Image URLs are assumed to be accessible and valid URLs.

## Future Enhancements

- **Database Integration**: Use a database for storing job statuses and other information persistently.
- **Error Logging and Monitoring**: Integrate logging and monitoring for better error tracking and reporting.
- **Advanced Image Processing**: Incorporate actual image recognition tasks (e.g., object detection or feature extraction).
- **Authentication**: Add authentication and authorization for secure API access.
- **Scaling with Job Queue**: Implement a job queue system like RabbitMQ to handle high loads and improve scalability.
- **Automatic Retry Mechanism**: Implement a retry mechanism for failed image downloads or processing attempts.

## License
- This project is licensed under the MIT License - see the LICENSE file for details.



### Clone the Repository
```bash
git clone https://github.com/your-username/retail-pulse-backend.git
cd retail-pulse-backend

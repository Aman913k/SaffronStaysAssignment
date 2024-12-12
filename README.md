# Saffron Stays - Hotel Management System

A hotel management system built using Go (Golang) and PostgreSQL, offering APIs for creating hotels, managing available dates, and fetching hotel details including occupancy and rate statistics.

## Features

- **Create Hotel**: API to add a new hotel with its details and availability.
- **Get Hotel Details**: API to fetch hotel details by `room_id`, including occupancy and rate statistics.
- **Hotel Availability**: Manage and view the availability of hotels over the next few months.
- **Occupancy Calculation**: Get occupancy percentage over the next 5 months.
- **Rate Statistics**: Get the average, highest, and lowest nightly rates for a hotel.

## Technologies Used

- **Go (Golang)**: Backend framework.
- **PostgreSQL**: Database for storing hotel and availability data.
- **Gin**: Web framework for handling HTTP requests and responses.
- **Gorm**: ORM for database management (if required).
- **Docker**: Containerization (optional for deployment).

## Setup and Installation

### Prerequisites

- Go 1.18+ installed on your machine.
- PostgreSQL installed or a PostgreSQL service to connect to.
- Git for cloning the repository.

### Steps to Run Locally

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/SaffronStaysAssignment.git
   cd SaffronStaysAssignment


2. **Run on localhost:**

  Post request: http://localhost:8000/hotel/create
  GET  request: http://localhost:8000/hotel/id


   

# TourTap

*Your control center for managing group tours!*

TourTap is built with a simple, event-driven setup:

- Backend: Go (REST API + booking workflow logic)
- Message Broker: RabbitMQ (asynchronous event processing)
- Frontend: Vue + TypeScript
- Database: PostgreSQL with sqlc
- Containerization: Docker
- Orchestration: Docker Compose

The application currently runs as a monolith, but the codebase is structured in a modular way. If needed, parts of the system could be split into separate services later.

The project intentionally starts as a monolith to keep deployment simple and operational overhead low.

## Motivation

First-hand experience of handling large quantities of tour bookings manually. This experience made me think how to automate the workflow.

### Goal

- An open source alternative for smaller companies to automate their workflow regarding bookings, without breaking the bank and without the need to be a coding virtuoso.
- Simplicity and mobility: Containerization makes it a breeze to either self-host or run in the cloud.

## Installation

1. Make sure to have **docker** and **docker-compose** installed
2. Clone this repo
3. Create a new .env file (or rename .env.example and change the credentials)
4. From the project root run:
`docker-compose up --build`

- Remember to backup your database!

## How to use

- Go to <http://localhost:3000/>
- Log in with the demo credentials
  - Email: `test@email.com`
  - Password: `password`

- Create mock bookings at:
  <http://localhost:3000/booking>

- View active bookings under the **Bookings** tab

## Workflow

- New bookings pop up under the **Pending** tab.
- Accepting a request creates a booking.
- If more groups join the same tour on the same date they are merged into the same booking.
- Active bookings can be viewed and filtered by date under the **Bookings** tab.

## Noteworthy features

- Secure API with JWT authentication and refresh tokens.
- Real-time toast notifications via SSE.
- Event-driven booking workflow using RabbitMQ
- Fully containerized setup with database, broker, migrations etc

- Includes automated database migrations and demo user seeding for local development.

## Status/Future of TourTap

This project is in active development. Contributions and feedback are welcome.

### Planned features

- Fully implement the auth package front/back
- Automatic Email Notifications
- PayPal integration
- Individual group detail page
- Admin page for creating users
- Expanded test suite
- Implement https

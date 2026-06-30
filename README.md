# 🚗 SpotSync API

A RESTful Parking Reservation Management API built with **Go (Golang)**, **Echo Framework**, **GORM**, **PostgreSQL (Neon)**, and **JWT Authentication**.

This project was developed as part of the Backend Development Assignment.

---

#  Project Links

## ✅ GitHub Repository

https://github.com/Mokarama/assignment-6-spotsync-api.git

---

##  Live Deployment

https://assignment-6-spotsync-api.onrender.com/

Health Check:

```json
{
  "message": "SpotSync API is running successfully!"
}
```

---

##  Interview Video

https://drive.google.com/drive/folders/1q64us1OUvCr0i-jUMeYkX4gISrRDDSE6

---

#  Features

##  Authentication

- User Registration
- User Login
- Password Hashing (bcrypt)
- JWT Authentication

---

##  Role Based Authorization

### Admin

- Create Parking Zone
- Update Parking Zone
- Delete Parking Zone

### Driver

- Create Reservation
- View Reservations
- Cancel Reservation

---

##  Parking Zone Management

- Create Parking Zone
- Get All Parking Zones
- Get Zone Details
- Update Zone
- Delete Zone

---

##  Reservation Management

- Create Reservation
- Get All Reservations
- Get Reservation By ID
- Get My Reservations
- Cancel Reservation
- Delete Reservation

---

##  Smart Features

- Available Parking Spot Calculation
- Parking Capacity Validation
- Reservation Ownership Check
- Database Transactions
- Row Locking (SELECT FOR UPDATE)

---

#  Tech Stack

- Go (Golang)
- Echo Framework
- GORM
- PostgreSQL (Neon)
- JWT
- bcrypt
- Render
- Git & GitHub

---

#  Project Structure

```text
assignment-6-spotsync-api
│
├── cmd/
│   └── server/
├── config/
├── database/
├── dto/
├── handler/
├── middleware/
├── models/
├── repository/
├── routes/
├── service/
├── .env
├── go.mod
└── README.md
```

---

#  Installation

Clone the repository

```bash
git clone https://github.com/Mokarama/assignment-6-spotsync-api.git
```

Go to project

```bash
cd assignment-6-spotsync-api
```

Install dependencies

```bash
go mod tidy
```

Create a `.env` file

```env
DATABASE_URL=your_neon_database_url
JWT_SECRET=your_secret_key
PORT=8080
```

Run the project

```bash
go run ./cmd/server
```

The server will start at

```text
http://localhost:8080
```

---

#  API Endpoints

## Authentication

| Method | Endpoint |
|---------|----------|
| POST | `/api/v1/auth/register` |
| POST | `/api/v1/auth/login` |

### Parking Zones

| Method | Endpoint |
|---------|----------|
| GET | `/api/v1/zones` |
| GET | `/api/v1/zones/:id` |
| POST | `/api/v1/zones` |
| PATCH | `/api/v1/zones/:id` |
| DELETE | `/api/v1/zones/:id` |

### Reservations

| Method | Endpoint |
|---------|----------|
| POST | `/api/v1/reservations` |
| GET | `/api/v1/reservations` |
| GET | `/api/v1/reservations/:id` |
| GET | `/api/v1/reservations/my-reservations` |
| PATCH | `/api/v1/reservations/:id/cancel` |
| DELETE | `/api/v1/reservations/:id` |

---

#  Security Features

- JWT Authentication
- Password Hashing using bcrypt
- Role-Based Authorization
- Reservation Ownership Validation
- Database Transactions
- Row Locking (SELECT ... FOR UPDATE)

---

#  Deployment

**Platform:** Render

**Live API:**

https://assignment-6-spotsync-api.onrender.com/

---

#  Testing

The API was tested using **Postman**.

---

#  License

This project was developed for educational purposes as part of a backend development assignment.
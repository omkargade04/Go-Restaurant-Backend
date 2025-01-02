# Restaurant Management Backend

A robust REST API backend system for restaurant management built with Go, Gin framework, and MongoDB. This system handles various aspects of restaurant operations including user management, food items, menu management, order processing, and invoicing.

## Features

- User Authentication & Authorization
- Table Management
- Menu Management
- Food Items Management
- Order Processing
- Order Items Tracking
- Invoice Generation
- Middleware for Authentication

## Tech Stack

- **Go** - Primary programming language
- **Gin** - Web framework
- **MongoDB** - Database
- **JWT** - Authentication

## Prerequisites

Before running this project, make sure you have the following installed:
- Go (1.16 or later)
- MongoDB
- Git

## Project Structure

```
go-restro-backend/
├── controllers/
├── database/
├── helpers/
├── middleware/
├── models/
├── routes/
├── main.go
└── README.md
```

### Models
- User
- Table
- Order
- OrderItem
- Menu
- Food
- Invoice
- Note

## API Endpoints

### User Routes
- `GET /users` - Get all users
- `GET /users/:user_id` - Get specific user
- `POST /users/signup` - Register new user
- `POST /users/login` - User login

### Table Routes
- `GET /tables` - Get all tables
- `GET /tables/:table_id` - Get specific table
- `POST /tables` - Create new table
- `PATCH /tables/:table_id` - Update table

### Food Routes
- `GET /foods` - Get all food items
- `GET /foods/:food_id` - Get specific food item
- `POST /foods` - Add new food item
- `PATCH /foods/:food_id` - Update food item

### Menu Routes
- `GET /menus` - Get all menus
- `GET /menus/:menu_id` - Get specific menu
- `POST /menus` - Create new menu
- `PATCH /menus/:menu_id` - Update menu

### Order Routes
- `GET /orders` - Get all orders
- `GET /orders/:order_id` - Get specific order
- `POST /orders` - Create new order
- `PATCH /orders/:order_id` - Update order

### Order Items Routes
- `GET /orderItems` - Get all order items
- `GET /orderItems/:orderItem_id` - Get specific order item
- `GET /orderItems-order/:order_id` - Get items by order
- `POST /orderItems` - Create order items
- `PATCH /orderItems/:orderItem_id` - Update order items

### Invoice Routes
- `GET /invoices` - Get all invoices
- `GET /invoices/:invoice_id` - Get specific invoice
- `POST /invoices` - Create new invoice
- `PATCH /invoices/:invoice_id` - Update invoice

## Setup and Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/go-restro-backend.git
```

2. Navigate to the project directory
```bash
cd go-restro-backend
```

3. Install dependencies
```bash
go mod download
```

4. Set up environment variables
```bash
export PORT=8080
export SECRET_KEY=your-secret-key
```

5. Run the application
```bash
go run main.go
```

The server will start running at `http://localhost:8080`

## Environment Variables

Create a `.env` file in the root directory and add the following:

```env
PORT=8080
MONGODB_URL=your_mongodb_connection_string
SECRET_KEY=your_jwt_secret_key
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Protected routes require a valid JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

# How to Run

> by Shercosta

## Prerequisites

- Go installed
- A MySQL server
- An empty database created for this project

## Setup

1. Clone the repository.
2. Run the following command to install all dependencies:

   ```bash
   go mod tidy
   ```

3. Update the `.env.local` file with your database credentials and the port you want the API to run on.

4. Make sure the database specified in `.env.local` already exists. The application will automatically create the required tables when it starts.

5. Import the Postman collection (`*.json`) from the project root into Postman.

6. Start the application:

   ```bash
   go run main.go
   ```

if things go well, you will automatically run the migrations needed, and then you can try the requests in the postman collection.

_No Agentic AI was used, i dont even use Claude (or vibe code for that matter). All code was templated from my previous project._

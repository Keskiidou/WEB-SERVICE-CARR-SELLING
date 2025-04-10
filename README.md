# 🚗 Car Selling Platform - Web Services Architecture

This project is a webservice-based architecture for a **car selling system**, composed of multiple services using different technologies and communication protocols.

---

##  Authentication Service (REST + JWT)

- This service handles **user authentication and authorization**.
- Built with **REST API** and uses **JWT (JSON Web Tokens)** for secure access.
- It includes a **PostgreSQL** database managed via **Docker Compose**.

> ⚠️ If the `GO Auth` folder in this repository does not open properly, you can view the full authentication service source code [here](https://github.com/Keskiidou/AUTH-JWT-GO).

### Docker Setup (`docker-compose.yml`):
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:latest
    container_name: postgres_auth_service
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: root
      POSTGRES_DB: GOAUTH
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

---

##  Bank Service (GraphQL)

- Built with **GraphQL**, this service provides full **CRUD operations** for bank accounts.
- It is responsible for **checking if a user can afford a car** based on their current bank balance.
- This service communicates with the SOAP service for financial calculations.

---

##  Financial Calculation Service (SOAP)

- A classic **SOAP web service** that calculates the **monthly payment** a user will need to make after buying a car.
- It is integrated into the **GraphQL bank service** and called internally.

---

##  Access Control

All operations in the bank and financial services are **secured** — users **must be authenticated via the REST JWT service** to access any features.

---

##  Project Structure Overview

- `AUTH JWT GO`: REST API service for user login/register and token-based access.
- `Bank graphqlWS`: GraphQL service handling user bank data and calling the SOAP service.
- `insurance-soap-go`: SOAP service performing financial calculations.
- `.gitignore`: Standard gitignore config.
- `queries and mutations.txt`: Sample GraphQL operations.
- `champs i probably can play.txt`: Just for fun 😉 (or maybe placeholder data?).

---

##  Getting Started

1. **Start the Authentication Service**:
     for downloading the dependencies 
   ```bash
   go mod tidy 
   ```
   
   ```bash
   go run main.go 
   ```
   
   ```bash
   docker-compose up 
   ```

3. **Run the other services (GraphQL and SOAP)** independently after auth is up and tokens are issued.
   for downloading the dependencies 
   ```bash
   go mod tidy 
   ```
   
   ```bash
   go run main.go 
   ```

---

Feel free to contribute or expand each service as needed!

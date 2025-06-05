# tourism-app-backend-go

This repository contains the **Go-based microservice part** of a tourism software project developed for **Teknofest 2025**.

The overall project backend was developed with a hybrid approach: some modules in Java, others in Go. This repo holds the Go microservice, which complements the main backend by handling specific functionalities.

> ⚠️ Note:  
> This service does not yet implement full microservice infrastructure features such as distributed logging or service discovery. However, it demonstrates solid backend practices with Go, focusing on clear separation of concerns, RESTful API design, and database interaction.

---

## Project Context

- Developed as part of the **Teknofest 2025** tourism software competition project.  
- Hybrid backend architecture: Java + Go (this repo covers the Go part).  
- The Go service covers translation, weather data, and user-generated content management such as reviews and chests.

---

## Features

- **Translation**: Endpoint to translate text.  
- **Weather**: Get weather data by geographical coordinates.  
- **Review Management**: CRUD operations for user reviews on places.  
- **Chest Management**: CRUD for in-app collectible "chests".  
- **Collected Chests**: Track which chests a user has collected.

---

## API Endpoints

- POST `/translate` - Translate text  
- GET `/weather` - Get weather by coordinates  

### Review Endpoints (`/api/review`)

- POST `/save` - Create review  
- GET `/findAll` - List all reviews  
- GET `/find/:id` - Get review by ID  
- GET `/findByPlace/:place_id` - Get reviews by place  
- GET `/findByPlaceAndUser/:place_id/:user_id` - Get reviews by place and user  
- PUT `/update/:id` - Update review  
- DELETE `/delete/:id` - Delete review  

### Chest Endpoints (`/api/chest`)

- POST `/create` - Create chest  
- PUT `/update/:id` - Update chest  
- DELETE `/delete/:id` - Delete chest  
- GET `/find/:id` - Get chest by ID  
- GET `/findAll` - List all chests  

### Collected Chest Endpoints (`/api/collectedChest`)

- POST `/create` - Create collected chest record  
- GET `/findByUser/:id` - Get collected chests for a user  
- PUT `/update/:id` - Update collected chest record  
- DELETE `/delete/:id` - Delete collected chest record  

---

## Tech Stack

- **Language:** Go (Golang)  
- **Web Framework:** Fiber  
- **ORM:** GORM  
- **Database:** PostgreSQL 

---

## Contact

Feel free to reach out to me via [LinkedIn](https://linkedin.com/in/caner-mastan) or [Email](mailto:jcanermastan@gmail.com).

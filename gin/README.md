# Blog API (Go + Gin + MongoDB)

A simple and efficient Blog CRUD API built using:

- Go
- Gin
- MongoDB
- Clean Architecture structure

---

## 📦 Features

- Create Blog
- Get All Blogs
- Get Single Blog
- Update Blog
- Delete Blog
- Environment based configuration
- Hot reload using Air

---

## 🏗 Project Structure

cmd/api/main.go – Entry point  
internal/config – Environment config  
internal/db – MongoDB connection  
internal/models – Blog schema  
internal/repository – Database logic  
internal/handlers – HTTP handlers  
internal/server – Route setup  

---

## ⚙️ Environment Variables

Create a `.env` file:


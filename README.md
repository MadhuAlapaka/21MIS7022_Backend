📁 21MIS7022_Backend - File Sharing & Management System
🚀 Overview
This is a backend system for secure file sharing and management, built using Go (Golang) with PostgreSQL, Redis, and AWS S3. It provides authentication, file upload/download, search, rate limiting, and automatic cleanup of expired files.

📌 Features
✅ User Authentication (JWT-based)
✅ File Upload & Download (AWS S3)
✅ File Search (by name, date, type)
✅ Rate Limiting (100 requests per min per user)
✅ Background Cleanup Job (Deletes expired files)
✅ File Encryption (Before upload)
✅ Real-time Upload Notifications (WebSockets)

🔧 Tech Stack
Backend: Golang

Database: PostgreSQL

Caching: Redis

Storage: AWS S3

Authentication: JWT

Testing: Go Unit Tests

Deployment: Docker & GitHub

🚀 Installation & Setup
1️⃣ Clone the Repository
sh
Copy
Edit
git clone https://github.com/MadhuAlapaka/21MIS7022_Backend.git
cd 21MIS7022_Backend
2️⃣ Configure Environment
Create a .env file and add the required credentials (DB, AWS, JWT secrets, etc.)

3️⃣ Run the Project
sh
Copy
Edit
go mod tidy
go run main.go
4️⃣ Run Tests
sh
Copy
Edit
go test ./tests/...
📌 API Endpoints
Method	Endpoint	Description
POST	/register	Register a new user
POST	/login	User login (JWT)
POST	/upload	Upload a file
GET	/files	List all files
GET	/files/{id}	Download a file
DELETE	/files/{id}	Delete a file
🤝 Contributing
Want to improve this project? Feel free to create a pull request!

📧 Contact: recruitments@trademarkia.com

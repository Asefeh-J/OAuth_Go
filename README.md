# OAuth Go Project

## Overview
This project demonstrates how to use **OAuth 2.0** in a Go application to integrate with Google services such as Gmail and Google Drive. The app provides APIs for user authentication and accessing Google Drive files and Gmail messages.

---

## Features
- **OAuth 2.0 Authentication**: Authenticate users using Google OAuth.
- **Gmail API Integration**: Fetch user emails with essential details.
- **Google Drive API Integration**: List and access files in the user's Google Drive.
- **Gin Web Framework**: Handle API routes and responses.

---

## Project Structure
    ```plaintext
    .
    ├── api/
    │   ├── routes.go         # API routes and handlers
    │   ├── AuthHandler.go    # OAuth authentication logic
    │   ├── GmailHandler.go   # Gmail API integration
    │   ├── DriveHandler.go   # Google Drive API integration
    ├── oauth/
    │   └── config.go         # OAuth configuration
    ├── main.go               # Entry point
    └── assets/
        └── home_logo.png     # Example logo


---

## Requirements
- **Go** (version 1.20 or later)
- Google Cloud project with OAuth credentials
- **Dependencies**:
  - Gin (`github.com/gin-gonic/gin`)
  - OAuth2 (`golang.org/x/oauth2`)
  - Gmail API (`google.golang.org/api/gmail/v1`)
  - Drive API (`google.golang.org/api/drive/v3`)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Asefeh-J/OAuth_Go.git
   cd OAuth_Go
   
2. Install dependencies:
   ```bash
   go mod tidy

3. Set up Google Cloud credentials:

    * Create a project in Google Cloud Console.
    * Enable Gmail and Google Drive APIs.
    * Generate OAuth 2.0 credentials and download the client_secrets.json file.
    * Save client_secrets.json in the project root or reference it in your code.

4. Run the app:
   ```bash
   go run main.go

---

## API Endpoints

### Authentication
- **GET `/api/v1/admin`**: Redirects user to Google OAuth login.
- **GET `/api/v1/admin/callback`**: Handles OAuth callback and exchanges authorization code for a token.

### Gmail
- **GET `/api/v1/emails`**: Fetches the user's recent emails (requires authentication).

### Google Drive
- **GET `/api/v1/drive/files`**: Lists files in the user's Google Drive.

---

## Example Usage

### Authenticate User
1. Visit `/api/v1/admin` in your browser.
2. Log in with a Google account and allow access.
3. Upon success, you'll see a success message.

### Fetch Emails
- Use an HTTP client (e.g., Postman) to call `/api/v1/emails`.

### List Drive Files
- Use an HTTP client to call `/api/v1/drive/files`.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.



    

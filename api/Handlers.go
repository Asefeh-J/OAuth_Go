package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Asefeh-J/OAuth_Go/oauth"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func AuthHandler(c *gin.Context) {
	if oauth.OauthConfig == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OAuth config is not initialized"})
		return
	}
	url := oauth.OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

var userTokens = make(map[string]*oauth2.Token)

func AuthCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/api/v1/error?message=No code in the query string")
		return
	}

	token, err := oauth.OauthConfig.Exchange(c, code)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/api/v1/error?message=Failed to exchange token")
		return
	}

	if !token.Valid() {
		c.Redirect(http.StatusTemporaryRedirect, "/api/v1/error?message=Invalid token")
		return
	}

	// Store the token for the user (e.g., using a session ID or user ID)
	userTokens["user"] = token

	c.String(http.StatusOK, "User authenticated successfully!")
}

func FetchEmailsHandler(c *gin.Context) {
	// Retrieve the token (for demonstration, using a hardcoded user key)
	token, exists := userTokens["user"]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	client := oauth.OauthConfig.Client(c, token)
	service, err := gmail.NewService(c, option.WithHTTPClient(client))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Gmail service"})
		return
	}

	// Fetch the first 10 messages
	msgs, err := service.Users.Messages.List("me").MaxResults(10).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	// Define essential header keys
	essentialKeys := []string{"From", "To", "Subject", "Date"}

	emailDetails := []map[string]interface{}{}
	for _, msg := range msgs.Messages {
		// Fetch the full message including headers
		fullMessage, err := service.Users.Messages.Get("me", msg.Id).Do()
		if err != nil {
			continue // Skip messages that failed to fetch
		}

		// Extract headers and filter for essential keys
		filteredHeaders := map[string]string{}
		for _, header := range fullMessage.Payload.Headers {
			for _, key := range essentialKeys {
				if header.Name == key {
					filteredHeaders[key] = header.Value
				}
			}
		}

		emailDetails = append(emailDetails, map[string]interface{}{
			"ID":      msg.Id,
			"Headers": filteredHeaders,
		})
	}

	c.JSON(http.StatusOK, gin.H{"emails": emailDetails})
}

func FetchDriveFilesHandler(c *gin.Context) {
	// Retrieve the token (for demonstration, using a hardcoded user key)
	token, exists := userTokens["user"]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	client := oauth.OauthConfig.Client(c, token)
	driveService, err := drive.NewService(c, option.WithHTTPClient(client))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Drive service"})
		return
	}

	// Fetch the first 10 files
	fileList, err := driveService.Files.List().Fields("files(id, name)").PageSize(10).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	// Format and return the file details
	files := []map[string]string{}
	for _, file := range fileList.Files {
		files = append(files, map[string]string{
			"ID":   file.Id,
			"Name": file.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func IconHandler(c *gin.Context) {
	iconPath := "./assets/home_logo.png"

	// Check if the file exists
	if _, err := os.Stat(iconPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Icon not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing the icon"})
		return
	}

	c.File(iconPath)
}

func ErrorHandler(c *gin.Context) {
	message := c.Query("message")
	if message == "" {
		message = "An unknown error occurred"
	}
	c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", message))
}

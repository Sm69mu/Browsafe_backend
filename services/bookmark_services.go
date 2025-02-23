package services

import (
	"browsafe_backend/configs"
	"browsafe_backend/models"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func AddBookmarkService(userID string, bookmark models.Bookmark) (*models.Bookmark, error) {
	now := time.Now()
	bookmark.CreateAt = now
	bookmark.UserID = userID
	// Get user's bookmarks collection
	bookmarksRef := configs.FirestoreClient.Collection("users").Doc(userID).Collection("bookmarks")
	// Add the bookmark
	docRef, _, err := bookmarksRef.Add(ctx, map[string]interface{}{
		"URL":      bookmark.URL,
		"CreateAt": bookmark.CreateAt,
		"UserID":   bookmark.UserID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to add bookmark: %v", err)
	}
	// Set the ID
	bookmark.ID = docRef.ID
	return &bookmark, nil
}

func GetBookmarksByUserID(userID string) ([]models.Bookmark, error) {
	bookmarks := []models.Bookmark{}
	// Get bookmarks collection
	iter := configs.FirestoreClient.Collection("users").Doc(userID).Collection("bookmarks").OrderBy("CreateAt", firestore.Desc).Documents(ctx)
	defer iter.Stop()
	// Iterate through doc
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate bookmarks: %v", err)
		}
		var bookmark models.Bookmark
		if err := doc.DataTo(&bookmark); err != nil {
			return nil, fmt.Errorf("failed to parse bookmark: %v", err)
		}
		//bm Id as doc iD
		bookmark.ID = doc.Ref.ID
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}

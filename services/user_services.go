package services

import (
	"browsafe_backend/configs"
	"browsafe_backend/models"
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"golang.org/x/crypto/bcrypt"
)

var ctx = context.Background()

//-------------------------------register user---------------------------------------------

func RegisterUser(email, password, name string) (*models.Users, error) {
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}
	//create FB auth user
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		DisplayName(name).
		EmailVerified(false)

	authUser, err := configs.AuthClient.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error creating Firebase user %v", err)
	}
	//custom claims for metadata
	// claims := map[string]interface{}{
	// 	//"created_at": time.Now().Unix(),
	// 	"last_login": time.Now().Unix(),
	// }
	// if err := configs.AuthClient.SetCustomUserClaims(ctx, authUser.UID, claims); err != nil {
	// 	return nil, fmt.Errorf("error setting user claims: %v", err)
	// }

	//create user doc in firestore
	user := models.Users{
		ID:        authUser.UID,
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		AuthType:  "email",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//save to firestore
	_, err = configs.FirestoreClient.Collection("users").Doc(user.ID).Set(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error saving user to firestore: %v", err)
	}
	return &user, nil
}

//-------------------------------------------sign in with google---------------------------------------

func HandleGoogleSignIn(idToken string) (*models.Users, string, error) {
	//verify the Google Id token
	token, err := configs.AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, "", fmt.Errorf("error verifying Google ID token: %v", err)
	}

	//get or create user
	user, err := getOrcreateGoogleUser(token)
	if err != nil {
		return nil, "", err
	}

	//create a custom token for client
	custmToken, err := configs.AuthClient.CustomToken(ctx, user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("error creating custom token %v", err)
	}
	return user, custmToken, nil
}

//---------------------------------------------get or create new user ---------------------------------------------

func getOrcreateGoogleUser(token *auth.Token) (*models.Users, error) {
	//check user exist or not
	existingUser, err := configs.AuthClient.GetUser(ctx, token.UID)
	if err == nil {
		doc, err := configs.FirestoreClient.Collection("users").Doc(existingUser.UID).Get(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting existing user from firestore: %v", err)
		}
		var userData models.Users
		if err := doc.DataTo(&userData); err != nil {
			return nil, fmt.Errorf("error parsing existing user data: %v", err)
		}
		return &userData, nil
	}

	//create new user if does not exist
	params := (&auth.UserToCreate{}).
		UID(token.UID).
		Email(token.Claims["email"].(string)).
		DisplayName(token.Claims["name"].(string)).
		PhotoURL(token.Claims["picture"].(string))

	authUser, err := configs.AuthClient.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error creating new user: %v", err)
	}

	user := models.Users{
		ID:        authUser.UID,
		Name:      authUser.DisplayName,
		Email:     authUser.Email,
		AuthType:  "google",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = configs.FirestoreClient.Collection("users").Doc(user.ID).Set(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error saving user to firestore: %v", err)
	}
	return &user, nil
}

//-------------------------------------------------login user ------------------------------------------------------

func LoginUser(email, password string) (string, error) {
	//get user by email
	users, err := configs.AuthClient.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("error getting user details: %v", err)
	}
	//get user from FS to check password
	doc, err := configs.FirestoreClient.Collection("users").Doc(users.UID).Get(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting user from firestore: %v", err)
	}
	var userData models.Users
	if err := doc.DataTo(&userData); err != nil {
		return "", fmt.Errorf("error parsing userData: %v", err)
	}
	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}
	//create custom token
	token, err := configs.AuthClient.CustomToken(ctx, users.UID)
	if err != nil {
		return "", fmt.Errorf("error creating custom token: %v", err)
	}
	return token, nil
}

//---------------------------------------------------update user deatils -------------------------------

func UpdateUserService(userID string, updates models.Users) (*models.Users, error) {
	// Use map data to update db
	updateMap := make(map[string]interface{})
	// Only update non-empty fields
	if updates.Email != "" {
		updateMap["Email"] = updates.Email
	}
	if updates.Password != "" {
		updateMap["Password"] = updates.Password
	}
	if updates.Name != "" {
		updateMap["Name"] = updates.Name
	}
	// If no fields to update, return early
	if len(updateMap) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}
	// Add timestamp
	updateMap["UpdatedAt"] = time.Now()
	// Get doc by id
	docRef := configs.FirestoreClient.Collection("users").Doc(userID)
	// Check doc exists
	_, err := docRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	// Update with merge
	_, err = docRef.Set(ctx, updateMap, firestore.MergeAll)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	// Get updated doc
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated user: %v", err)
	}
	// Convert doc to model
	var updatedUser models.Users
	err = doc.DataTo(&updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to convert document: %v", err)
	}
	return &updatedUser, nil
}

//------------------------------------------------get user by details-----------------------------------

func GetUserDetailsByID(userID string) (*models.Users, error) {
	//validate userID
	if userID == "" {
		return nil, fmt.Errorf("UserID is empty")
	}
	//get doc by id
	docRef := configs.FirestoreClient.Collection("users").Doc(userID)
	//get doc
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}
	//convert doc to model
	var user models.Users
	if err := doc.DataTo(&user); err != nil {
		return nil, fmt.Errorf("unable to convert document to user model: %v", err)
	}
	//set doc id
	user.ID = doc.Ref.ID
	return &user, nil
}

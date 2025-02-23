package configs

import (
    "context"
    "encoding/base64"
    "log"
    "os"
    "path/filepath"

    "cloud.google.com/go/firestore"
    firebase "firebase.google.com/go"
    "firebase.google.com/go/auth"
    "google.golang.org/api/option"
)

var (
    FirestoreClient *firestore.Client
    AuthClient      *auth.Client
)

func InitFirebase() {
    ctx := context.Background()
    var opt option.ClientOption
    if credJSON := os.Getenv("FIREBASE_CREDENTIALS"); credJSON != "" {
        // Decode base64 encoded credentials if needed
        decoded, err := base64.StdEncoding.DecodeString(credJSON)
        if err == nil {
            opt = option.WithCredentialsJSON(decoded)
        } else {
            opt = option.WithCredentialsJSON([]byte(credJSON))
        }
    } else {
        // Fallback to file
        cwd, err := os.Getwd()
        if err != nil {
            log.Fatal("Error getting working directory")
        }

        credentialPaths := []string{
            filepath.Join(cwd, "firebase.json"),
            filepath.Join(cwd, "../firebase.json"),
        }
        var credentialsFile string
        for _, path := range credentialPaths {
            if _, err := os.Stat(path); err == nil {
                credentialsFile = path
                break
            }
        }
        if credentialsFile == "" {
            log.Fatal("Firebase credentials not found in file or environment")
        }
        opt = option.WithCredentialsFile(credentialsFile)
    }
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("error init firebase: %v", err)
    }
    FirestoreClient, err = app.Firestore(ctx)
    if err != nil {
        log.Fatalf("error init firestoreClient: %v", err)
    }
    AuthClient, err = app.Auth(ctx)
    if err != nil {
        log.Fatalf("error init auth: %v", err)
    }
    log.Println("Firebase init success")
}
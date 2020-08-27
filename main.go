package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/karansinghgit/feelrGo/graphql"
	"github.com/karansinghgit/feelrGo/graphql/generated"
	"google.golang.org/api/option"
)

const defaultPort = ":8080"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

//verifyIDToken is a local function which verifies the ID
func verifyIDToken(ctx context.Context, idToken string) error {
	keyFile := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, keyFile)

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return err
	}

	log.Printf("Verified ID token: %v\n", token)
	return nil
}

// TokenVerification is a Middleware for Token Verification through firebase
func TokenVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			verifyIDToken(c, jwtToken)
			c.Writer.Write([]byte("Verified Token"))
		}
	}
}

func main() {
	r := gin.Default()
	// r.Use(TokenVerification())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(defaultPort)
}

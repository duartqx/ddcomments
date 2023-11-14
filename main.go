package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	r "github.com/duartqx/ddcomments/api/router"
	repo "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"
)

func main() {

	db, err := repo.GetDBConnection()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	mux := r.NewRouterBuilder().SetDb(db).SetSecret([]byte("secret")).Build()

	port := ":8000"

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(
		fmt.Sprintf(`
______________________________________________________________________________

 _____  _____  ______ _______ _______ _______ _______ _______ _______ _______ 
|     \|     \|      |       |   |   |   |   |    ___|    |  |_     _|     __|
|  --  |  --  |   ---|   -   |       |       |    ___|       | |   | |__     |
|_____/|_____/|______|_______|__|_|__|__|_|__|_______|__|____| |___| |_______|


DDComments running @ %s://%s
______________________________________________________________________________
`,
			"http", srv.Addr,
		),
	)

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// Graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)

	// Block until we interrupt signal
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(0)
}

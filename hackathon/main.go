package main

import (
	"database/sql"
	"fmt"
	"hackathon/controller"
	"hackathon/dao"
	"hackathon/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/oklog/ulid/v2"
)

var db *sql.DB

func main() {
	db = initDB()

	userDao := &dao.UserDao{DB: db}
	searchUserController := &controller.SearchUserController{SearchUserUseCase: &usecase.SearchUserUseCase{UserDao: userDao}}
	registerUserController := &controller.RegisterUserController{RegisterUserUseCase: &usecase.RegisterUserUseCase{UserDao: userDao}}
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		switch r.Method {
		case http.MethodGet:
			searchUserController.Handle(w, r)
		case http.MethodPost:
			registerUserController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			fmt.Printf(r.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	messageDao := &dao.MessageDao{DB: db}
	searchMessageController := &controller.SearchMessageController{SearchMessageUseCase: &usecase.SearchMessageUseCase{MessageDao: messageDao}}
	registerMessageController := &controller.RegisterMessageController{RegisterMessageUseCase: &usecase.RegisterMessageUseCase{MessageDao: messageDao}}
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		switch r.Method {
		case http.MethodGet:
			searchMessageController.Handle(w, r)
		case http.MethodPost:
			registerMessageController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	deleteMessageController := &controller.DeleteMessageController{DeleteMessageUseCase: &usecase.DeleteMessageUseCase{MessageDao: messageDao}}
	http.HandleFunc("/delete_message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			deleteMessageController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	EditMessageController := &controller.EditMessageController{EditMessageUseCase: &usecase.EditMessageUseCase{MessageDao: messageDao}}
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		switch r.Method {
		case http.MethodPost:
			EditMessageController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	channelDao := &dao.ChannelDao{DB: db}
	searchChannelController := &controller.SearchChannelController{SearchChannelUseCase: &usecase.SearchChannelUseCase{ChannelDao: channelDao}}
	registerChannelController := &controller.RegisterChannelController{RegisterChannelUseCase: &usecase.RegisterChannelUseCase{ChannelDao: channelDao}}
	getChannelController := &controller.GetChannelController{GetChannelUseCase: &usecase.GetChannelUseCase{ChannelDao: channelDao}}

	http.HandleFunc("/channel", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		switch r.Method {
		case http.MethodGet:
			searchChannelController.Handle(w, r)
		case http.MethodPost:
			registerChannelController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
	http.HandleFunc("/getchannels", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getChannelController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
	deleteChannelController := &controller.DeleteChannelController{DeleteChannelUseCase: &usecase.DeleteChannelUseCase{ChannelDao: channelDao}}
	http.HandleFunc("/delete_channel", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		switch r.Method {
		case http.MethodDelete:
			deleteChannelController.Handle(w, r)
		default:
			log.Printf("BadRequest(status code = 400)")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	closeDBWithSysCall()

	log.Println("Listening...")
	//if err := http.ListenAndServe(":6000", nil); err != nil {
	//	log.Fatal(err)
	//}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func initDB() *sql.DB {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	//mysqlUser := "uttc"
	//mysqlPwd := "ramen102"
	//mysqlHost := "34.27.193.191:3306"
	//mysqlDatabase := "hackathon"

	//mysqlUser := "test_user"
	//mysqlPwd := "password"
	//mysqlHost := "(localhost:3306)"
	//mysqlDatabase := "test_database"

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)

	_db, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	return _db
}

func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

type KarmaApi struct {
	key     string
	redisdb *redis.Client

	http.ServeMux
}

type KarmaResult struct {
	Name  string `json:"name"`
	Karma int64  `json:"karma"`
}

func NewKarmaApi(redisdb *redis.Client, key string) *KarmaApi {
	mux := KarmaApi{redisdb: redisdb, key: key}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		err := redisdb.Ping().Err()
		if err != nil {
			newErr := fmt.Errorf("redis: %s", err)
			mux.errorOut(w, newErr)
			return
		}

		fmt.Fprintf(w, "OK!")
	})

	mux.HandleFunc("/karma", func(w http.ResponseWriter, r *http.Request) {
		result, err := mux.get()
		if err != nil {
			newErr := fmt.Errorf("redis: %s", err)
			mux.errorOut(w, newErr)
			return
		}

		json.NewEncoder(w).Encode(result)
	})

	mux.HandleFunc("/karma/decrement", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result, err := mux.decrement()
		if err != nil {
			newErr := fmt.Errorf("redis: %s", err)
			mux.errorOut(w, newErr)
			return
		}

		json.NewEncoder(w).Encode(result)
	})

	mux.HandleFunc("/karma/increment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result, err := mux.increment()
		if err != nil {
			newErr := fmt.Errorf("redis: %s", err)
			mux.errorOut(w, newErr)
			return
		}

		json.NewEncoder(w).Encode(result)
	})

	return &mux
}

func (a *KarmaApi) errorOut(w http.ResponseWriter, err error) {
	log.Printf("error: %s", err)

	w.WriteHeader(http.StatusInternalServerError)
	msg := fmt.Sprintf("Error! -> %s", err.Error())
	w.Write([]byte(msg))
}

func (a *KarmaApi) decrement() (KarmaResult, error) {
	result, err := a.redisdb.Decr(a.key).Result()
	if err != nil {
		return KarmaResult{}, err
	}

	return KarmaResult{Name: a.key, Karma: result}, nil
}

func (a *KarmaApi) increment() (KarmaResult, error) {
	result, err := a.redisdb.Incr(a.key).Result()
	if err != nil {
		return KarmaResult{}, err
	}

	return KarmaResult{Name: a.key, Karma: result}, nil
}

func (a *KarmaApi) get() (KarmaResult, error) {
	result, err := a.redisdb.IncrBy(a.key, 0).Result()
	if err != nil {
		return KarmaResult{}, err
	}

	return KarmaResult{Name: a.key, Karma: result}, nil
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Print("Starting CiBO Karma server...")

	redisAddr := os.Getenv("REDIS_HOST")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	server := http.NewServeMux()

	apiServer := NewKarmaApi(redisdb, "karma")
	server.Handle("/api/", http.StripPrefix("/api", apiServer))

	fs := http.FileServer(http.Dir("static"))
	server.Handle("/static/", http.StripPrefix("/static", fs))

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "static/index.html")
	})

	http.ListenAndServe(":8080", logger(server))
}

package useroauth2

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type gatewayInfo struct {
	gatewayID string
	secret    string
}

func createClientAuthorize(gi gatewayInfo) *store.ClientStore {
	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set(gi.gatewayID, &models.Client{
		ID:     gi.gatewayID,
		Secret: gi.secret,
		Domain: "",
	})
	return clientStore
}

//Oauth2 ..
func Oauth2() {
	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   15,
	}))

	gio := &gatewayInfo{
		gatewayID: "yt123456",
		secret:    "123456",
	}
	manager.MapClientStorage(createClientAuthorize(*gio))

	srv := server.NewDefaultServer(manager)

	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)

		srv.HandleTokenRequest(w, r)

	})
	log.Fatal(http.ListenAndServe(":9096", nil))

}

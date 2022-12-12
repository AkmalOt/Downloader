package server

import (
	"Uploader/internal/repository"
	logging "Uploader/pkg"
	"context"
	"net/http"
)

const userID = "user_id"

type Auth struct {
	Repository *repository.Repository
}

type contextKey struct {
	key string
}

//type IDFunc func(ctx context.Context, token string) (id int64, err error)
//
//var AuthenticateContextKey = &contextKey{key: "authentication key"}
//
//func Authentication(idFunc IDFunc) func(handler http.Handler) http.Handler {
//	return func(handler http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			token := r.Header.Get("Authorization")
//			id, err := idFunc(r.Context(), token)
//			if err != nil {
//				log.Println(err)
//			}
//			//if errors.Is(err, services.ErrExpired) || errors.Is(err, services.ErrNoAuthorization) {
//			//	log.Println(err)
//			//	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//			//	return
//			//}
//			//if err != nil {
//			//	InternalServerError(w, err)
//			//	return
//			//}
//			ctx := context.WithValue(r.Context(), AuthenticateContextKey, id)
//			r = r.WithContext(ctx)
//			handler.ServeHTTP(w, r)
//		})
//	}
//}

func (s *Server) TokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log := logging.GetLogger()

		//values := request.URL.Query()
		//token := values.Get("token")

		token := request.Header.Get("token")
		if len(token) == 0 {
			log.Println("fuck ")
			http.Error(writer, http.StatusText(http.StatusBadRequest), 400)
			return
		}

		log.Println(token)
		tokenId, userId, err := s.Services.Repository.ValidateToken(token)
		if err != nil {
			log.Println(err)
			return
		}

		request = request.WithContext(context.WithValue(request.Context(), userID, userId))
		log.Println("test in WithContext", tokenId, userId)
		//----------------------------
		//log.Println(token)
		//ctx := request.Context()
		//value := ctx.Value(userID)
		//userIdTest := value.(string)
		//log.Println(userIdTest, userId)

		next.ServeHTTP(writer, request)
	})

}

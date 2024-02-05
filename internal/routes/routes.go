package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	//how are we importing these becuase akil is using a github path??
	ServerConfig "github.com/akhil/dynamo-go-cru-yt/config"
	HealthHandler "github.com/akhil/dynamodb-go-cur-yt/internal/handlers/health"
	ProductHandler "github.com/akhil/dynamodb-go-cru-yt/internal/handlers/product"
 )

type Router struct {
config *Config 
router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout), 
		router: chi.NewRouter(),
	}
}

func r(*Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigsRouters(repository)
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router

}

func (r *Router) setConfigsRouters() {

	r.EnableCORDS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnablerealIP()
}


// so here we are defining the routes 
func (r *Router)RouterHealth (repositroy adapter.Interface){
heandler := HealthHandler.newHandler(repository)

r.router.Route("/health", func(route chi.router){
	route.Post("/", handler.post)
	route.Get("/", handler.Get)
	route.Put("/", handler.Put)
	route.Delete("/", handler.Delete)
	route.Options("/", handler.Options)
})
}

func (r *Router)RouterProduct(repository adapter.Interface) {
	handler := Producthandler.NewHandler(repository)

	r.router.Route("/product", func(router chi.Router){
		route.Post("/", handler.post)
		route.Get("/", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/{ID}", handler.Options)

}

// these are all methods on a struct, which struct is it
func (r *Router)EnableLogger() *Router{
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router{
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router)EnableCORS()*Router{
	r.router.Use(r.config.Cors)
	return r 
}

func (r *Router)EnableReovery()*Router{
	r.router.Use(middleware.Recoverer)
	return r 
}

func (r *Router)EnableRequestId() *Router{
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router)EnableRealIP()*Router{
	r.router.User(middleware.RealIP)
	return r 
}

//blah
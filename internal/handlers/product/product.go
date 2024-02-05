package product

import (
	"errors"
	"http"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	//need the adapter in this function here
	//we also need controller/product
	//we also need internal/entities/product
	//we also need internal/handlers
	//we also need internal/rules
	//and interal/rules/product
	//and also need utils/http
	//some of these will need defining
)

type Handler struct {
	handler.Interface
	Controller product.Interface
	Rules Rules.Interface
}

func NewHandler(repository adapter.Interface) handlers. Interface{
	return &Handler {
		Controller: : product.NewController(repository),
		Rules: RulesProduct.NewRules(), 
	}
}

// just look at this function for now,
// how would you like the opportunity to win a 10000 point trading account to trade with
func (h *Handler)Get(w http.ResponseWriter, r *http.Request) {
if chi.URLParam(r, "ID") != "" {
	h.getONe(w, r)
} else {
	h.getAll(w, r)
}
	

}

//frome routers to handlers to controllers to repository 
func (h *Handler)getOne(http.ResponseWriter, r *http.Request) {
ID, err  := uuid.Parese(chi.URLParam(r, "ID"))

if err!= nil {
	HttpStatus.StatusInternalServerError(w, r, err)
	return
}
HttpStatus.StatusOk(w, r, response)
}

func(h *Handler) getAll(http.ResponseWriter, r, *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
HttpStatus.StatusInternalServerError(w, r, err)
	}
	HttpsStatusOK(w, r, response)
}



// lets do this post function
func (h *Handler) Post (w http.ResponseWriter, r *http.Request){
// once its validated its the correct body to work on???
	productBody, err := h.getBoadyAndValidate(r, uuid.nil )
if err != nil {
	HttpStatus.StatusBadREquest(w, r, err)
	return
}	

ID, err:= h.controller.Create(productBody)

if err != nil {
	HttpStatus.StatusInternalServerError(u,r,err)
	return
}
HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})
}

func Put (h *Handler)(w. http.responseWriter, r *http.Request)  {
	ID, err = uuid.Parse(chi. URLParam(r, "ID"))
if err != nil {
	HttepStatus.StatusBadRequest(w,r errors.New("ID is not uuid valid"))
	return
}
//when we have recieved the new data from the request we want to validate it
producBody, err := h.getBodyAndValidate 
if err != nil {
	HttpStatus.StatusBadRequest(w, r, err)
	return
}

if err := h.Controller.Update(ID, productBody); err != nil{
HttpStatus.StatusInternalServerError(w, r, err)
return}
}


func (h *Handler)Delete (w, r http.ResponseWriter, r *http.request){
	uuid.Parse(chi.URLParam(re, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}
if err != hmcontroller.Remove(ID); err != nil{
	HttpStatus.StatusInternalServerError(w, r, err)
	return

}
HttpStatus.StatusNoContent(w, r)
}

func Options(h *Handler)(w http.ResponseWriter, r *http.request) {
	HttpStatus.StatusNoContent(w, r)
}

//below is marshalling and unmarshalling and setting default values
func (h *Handler) getBodyAndValidate(r, *http.Request, ID uuid.UUID)(*EntityProduct.Product, error) {
	 productBody := &EntityProduct.Product()

	 body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)

	 if err != nil{
		reuturn &EntityProduct{}, errors.New("body is required")
	 }
productParsed, err := EntityProduct.InterfaceToModel(body)
	
if err!= nil{
return &EntityProduct.Product(), errors.New("error converting body to model")
}

setDefaultValues(productParse, ID)
return productParsed, h.Rule.Validate(productParsed)

}

func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID){
	product.UpdatedAt = time.Now()
if ID == uuid.Nil{
	product.ID = uuid.New()
	product.CreatedAt = time.Now
}else {
	product.ID = ID
}
	
}


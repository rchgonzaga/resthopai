package kittiesbundle

import (
	"net/http"
	"strconv"

	"resthopai/app/core"

	"github.com/gorilla/mux"
)

// KittiesController struct
type KittiesController struct {
	core.Controller
	km KittiesMapper
}

// NewKittiesController instance
func NewKittiesController(km KittiesMapper) *KittiesController {
	return &KittiesController{
		Controller: core.Controller{},
		km:         km,
	}
}

// Index func return all kitties in database
func (c *KittiesController) Index(w http.ResponseWriter, r *http.Request) {
	k, err := c.km.FindAll()

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}

// Create a kitty
func (c *KittiesController) Create(w http.ResponseWriter, r *http.Request) {
	var k Kitty

	if err := c.GetContent(&k, r); err != nil {
		return
	}

	if !k.Validate() {
		c.SendJSON(w, k.Errors, http.StatusBadRequest)
		return
	}

	// Insert kitty and handle error
	if c.HandleError(c.km.Insert(&k), w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}

// Delete a kitty
func (c *KittiesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if c.HandleError(err, w) {
		return
	}

	// Insert kitty and handle error
	if c.HandleError(c.km.Delete(id), w) {
		return
	}

	c.SendJSON(w, nil, http.StatusNoContent)
}

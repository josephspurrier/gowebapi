package controller

import (
	"log"
	"net/http"

	"app/webapi/model/user"
	"app/webapi/pkg/form"
	"app/webapi/pkg/response"
	"app/webapi/pkg/router"
)

// Routes.
func init() {
	router.Post("/users", UserOnePOST)
	router.Get("/users/:id", UserOneGET)
	router.Get("/users", UserAllGET)
	router.Put("/users/:id", UserOnePUT)
	router.Delete("/users/:id", UserOneDELETE)
	router.Delete("/users", UserAllDELETE)
}

const (
	itemCreated      = "item created"
	itemExists       = "item already exists"
	itemNotFound     = "item not found"
	itemFound        = "item found"
	itemsFound       = "items found"
	itemsFindEmpty   = "no items to find"
	itemUpdated      = "item updated"
	itemDeleted      = "item deleted"
	itemsDeleted     = "items deleted"
	itemsDeleteEmpty = "no items to delete"

	friendlyError = "an error occurred, please try again later"
)

// *****************************************************************************
// Create
// *****************************************************************************

// UserOnePOST creates a user.
func UserOnePOST(w http.ResponseWriter, r *http.Request) {
	m, err := user.New()
	if err != nil {
		log.Println("UUID Error", err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Validate the required fields are present
	errMsg, err := form.Validate(r, m)
	if err == form.ErrRequiredMissing || err == form.ErrWrongContentType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrBadStruct || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Validate value types and copy values to struct
	errMsg, err = form.StructCopy(r, m)
	if err == form.ErrWrongType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrNotSupported || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Create item
	count, err := m.Create()
	if err == user.ErrExists {
		response.SendError(w, http.StatusBadRequest, itemExists)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	response.Send(w, http.StatusCreated, itemCreated, count, nil)
}

// *****************************************************************************
// Read
// *****************************************************************************

// UserOneGET returns one user.
func UserOneGET(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Get an item
	entity, err := user.Read(ID)
	if err == user.ErrNoResult {
		response.Send(w, http.StatusOK, itemNotFound, 0, nil)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	response.Send(w, http.StatusOK, itemFound, 1, user.Group{*entity})
}

// UserAllGET returns all users.
func UserAllGET(w http.ResponseWriter, r *http.Request) {
	// Get all items
	group, err := user.ReadAll()
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	} else if len(group) < 1 {
		response.Send(w, http.StatusOK, itemsFindEmpty, len(group), nil)
		return
	}

	response.Send(w, http.StatusOK, itemsFound, len(group), group)
}

// *****************************************************************************
// Update
// *****************************************************************************

// UserOnePUT updates a user.
func UserOnePUT(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Get an item
	m, err := user.Read(ID)
	if err == user.ErrNoResult {
		response.Send(w, http.StatusOK, itemNotFound, 0, nil)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Validate the required fields are present
	errMsg, err := form.Validate(r, m)
	if err == form.ErrRequiredMissing {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrBadStruct || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Validate value types and copy values to struct
	errMsg, err = form.StructCopy(r, m)
	if err == form.ErrWrongType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrNotSupported || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Update item
	count, err := m.Update()
	if err == user.ErrNotExist {
		response.SendError(w, http.StatusBadRequest, itemNotFound)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	response.Send(w, http.StatusCreated, itemUpdated, count, nil)
}

// *****************************************************************************
// Delete
// *****************************************************************************

// UserOneDELETE deletes one user.
func UserOneDELETE(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Delete an item
	count, err := user.Delete(ID)
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	} else if count < 1 {
		response.Send(w, http.StatusOK, itemNotFound, count, nil)
		return
	}

	response.Send(w, http.StatusOK, itemDeleted, count, nil)
}

// UserAllDELETE deletes all users.
func UserAllDELETE(w http.ResponseWriter, r *http.Request) {
	// Delete all items
	count, err := user.DeleteAll()
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	} else if count < 1 {
		response.Send(w, http.StatusOK, itemsDeleteEmpty, count, nil)
		return
	}

	response.Send(w, http.StatusOK, itemsDeleted, count, nil)
}

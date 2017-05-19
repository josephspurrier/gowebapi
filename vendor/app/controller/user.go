package controller

import (
	"log"
	"net/http"

	"app/model/user"
	"app/shared/form"
	"app/shared/response"
	"app/shared/router"
)

// Routes
func init() {
	router.Post("/users", UserOnePOST)
	router.Get("/users/:id", UserOneGET)
	router.Get("/users", UserAllGET)
	router.Put("/users/:id", UserOnePUT)
	router.Delete("/users/:id", UserOneDELETE)
	router.Delete("/users", UserAllDELETE)
}

const (
	ItemCreated      = "item created"
	ItemExists       = "item already exists"
	ItemNotFound     = "item not found"
	ItemFound        = "item found"
	ItemsFound       = "items found"
	ItemsFindEmpty   = "no items to find"
	ItemUpdated      = "item updated"
	ItemDeleted      = "item deleted"
	ItemsDeleted     = "items deleted"
	ItemsDeleteEmpty = "no items to delete"

	FriendlyError = "an error occurred, please try again later"
)

// *****************************************************************************
// Create
// *****************************************************************************
func UserOnePOST(w http.ResponseWriter, r *http.Request) {
	m, err := user.New()
	if err != nil {
		log.Println("UUID Error", err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Validate the required fields are present
	err, errMsg := form.Validate(r, m)
	if err == form.ErrRequiredMissing || err == form.ErrWrongContentType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrBadStruct || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Validate value types and copy values to struct
	err, errMsg = form.StructCopy(r, m)
	if err == form.ErrWrongType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrNotSupported || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Create item
	count, err := m.Create()
	if err == user.ErrExists {
		response.SendError(w, http.StatusBadRequest, ItemExists)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	response.Send(w, http.StatusCreated, ItemCreated, count, nil)
}

// *****************************************************************************
// Read
// *****************************************************************************

func UserOneGET(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Get an item
	entity, err := user.Read(ID)
	if err == user.ErrNoResult {
		response.Send(w, http.StatusOK, ItemNotFound, 0, nil)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	response.Send(w, http.StatusOK, ItemFound, 1, user.Group{*entity})
}

func UserAllGET(w http.ResponseWriter, r *http.Request) {
	// Get all items
	group, err := user.ReadAll()
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	} else if len(group) < 1 {
		response.Send(w, http.StatusOK, ItemsFindEmpty, len(group), nil)
		return
	}

	response.Send(w, http.StatusOK, ItemsFound, len(group), group)
}

// *****************************************************************************
// Update
// *****************************************************************************

func UserOnePUT(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Get an item
	m, err := user.Read(ID)
	if err == user.ErrNoResult {
		response.Send(w, http.StatusOK, ItemNotFound, 0, nil)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Validate the required fields are present
	err, errMsg := form.Validate(r, m)
	if err == form.ErrRequiredMissing {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrBadStruct || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Validate value types and copy values to struct
	err, errMsg = form.StructCopy(r, m)
	if err == form.ErrWrongType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	} else if err == form.ErrNotSupported || err == form.ErrNotStruct {
		log.Println(errMsg)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Update item
	count, err := m.Update()
	if err == user.ErrNotExist {
		response.SendError(w, http.StatusBadRequest, ItemNotFound)
		return
	} else if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	response.Send(w, http.StatusCreated, ItemUpdated, count, nil)
}

// *****************************************************************************
// Delete
// *****************************************************************************

func UserOneDELETE(w http.ResponseWriter, r *http.Request) {
	// Get the parameter id
	params := router.Params(r)
	ID := params.ByName("id")

	// Delete an item
	count, err := user.Delete(ID)
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	} else if count < 1 {
		response.Send(w, http.StatusOK, ItemNotFound, count, nil)
		return
	}

	response.Send(w, http.StatusOK, ItemDeleted, count, nil)
}

func UserAllDELETE(w http.ResponseWriter, r *http.Request) {
	// Delete all items
	count, err := user.DeleteAll()
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	} else if count < 1 {
		response.Send(w, http.StatusOK, ItemsDeleteEmpty, count, nil)
		return
	}

	response.Send(w, http.StatusOK, ItemsDeleted, count, nil)
}

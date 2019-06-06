package controllers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"api/models"
	"api/workers"
)

type IBookController interface {
	CreateBook(response http.ResponseWriter, request *http.Request)
	RetrieveBookById(response http.ResponseWriter, request *http.Request)
	RetrieveAllBooks(response http.ResponseWriter, request *http.Request)
	UpdateBookById(response http.ResponseWriter, request *http.Request)
	DeleteBookById(response http.ResponseWriter, request *http.Request)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	book := models.Book{}
	e := json.NewDecoder(request.Body).Decode(&book)
	if e != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
		return
	}
	id, err := workers.CreateBook(book)
	idJson, errJson := json.Marshal(models.BookLocator{Id: id})
	if err != nil || errJson != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
	} else {
		response.Header().Set("Content-Type","application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(idJson)
	}
}

func RetrieveBookById(response http.ResponseWriter, request *http.Request) {
	idContext := request.Context()
	id, e := strconv.Atoi(idContext.Value("id").(string))
	if e != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("internal server error"))
		return
	}
	book, err := workers.RetrieveBookById(id)
	bookJson, errJson := json.Marshal(book)
	if err != nil || errJson != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
	} else {
		response.Header().Set("Content-Type","application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(bookJson)
	}
}

func RetrieveAllBooks(response http.ResponseWriter, request *http.Request) {
	books, err := workers.RetrieveAllBooks()
	booksJson, errJson := json.Marshal(books)
	if err != nil || errJson != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
	} else {
		response.Header().Set("Content-Type","application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(booksJson)
	}
}

func UpdateBookById(response http.ResponseWriter, request *http.Request) {
	idContext := request.Context()
	id, e := strconv.Atoi(idContext.Value("id").(string))
	if e != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("internal server error"))
		return
	}
	book := models.Book{}
	errJson := json.NewDecoder(request.Body).Decode(&book)
	if errJson != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
		return
	}
	err := workers.UpdateBookById(book, id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
	} else {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("ok"))
	}
}

func DeleteBookById(response http.ResponseWriter, request *http.Request) {
	idContext := request.Context()
	id, e := strconv.Atoi(idContext.Value("id").(string))
	if e != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("internal server error"))
		return
	}
	err := workers.DeleteBookById(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("bad request"))
	} else {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("ok"))
	}
}
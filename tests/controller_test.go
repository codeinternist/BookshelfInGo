package controller_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"api/models"
	"encoding/json"
	"net/http"
	"time"
	"bytes"
	"io/ioutil"
	"fmt"
)

func TestCreateSuccess(t *testing.T) {
	book := models.Book{
		Title: "someTitle",
		Author: "someAuthor",
		Publisher: "somePublisher",
		PublishDate: time.Now(),
		Rating: 2,
		Status: models.CheckedIn,
	}
	bookJson, _ := json.Marshal(book)
	bookBytes := bytes.NewBuffer(bookJson)
	resp, err := http.Post("http://localhost:5000/books/", "application/json", bookBytes)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCreateFailInvalidRating(t *testing.T) {
	book := models.Book{
		Title: "someTitle",
		Author: "someAuthor",
		Publisher: "somePublisher",
		PublishDate: time.Now(),
		Rating: 0,
		Status: models.CheckedIn,
	}
	bookJson, _ := json.Marshal(book)
	bookBytes := bytes.NewBuffer(bookJson)
	resp, err := http.Post("http://localhost:5000/books/", "application/json", bookBytes)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestCreateFailInvalidStatus(t *testing.T) {
	book := models.Book{
		Title: "someTitle",
		Author: "someAuthor",
		Publisher: "somePublisher",
		PublishDate: time.Now(),
		Rating: 3,
		Status: 42,
	}
	bookJson, _ := json.Marshal(book)
	bookBytes := bytes.NewBuffer(bookJson)
	resp, err := http.Post("http://localhost:5000/books/", "application/json", bookBytes)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestRetrieveOneSuccess(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:5000/books/%d/", AddBookGetId(t)))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestRetrieveOneFailNotFound(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/books/0/")
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestRetrieveAllSuccess(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/books/")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdateSuccess(t *testing.T) {
	newBook := GetDifferentBook(2)
	newBookJson, _ := json.Marshal(newBook)
	newBookBytes := bytes.NewBuffer(newBookJson)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:5000/books/%d/", AddBookGetId(t)), newBookBytes)
	assert.Nil(t, err)

	resp, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdateFailNotFound(t *testing.T) {
	newBook := GetDifferentBook(2)
	newBookJson, _ := json.Marshal(newBook)
	newBookBytes := bytes.NewBuffer(newBookJson)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:5000/books/%d/", 0), newBookBytes)
	assert.Nil(t, err)

	resp, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdateFailInvalidRating(t *testing.T) {
	newBook := GetDifferentBook(7)
	newBookJson, _ := json.Marshal(newBook)
	newBookBytes := bytes.NewBuffer(newBookJson)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:5000/books/%d/", AddBookGetId(t)), newBookBytes)
	assert.Nil(t, err)

	resp, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestDeleteSuccess(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:5000/books/%d/", AddBookGetId(t)), nil)
	assert.Nil(t, err)

	resp, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteFail(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/books/0/", nil)
	assert.Nil(t, err)

	resp, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func AddBookGetId(t *testing.T) int {
	book := models.Book{
		Title: "someTitle",
		Author: "someAuthor",
		Publisher: "somePublisher",
		PublishDate: time.Now(),
		Rating: 2,
		Status: models.CheckedIn,
	}
	bookJson, _ := json.Marshal(book)
	bookBytes := bytes.NewBuffer(bookJson)
	resp, err := http.Post("http://localhost:5000/books/", "application/json", bookBytes)
	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	var id models.BookLocator

	err = json.Unmarshal([]byte(jsonData), &id)
	assert.Nil(t, err)
	assert.NotEmpty(t, id.Id)

	return id.Id
}

func GetDifferentBook(rating int) models.Book {
	return models.Book {
		Title: "updatedTitle",
		Author: "updatedAuthor",
		Publisher: "updatedPublisher",
		PublishDate: time.Now(),
		Rating: rating,
		Status: models.CheckedIn,
	}
}
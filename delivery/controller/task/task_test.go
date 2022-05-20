package task

import (
	"encoding/json"
	"errors"
	middlewares "projek/todo/delivery/middleware"

	"net/http"
	"net/http/httptest"
	"projek/todo/entities"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var token string

// INITIATE TOKEN
func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1, "Galih", "Galih@gmail.com")
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"userid":   1,
			"projekid": 1,
			"name":     "nana",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		taskC := NewControlTask(&mockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(taskC.CreateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Task", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"userid":   1,
			"projekid": 1,
			"name":     "nana",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		taskC := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(taskC.CreateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()

		requestBody := "kecantikan"

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		taskC := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(taskC.CreateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{

			"projekid": 1,
			"name":     "nana",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		taskC := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(taskC.CreateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
	})
}

func TestGetAllTask(t *testing.T) {
	t.Run("Success Get All Task", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		task := NewControlTask(&mockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.GetAllTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.GetAllTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
}

func TestGetTaskID(t *testing.T) {
	t.Run("Success Get TaskBy ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		task := NewControlTask(&mockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.GetTaskID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data ID", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.GetTaskID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.GetTaskID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":    "kecantikan",
			"status":  "reopen",
			"duedate": "2022-06-12",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		task := NewControlTask(&mockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.UpdateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.UpdateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "kecantikan"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.UpdateTask())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.UpdateTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Success Delete Task", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		task := NewControlTask(&mockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.DeleteTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Delete Task", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.DeleteTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		task := NewControlTask(&errMockTask{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(task.DeleteTask())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

type mockTask struct{}

func (c *mockTask) CreateTask(newAdd entities.Task, userid uint) (entities.Task, error) {
	return entities.Task{UserID: 1, ProjekID: 1, Name: "List of guest", Duedate: "2022-05-23", Status: "completed"}, nil
}
func (c *mockTask) GetAllTask(Userid uint) ([]entities.Task, error) {
	return []entities.Task{{UserID: 1, ProjekID: 1, Name: "List of guest", Duedate: "2022-05-23", Status: "completed"}}, nil
}
func (c *mockTask) GetTaskID(id uint, userid uint) (entities.Task, error) {
	return entities.Task{UserID: 1, ProjekID: 1, Name: "List of guest", Duedate: "2022-05-23", Status: "completed"}, nil
}
func (c *mockTask) UpdateTask(id uint, updatedAddress entities.Task, userid uint) (entities.Task, error) {
	return entities.Task{UserID: 1, ProjekID: 1, Name: "List of guest", Duedate: "2022-05-23", Status: "completed"}, nil
}
func (c *mockTask) DeleteTask(id uint, Userid uint) error {
	return nil
}

type errMockTask struct {
}

// METHOD MOCK ERROR
func (e *errMockTask) CreateTask(newAdd entities.Task, userid uint) (entities.Task, error) {
	return entities.Task{}, errors.New("Access Database Error")
}

func (e *errMockTask) GetAllTask(Userid uint) ([]entities.Task, error) {
	return nil, errors.New("Access Database Error")
}

func (e *errMockTask) GetTaskID(id uint, userid uint) (entities.Task, error) {
	return entities.Task{}, errors.New("Access Database Error")
}

func (e *errMockTask) UpdateTask(id uint, updatedAddress entities.Task, userid uint) (entities.Task, error) {
	return entities.Task{}, errors.New("Access Database Error")
}

func (e *errMockTask) DeleteTask(id uint, Userid uint) error {
	return errors.New("Access Database Error")
}

package projek

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

func TestCreateProjek(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"userid":      1,
			"title":       "Live Concert Kahitna",
			"description": "live concert outdoor",
			"contributor": "astuti",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek")
		projekC := NewControlPro(&mockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.CreateProjek())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Projek", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"userid":      1,
			"title":       "Live Concert Kahitna",
			"description": "live concert outdoor",
			"contributor": "astuti",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.CreateProjek())(context)

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
		context.SetPath("/projek")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.CreateProjek())(context)

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

			"title":       "Live Concert Kahitna",
			"description": "live concert outdoor",
			"contributor": "astuti",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.CreateProjek())(context)

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

func TestGetAllProjek(t *testing.T) {
	t.Run("Success Get All Projek", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek")
		projekC := NewControlPro(&mockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.GetAllProjek())(context)

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
		context.SetPath("/projek")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.GetAllProjek())(context)

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

func TestGetProjekID(t *testing.T) {
	t.Run("Success Get Task By ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		projekC := NewControlPro(&mockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.GetProjekID())(context)
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
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.GetProjekID())(context)

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
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.GetProjekID())(context)

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

func TestUpdateProjek(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":       "Live Concert Kahitna",
			"description": "live concert and record international ft kahitna",
			"contributor": "indah",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/task/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		projekC := NewControlPro(&mockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.UpdateProjek())(context)

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
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.UpdateProjek())(context)

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
		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.UpdateProjek())(context)
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

		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.UpdateProjek())(context)

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

func TestDeleteProjek(t *testing.T) {
	t.Run("Success Delete Projek", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		projekC := NewControlPro(&mockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.DeleteProjek())(context)

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
	t.Run("Error Delete Projek", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")

		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.DeleteProjek())(context)

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
		context.SetPath("/projek/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")

		projekC := NewControlPro(&errMockProjek{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TODO")})(projekC.DeleteProjek())(context)

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

type mockProjek struct{}

func (c *mockProjek) CreateProjek(newAdd entities.Projek, userid uint) (entities.Projek, error) {
	return entities.Projek{UserID: 1, Title: "Live Concert Kahitna", Description: "live music outdoor", Date: "2022-07-1", Contributor: "astuti"}, nil
}
func (c *mockProjek) GetAllProjek(userid uint) ([]entities.Projek, error) {
	return []entities.Projek{{UserID: 1, Title: "Live Concert Kahitna", Description: "live music outdoor", Date: "2022-07-1", Contributor: "astuti"}}, nil
}
func (c *mockProjek) GetProjekID(id uint, userid uint) (entities.Projek, error) {
	return entities.Projek{UserID: 1, Title: "Live Concert Kahitna", Description: "live music outdoor", Date: "2022-07-1", Contributor: "astuti"}, nil
}
func (c *mockProjek) UpdateProjek(id uint, UpdatePro entities.Projek, userid uint) (entities.Projek, error) {
	return entities.Projek{UserID: 1, Title: "Live Concert Kahitna", Description: "live music outdoor", Date: "2022-07-1", Contributor: "astuti"}, nil
}
func (c *mockProjek) DeleteProjek(id uint, userid uint) error {
	return nil
}

type errMockProjek struct {
}

// METHOD MOCK ERROR
func (e *errMockProjek) CreateProjek(newAdd entities.Projek, userid uint) (entities.Projek, error) {
	return entities.Projek{}, errors.New("Access Database Error")
}

func (e *errMockProjek) GetAllProjek(userid uint) ([]entities.Projek, error) {
	return nil, errors.New("Access Database Error")
}

func (e *errMockProjek) GetProjekID(id uint, userid uint) (entities.Projek, error) {
	return entities.Projek{}, errors.New("Access Database Error")
}

func (e *errMockProjek) UpdateProjek(id uint, UpdatePro entities.Projek, userid uint) (entities.Projek, error) {
	return entities.Projek{}, errors.New("Access Database Error")
}

func (e *errMockProjek) DeleteProjek(id uint, userid uint) error {
	return errors.New("Access Database Error")
}

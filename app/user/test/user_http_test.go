package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var jwtKey = []byte("testKey")

func setupHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	})
	return token.SignedString(jwtKey)
}

func TestCreateUserHTTP(t *testing.T) {
	hClient := setupHTTPClient()

	reqBody, err := json.Marshal(map[string]string{
		"mobile":   "13811211501",
		"password": "password",
		"nickName": "nickname",
		"role":     "1",
	})
	assert.NoError(t, err)

	token, err := generateJWT()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "http://127.0.0.1:8002/v1/users", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // 使用生成的JWT

	resp, err := hClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	log.Printf("[http] CreateUser %+v\n", resp)
}

func TestUpdateUserHTTP(t *testing.T) {
	hClient := setupHTTPClient()
	reqBody, err := json.Marshal(map[string]string{
		"mobile":   "13811211501",
		"password": "newpassword",
		"nickName": "newnickname",
		"gender":   "female",
		"role":     "3",
	})
	assert.NoError(t, err)

	token, err := generateJWT()
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "http://127.0.0.1:8002/v1/users/2", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // 使用生成的JWT

	resp, err := hClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	log.Printf("[http] UpdateUser %+v\n", resp)
}

func TestGetUserHTTP(t *testing.T) {
	hClient := setupHTTPClient()

	token, err := generateJWT()
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "http://127.0.0.1:8002/v1/users/2", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token) // 使用生成的JWT
	req.Header.Set("Content-Type", "application/json")

	resp, err := hClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	log.Printf("[http] GetUser %+v\n", resp)
}

func TestListUserHTTP(t *testing.T) {
	hClient := setupHTTPClient()

	token, err := generateJWT()
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "http://127.0.0.1:8002/v1/users", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token) // 使用生成的JWT
	req.Header.Set("Content-Type", "application/json")

	resp, err := hClient.Do(req)
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	// 打印响应体中的用户列表数据
	log.Printf("[http] ListUser response body: %s\n", string(body))
}

package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь -проверяет, что данные из таблицы корректно извлекаются
	gotClient, err := selectClient(db, clientID)
	require.NoError(t, err)
	require.Equal(t, clientID, gotClient.ID)
	assert.NotEmpty(t, gotClient.Login)
	assert.NotEmpty(t, gotClient.Birthday)
	assert.NotEmpty(t, gotClient.Email)
	assert.NotEmpty(t, gotClient.FIO)

}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1

	// напиши тест здесь проверяет, что в случае отсутствия записи возвращается ошибка
	gotClient, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, gotClient.ID)
	assert.Empty(t, gotClient.Login)
	assert.Empty(t, gotClient.Birthday)
	assert.Empty(t, gotClient.Email)
	assert.Empty(t, gotClient.FIO)

}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь добавляет запись в таблицу, а затем проверяет, что она добавилась корректно
	clientID, err := insertClient(db, cl)
	require.Empty(t, err)
	require.NotEmpty(t, clientID)

	gotClient, err := selectClient(db, clientID)
	require.Empty(t, err)
	assert.Equal(t, clientID, gotClient.ID)
	assert.Equal(t, cl.Birthday, gotClient.Birthday)
	assert.Equal(t, cl.Email, gotClient.Email)
	assert.Equal(t, cl.FIO, gotClient.FIO)
	assert.Equal(t, cl.Login, gotClient.Login)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь проверяет удаление записи.
	clientID, err := insertClient(db, cl)
	require.Empty(t, err)
	require.NotEmpty(t, clientID)

	_, err = selectClient(db, clientID)
	require.Empty(t, err)

	err = deleteClient(db, clientID)
	require.Empty(t, err)

	_, err = selectClient(db, clientID)
	require.NotEmpty(t, err)
	require.Equal(t, sql.ErrNoRows, err)

}

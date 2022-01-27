package rediscache

import (
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdd_Success(t *testing.T) {

	//arrange
	client, mock := redismock.NewClientMock()

	key := "foo"
	value := "bar"
	duration := time.Duration(0)

	mock.ExpectSet(key, value, duration).SetVal("")

	helper := Helper{
		client,
	}

	//act
	err := helper.Add(key, value, duration)

	//assert
	assert.NoError(t, err)
}

func TestAdd_Fail(t *testing.T) {

	//arrange
	client, mock := redismock.NewClientMock()

	key := "foo"
	value := "bar"
	duration := time.Duration(0)
	errMessage := "FAIL"
	mock.ExpectSet(key, value, duration).SetErr(errors.New(errMessage))

	helper := Helper{
		client,
	}

	//act
	err := helper.Add(key, value, duration)

	//assert
	assert.Error(t, err)
	assert.Equal(t, errMessage, err.Error())
}

func TestGet_ExistingKey(t *testing.T) {
	//arrange
	client, mock := redismock.NewClientMock()

	key := "foo"
	value := "bar"

	mock.ExpectGet(key).SetVal(value)

	helper := Helper{
		client,
	}

	//act
	actual, err := helper.Get(key)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, value, actual)
}

func TestGet_NonExistingKey(t *testing.T) {

	//arrange
	client, mock := redismock.NewClientMock()

	key := "foo"

	mock.ExpectGet(key).RedisNil()

	helper := Helper{
		client,
	}

	//act
	actual, err := helper.Get(key)

	//assert
	assert.NoError(t, err)
	assert.Nil(t, actual)
}

func TestGet_Fail(t *testing.T) {

	//arrange
	client, mock := redismock.NewClientMock()

	key := "foo"
	errMessage := "FAIL"

	mock.ExpectGet(key).SetErr(errors.New(errMessage))

	helper := Helper{
		client,
	}

	//act
	actual, err := helper.Get(key)

	//assert
	assert.Error(t, err)
	assert.Nil(t, actual)
	assert.Equal(t, errMessage, err.Error())

}

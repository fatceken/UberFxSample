package rediscache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m MockObject) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	r := args.Get(0)
	v, _ := r.(*redis.StatusCmd)
	return v
}

func (m MockObject) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	r := args.Get(0)
	v, _ := r.(*redis.StringCmd)
	return v

}

func (m MockObject) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestAdd_Success(t *testing.T) {
	//arrange
	key := "foo"
	value := "bar"
	duration := time.Duration(0)

	mockObject := new(MockObject)

	mockObject.On("Set", context.Background(), key, value, duration).Return(redis.NewStatusResult("A", nil))

	helper := Helper{
		mockObject,
	}

	//act
	err := helper.Add(key, value, duration)

	//assert
	assert.NoError(t, err)
}

func TestAdd_Fail(t *testing.T) {
	//arrange
	key := "foo"
	value := "bar"
	duration := time.Duration(0)
	errMessage := "FAIL"

	mockObject := new(MockObject)

	mockObject.On("Set", context.Background(), key, value, duration).Return(redis.NewStatusResult("A", errors.New(errMessage)))

	helper := Helper{
		mockObject,
	}

	//act
	err := helper.Add(key, value, duration)

	//assert
	assert.Error(t, err)
	assert.Equal(t, errMessage, err.Error())
}

func TestGet_ExistingKey(t *testing.T) {
	//arrange
	key := "foo"
	value := "bar"

	mockObject := new(MockObject)

	mockObject.On("Get", context.Background(), key).Return(redis.NewStringResult(value, nil))

	helper := Helper{
		mockObject,
	}

	//act
	actual, err := helper.Get(key)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, value, actual)
}

func TestGet_NonExistingKey(t *testing.T) {
	//arrange
	key := "foo"

	mockObject := new(MockObject)

	mockObject.On("Get", context.Background(), key).Return(redis.NewStringResult("", redis.Nil))

	helper := Helper{
		mockObject,
	}

	//act
	actual, err := helper.Get(key)

	//assert
	assert.NoError(t, err)
	assert.Nil(t, actual)
}

func TestGet_Fail(t *testing.T) {
	//arrange
	key := "foo"
	errMessage := "FAIL"

	mockObject := new(MockObject)

	mockObject.On("Get", context.Background(), key).Return(redis.NewStringResult("", errors.New(errMessage)))

	helper := Helper{
		mockObject,
	}

	//act
	actual, err := helper.Get(key)

	//assert
	assert.Error(t, err)
	assert.Nil(t, actual)
	assert.Equal(t, errMessage, err.Error())
}

func Test_Close(t *testing.T) {
	//arrange

	mockObject := new(MockObject)
	mockObject.On("Close").Return(nil)

	helper := Helper{
		mockObject,
	}

	//act
	err := helper.Close()

	//assert
	assert.NoError(t, err)
}

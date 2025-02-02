package controller_test

import (
	"context"
	"encoding/json"
	"example/test-golang/controller"
	"example/test-golang/entity"
	"example/test-golang/entity/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"net/http"
	"net/http/httptest"
)

type ItemControllerSuite struct {
	suite.Suite
	itemService *ItemServiceMock
	underTest   controller.ItemController
}

type ItemServiceMock struct {
	mock.Mock
}

func (m *ItemServiceMock) GetAll(ctx context.Context) []entity.Item {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Item)
}

func (s *ItemControllerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	s.itemService = new(ItemServiceMock)
	s.underTest = controller.NewItemController(s.itemService)
}

func (s *ItemControllerSuite) TestGetAll() {
	assert := s.Assertions

	mockItems := getMockItems()
	expectedResponse := dto.GetAllResponse{Items: mockItems}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	s.itemService.On("GetAll", ctx).Return(mockItems)

	s.underTest.GetAll(ctx)

	var response dto.GetAllResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)

	s.itemService.AssertExpectations(s.T())
	assert.Equal(http.StatusOK, recorder.Code)
	assert.Equal(expectedResponse, response)
}

func getMockItems() []entity.Item {
	item1 := entity.Item{Name: "item1", Price: 100.00}
	item2 := entity.Item{Name: "item2", Price: 200.00}
	item3 := entity.Item{Name: "item3", Price: 300.00}

	return []entity.Item{item1, item2, item3}
}

func TestItemControllerSuite(t *testing.T) {
	suite.Run(t, new(ItemControllerSuite))
}

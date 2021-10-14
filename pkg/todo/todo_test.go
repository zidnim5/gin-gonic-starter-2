package todo

import (
	"starter/pkg/todo/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUnit(t *testing.T) {
	var a int = 2

	if a != 2 {
		panic("error")
	}

}

func TestMocking(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTodo := mocks.NewMockTodoRepo(mockCtrl)
	testTodo := &todoMock{mock: mockTodo}

	// Expect Do to be called once with 1 and "abc" as parameters, and return nil from the mocked call.
	mockTodo.EXPECT().FindId("abc").Return(nil).Times(1)
	testTodo.FindId()

	mockCtrl.Finish()
}

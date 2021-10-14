package todo

type todoMock struct {
	mock TodoRepo
}

func (tm *todoMock) FindId() {
	tm.mock.FindId("abc")
}

func (tm *todoMock) Insert() {
	tm.mock.Insert()
}

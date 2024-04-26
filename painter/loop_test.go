package painter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"testing"
	"time"
)

type MockOperation struct {
	mock.Mock
}

func (mockOperation *MockOperation) Do(t screen.Texture) bool {
	args := mockOperation.Called(t)
	return args.Bool(0)
}

type MockScreen struct {
	mock.Mock
}

func (_ *MockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	return nil, nil
}

func (_ *MockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	return nil, nil
}

func (mockScreen *MockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	args := mockScreen.Called(size)
	return args.Get(0).(screen.Texture), args.Error(1)
}

type MockTexture struct {
	mock.Mock
}

func (mockTexture *MockTexture) Release() {
	mockTexture.Called()
}

func (mockTexture *MockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	mockTexture.Called(dp, src, sr)
}

func (mockTexture *MockTexture) Bounds() image.Rectangle {
	args := mockTexture.Called()
	return args.Get(0).(image.Rectangle)
}

func (mockTexture *MockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	mockTexture.Called(dr, src, op)
}

func (mockTexture *MockTexture) Size() image.Point {
	args := mockTexture.Called()
	return args.Get(0).(image.Point)
}

type MockReceiver struct {
	mock.Mock
}

func (mockReceiver *MockReceiver) Update(t screen.Texture) {
	mockReceiver.Called(t)
}

func setupTestResources() (
	screenmock *MockScreen,
	receivermock *MockReceiver,
	texturemock *MockTexture,
	loop Loop) {

	screenmock = new(MockScreen)
	receivermock = new(MockReceiver)
	texturemock = new(MockTexture)
	texture := image.Pt(800, 800)

	screenmock.On("NewTexture", texture).Return(texturemock, nil)
	receivermock.On("Update", texturemock).Return()
	texturemock.On("Bounds").Return(image.Rectangle{})

	loop = Loop{
		Receiver: receivermock,
	}

	return
}

func TestLoop_Post_Operation_WasRun(t *testing.T) {
	screenMock, receiverMock, textureMock, loop := setupTestResources()

	operation := new(MockOperation)
	operation.On("Do", textureMock).Return(true)

	loop.Start(screenMock)
	assert.Empty(t, loop.mq.data)

	loop.Post(operation)
	time.Sleep(1000 * time.Millisecond)

	operation.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))

	assert.Empty(t, loop.mq.data)
}

func TestLoop_Post_Operation_WasNotRun(t *testing.T) {
	screenMock, receiverMock, textureMock, loop := setupTestResources()

	operation := new(MockOperation)
	operation.On("Do", textureMock).Return(false)

	loop.Start(screenMock)
	assert.Empty(t, loop.mq.data)

	loop.Post(operation)
	time.Sleep(1000 * time.Millisecond)

	operation.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertNotCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))

	assert.Empty(t, loop.mq.data)
}

func TestLoop_Post_Three_Operations_WereRun(t *testing.T) {
	screenMock, receiverMock, textureMock, loop := setupTestResources()

	operations := make([]*MockOperation, 3)
	for i := range operations {
		operations[i] = new(MockOperation)
		operations[i].On("Do", textureMock).Return(true)
	}

	loop.Start(screenMock)
	assert.Empty(t, loop.mq.data)

	for i := range operations {
		loop.Post(operations[i])
	}
	time.Sleep(1000 * time.Millisecond)

	for i := range operations {
		operations[i].AssertCalled(t, "Do", textureMock)
	}
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))

	assert.Empty(t, loop.mq.data)
}

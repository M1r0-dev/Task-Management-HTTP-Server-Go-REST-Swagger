package http

import (
	"http_server_arch/api/http/types"
	"http_server_arch/domain"
	"http_server_arch/usecases"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Object struct {
	service usecases.Object
}

func NewHandler(service usecases.Object) *Object {
	return &Object{service: service}
}

// Get Status Handler
// @Summary Получение статуса задачи по ID
// @Description Получить статус задачи по указанному ID.
// @Tags task
// @Accept json
// @Produce json
// @Param id query string true "ID задачи"
// @Success 200 {object} types.GetStatusHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Router /task/status [get]
func (s *Object) getStatusHandler(w http.ResponseWriter, r *http.Request) {
	request, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	task, err := s.service.Get(request.ID)
	types.ProcessError(w, err, &types.GetStatusHandlerResponse{
		Status: &task.Status})
}

// Get Result Handler
// @Summary Получение результата задачи по ID
// @Description Получить результат выполнения задачи по указанному ID.
// @Tags task
// @Accept json
// @Produce json
// @Param id query string true "ID задачи"
// @Success 200 {object} types.GetResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Router /task/result [get]
func (s *Object) getResultHandler(w http.ResponseWriter, r *http.Request) {
	request, err := types.CreateGetObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	task, err := s.service.Get(request.ID)
	types.ProcessError(w, err, &types.GetResultHandlerResponse{
		Result: &task.Result})
}

// Put Object Handler
// @Summary Создание или обновление задачи
// @Description Создать новую задачу с переданным кодом и компилятором. После создания возвращает ID и начальный статус.
// @Tags task
// @Accept json
// @Produce json
// @Param body body types.PutObjectHandlerRequest true "Данные задачи"
// @Success 200 {object} types.PutObjectHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Router /task [put]
func (s *Object) putHandler(w http.ResponseWriter, r *http.Request) {
	request, err := types.CreatePutObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	task := &domain.Task{
		ID:       id,
		Code:     request.Code,
		Compiler: request.Compiler,
		Status:   "in_progress",
		Result:   "",
	}

	err = s.service.Put(task)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}

	types.ProcessError(w, nil, &types.PutObjectHandlerResponse{ID: task.ID, Status: task.Status})
	go func(t *domain.Task) {
		time.Sleep(10 * time.Second)
		t.Result = "qww qqq eew qwe eee"
		t.Status = "ready"
		_ = s.service.Put(t)
	}(task)
}

// Post Object Handler
// @Summary Создание задачи
// @Description Создать новую задачу с переданным кодом и компилятором.
// @Tags task
// @Accept json
// @Produce json
// @Param body body types.PostObjectHandlerRequest true "Данные задачи"
// @Success 200 {object} types.PostObjectHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Router /task [post]
func (s *Object) postHandler(w http.ResponseWriter, r *http.Request) {
	request, err := types.CreatePostObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	task := &domain.Task{
		ID:       id,
		Code:     request.Code,
		Compiler: request.Compiler,
		Status:   "in_progress",
		Result:   "",
	}

	err = s.service.Post(task)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}

	types.ProcessError(w, nil, &types.PostObjectHandlerResponse{ID: task.ID, Status: task.Status})
	go func(t *domain.Task) {
		time.Sleep(10 * time.Second)
		t.Result = "qww qqq eew qwe eee"
		t.Status = "ready"
		_ = s.service.Put(t)
	}(task)
}

// Delete Object Handler
// @Summary Удаление задачи
// @Description Удалить задачу по указанному ID.
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "ID задачи"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Router /task/{id} [delete]
func (s *Object) deleteHandler(w http.ResponseWriter, r *http.Request) {
	request, err := types.CreateDeleteObjectHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = s.service.Delete(request.ID)
	types.ProcessError(w, err, nil)
}

// WithObjectHandlers регистрирует маршруты для работы с задачами
func (s *Object) WithObjectHandlers(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
		r.Post("/", s.postHandler)
		r.Put("/", s.putHandler)
		r.Get("/status", s.getStatusHandler)
		r.Get("/result", s.getResultHandler)
		r.Delete("/{id}", s.deleteHandler)
	})
}

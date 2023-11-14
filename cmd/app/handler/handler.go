package handler

import (
	"board/internal/controller"
	"board/internal/controller/req"
	page2 "board/internal/page"
	"board/internal/service/res"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	// 상세
	r.HandleFunc("/boards/{id:[0-9]+}/detail", h.getDetail).Methods(http.MethodGet)
	// query 로 최신순,boardType,writer 검색을 넣을거임
	r.HandleFunc("/boards/{cafeId:[0-9]+}", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/boards/{cafeId:[0-9]+}/{boardType:[0-9]+}", h.create).Methods(http.MethodPost)
	r.HandleFunc("/boards/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/boards/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InvalidCafeId       = "invalid cafe id"
	InvalidBoardType    = "invalid board type"
	InvalidBoardId      = "invalid board id"
	InternalServerError = "internal server error"
)

func (h Handler) getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}
	detail, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(detail)
	if err != nil {
		log.Println("getDetail json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}

	q := r.URL.Query()
	boardType, err := strconv.Atoi(q.Get("board-type"))
	if err != nil {
		boardType = 0
	}
	writer, err := strconv.Atoi(q.Get("writer"))
	if err != nil {
		writer = 0
	}
	page, err := strconv.Atoi(q.Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(q.Get("size"))
	if err != nil {
		size = 0
	}
	reqPage := page2.NewReqPage(page, size)

	list, total, err := h.c.GetList(r.Context(), cafeId, boardType, writer, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dto := res.ListTotalDto{
		Content: list,
		Total:   total,
	}
	data, err := json.Marshal(&dto)
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}

	boardType, err := strconv.Atoi(vars["boardType"])
	if err != nil {
		http.Error(w, InvalidBoardType, http.StatusBadRequest)
		return
	}

	var dto req.Create
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Create(r.Context(), cafeId, boardType, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}

	err = h.c.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}

	var dto req.Patch
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Patch(r.Context(), id, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

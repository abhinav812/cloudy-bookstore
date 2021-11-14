package app

import (
	"encoding/json"
	"fmt"
	"github.com/abhinav812/cloudy-bookstore/internal/model"
	"github.com/abhinav812/cloudy-bookstore/internal/repository"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJSONCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
)

//HandleListBooks - http route to list all the books. Returns result in form of model.BookDtos
func (s *Server) HandleListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := repository.ListBooks(s.db)
	if err != nil {
		s.Logger().Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, `{"error": "%v""}`, appErrDataAccessFailure)
		return
	}

	if books == nil {
		_, _ = fmt.Fprintf(w, "[]")
		return
	}

	dtos := books.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		s.Logger().Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, `{"error": "%v""}`, appErrJSONCreationFailure)
		return
	}

	/*w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("[]"))*/
}

//HandleCreateBook - http route to create new book, with supplied model.Book json
func (s *Server) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	book := &model.Book{}

	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	resp, err := repository.CreateBook(s.db, book)
	if err != nil {
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, `{"error": "%v"}`, appErrDataCreationFailure)
		return
	}

	s.Logger().Info().Msgf("New book created. ID: %d", resp.ID)
	w.WriteHeader(http.StatusCreated)
}

//HandleReadBook - http route to handle book by id request. Returns result in form of model.BookDto json
func (s *Server) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil {
		s.logger.Info().Msgf("can not parse ID: %v", id)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if id == 0 {
		s.logger.Info().Msg("id value has to be greater that 0")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	book, err := repository.ReadBook(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	dto := book.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJSONCreationFailure)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//HandleUpdateBook - http route to handle update book based on id and supplied model.Book json payload
func (s *Server) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil {
		s.logger.Info().Msgf("can not parse ID: %v", id)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if id == 0 {
		s.logger.Info().Msg("id value has to be greater that 0")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	book := &model.Book{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		return
	}

	book.ID = uint(id)
	if err := repository.UpdateBook(s.db, book); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataUpdateFailure)
		return
	}

	s.logger.Info().Msgf("Book updated: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

//HandleDeleteBook - http request to handle delete book request by bookId
func (s *Server) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil {
		s.logger.Info().Msgf("can not parse ID: %v", id)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if id == 0 {
		s.logger.Info().Msg("id value has to be greater that 0")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := repository.DeleteBook(s.db, uint(id)); err != nil {
		s.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}
	s.logger.Info().Msgf("Book deleted: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

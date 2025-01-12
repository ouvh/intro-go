package httpHandlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"um6p.ma/project-04/dao"
	"um6p.ma/project-04/httpHandlers/httpError"
	"um6p.ma/project-04/httpHandlers/httpJsonParser"
	"um6p.ma/project-04/httpHandlers/httpStructVerifier"
	"um6p.ma/project-04/models/author"
	"um6p.ma/project-04/models/book"
	"um6p.ma/project-04/models/customer"
	"um6p.ma/project-04/models/order"
	"um6p.ma/project-04/models/salesReport"
	service "um6p.ma/project-04/services"
	"um6p.ma/project-04/store/searchCriteria"
)

func DispatcherWrapper(requestHandler func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		clientContext := r.Context()
		requestContext, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		requestChannel := make(chan bool)
		go func() {
			requestHandler(w, r)
			requestChannel <- true
		}()

		select {
		case <-clientContext.Done():
			log.Println("Connection Lost")
			return
		case <-requestChannel:
			log.Println("Request Done with Success")

		case <-requestContext.Done():
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: "request timeout"}, w, http.StatusRequestTimeout)
			log.Println("Request timeout")
		}
	}

}

func ServiceInjector(service *service.Service, handler func(w http.ResponseWriter, r *http.Request, service *service.Service)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, service)
	}

}

func BookHandler(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		params := make(map[string]any, 0)
		searchCriteria := searchCriteria.SearchCriteria{}
		for key, value := range r.URL.Query() {
			if len(value) == 2 {
				params[key] = value[0] + "+" + value[1]
			} else {
				params[key] = value[0]
			}
		}
		e := httpStructVerifier.ValidateSearchCriteria(book.Fields, params, make([]string, 0))
		if e != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: e.Error()}, w, http.StatusBadRequest)
			log.Println(e.Error())
			return
		}
		searchCriteria.Parameters = params

		result, err := service.BookService.SearchBook(searchCriteria)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		} else {
			httpJsonParser.SetJson(result, w, http.StatusOK)
			log.Println("Search Request Accepted")
			return

		}

	} else if r.Method == http.MethodPost {
		requestBody := dao.BookDAO{}
		err := httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.BookService.CreateBook(requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Book Created", result)
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}
}

func BookHandlerWithPath(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		id, err := httpStructVerifier.ValidatePathId(r, "/books/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.BookService.GetBook(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("request Ok")
		return

	} else if r.Method == http.MethodPut {
		id, err := httpStructVerifier.ValidatePathId(r, "/books/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		requestBody := dao.BookDAO{}
		err = httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.BookService.UpdateBook(int(id), requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Book Updated", result)

	} else if r.Method == http.MethodDelete {
		id, err := httpStructVerifier.ValidatePathId(r, "/books/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = service.BookService.DeleteBook(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson("Book Deleted", w, http.StatusOK)
		log.Println("request Ok")
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
	}

}

func AuthorHandler(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		params := make(map[string]any, 0)
		searchCriteria := searchCriteria.SearchCriteria{}
		for key, value := range r.URL.Query() {
			if len(value) == 2 {
				params[key] = value[0] + "+" + value[1]
			} else {
				params[key] = value[0]
			}
		}
		e := httpStructVerifier.ValidateSearchCriteria(author.Fields, params, make([]string, 0))
		if e != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: e.Error()}, w, http.StatusBadRequest)
			log.Println(e.Error())
			return
		}
		searchCriteria.Parameters = params

		result, err := service.AuthorService.SearchAuthor(searchCriteria)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		} else {
			httpJsonParser.SetJson(result, w, http.StatusOK)
			log.Println("Search Request Accepted")
			return

		}

	} else if r.Method == http.MethodPost {
		requestBody := dao.AuthorDAO{}
		err := httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.AuthorService.CreateAuthor(requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Author Created", result)
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}
}

func AuthorHandlerWithPath(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		id, err := httpStructVerifier.ValidatePathId(r, "/authors/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.AuthorService.GetAuthor(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("request Ok")
		return

	} else if r.Method == http.MethodPut {
		id, err := httpStructVerifier.ValidatePathId(r, "/authors/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		requestBody := dao.AuthorDAO{}
		err = httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.AuthorService.UpdateAuthor(int(id), requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Author Updated", result)

	} else if r.Method == http.MethodDelete {
		id, err := httpStructVerifier.ValidatePathId(r, "/authors/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = service.AuthorService.DeleteAuthor(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson("Author Deleted", w, http.StatusOK)
		log.Println("request Ok")
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
	}

}

func CustomerHandler(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		params := make(map[string]any, 0)
		searchCriteria := searchCriteria.SearchCriteria{}
		for key, value := range r.URL.Query() {
			if len(value) == 2 {
				params[key] = value[0] + "+" + value[1]
			} else {
				params[key] = value[0]
			}
		}
		e := httpStructVerifier.ValidateSearchCriteria(customer.Fields, params, make([]string, 0))
		if e != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: e.Error()}, w, http.StatusBadRequest)
			log.Println(e.Error())
			return
		}
		searchCriteria.Parameters = params

		result, err := service.CustomerService.SearchCustomer(searchCriteria)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		} else {
			httpJsonParser.SetJson(result, w, http.StatusOK)
			log.Println("Search Request Accepted")
			return

		}

	} else if r.Method == http.MethodPost {
		requestBody := dao.CustomerDAO{}
		err := httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.CustomerService.CreateCustomer(requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Customer Created", result)
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}
}

func CustomerHandlerWithPath(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		id, err := httpStructVerifier.ValidatePathId(r, "/customers/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.CustomerService.GetCustomer(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("request Ok")
		return

	} else if r.Method == http.MethodPut {
		id, err := httpStructVerifier.ValidatePathId(r, "/customers/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		requestBody := dao.CustomerDAO{}
		err = httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.CustomerService.UpdateCustomer(int(id), requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Customer Updated", result)

	} else if r.Method == http.MethodDelete {
		id, err := httpStructVerifier.ValidatePathId(r, "/customers/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = service.CustomerService.DeleteCustomer(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson("Customer Deleted", w, http.StatusOK)
		log.Println("request Ok")
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
	}

}

func OrderHandler(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		params := make(map[string]any, 0)
		searchCriteria := searchCriteria.SearchCriteria{}
		for key, value := range r.URL.Query() {
			if len(value) == 2 {
				params[key] = value[0] + "+" + value[1]
			} else {
				params[key] = value[0]
			}
		}

		e := httpStructVerifier.ValidateSearchCriteria(order.Fields, params, make([]string, 0))
		if e != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: e.Error()}, w, http.StatusBadRequest)
			log.Println(e.Error())
			return
		}
		searchCriteria.Parameters = params

		result, err := service.OrderService.SearchOrder(searchCriteria)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		} else {
			httpJsonParser.SetJson(result, w, http.StatusOK)
			log.Println("Search Request Accepted")
			return

		}

	} else if r.Method == http.MethodPost {
		requestBody := dao.OrderDAO{}
		err := httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.OrderService.CreateOrder(requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("Order Created", result)
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}
}

func OrderHandlerWithPath(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		id, err := httpStructVerifier.ValidatePathId(r, "/orders/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.OrderService.GetOrder(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("request Ok")
		return

	} else if r.Method == http.MethodPut {
		id, err := httpStructVerifier.ValidatePathId(r, "/orders/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		requestBody := dao.OrderDAO{}
		err = httpJsonParser.LoadJson(r, &requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = httpStructVerifier.ValidateStruct(requestBody)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		result, err := service.OrderService.UpdateOrder(int(id), requestBody)

		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		httpJsonParser.SetJson(result, w, http.StatusOK)
		log.Println("order Updated", result)

	} else if r.Method == http.MethodDelete {
		id, err := httpStructVerifier.ValidatePathId(r, "/orders/")
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = service.OrderService.DeleteOrder(int(id))
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusNotFound)
			log.Println(err.Error())
			return
		}
		httpJsonParser.SetJson("Order Deleted", w, http.StatusOK)
		log.Println("request Ok")
		return

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
	}

}

func SalesReportHandler(w http.ResponseWriter, r *http.Request, service *service.Service) {
	if r.Method == http.MethodGet {
		params := make(map[string]any, 0)
		searchCriteria := searchCriteria.SearchCriteria{}
		for key, value := range r.URL.Query() {
			if len(value) == 2 {
				params[key] = value[0] + "+" + value[1]
			} else {
				params[key] = value[0]
			}
		}
		e := httpStructVerifier.ValidateSearchCriteria(salesReport.Fields, params, make([]string, 0))
		if e != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: e.Error()}, w, http.StatusBadRequest)
			log.Println(e.Error())
			return
		}
		searchCriteria.Parameters = params

		result, err := service.SalesReportService.SearchSalesReport(searchCriteria)
		if err != nil {
			httpJsonParser.SetJson(httpError.ErrorResponse{Error: err.Error()}, w, http.StatusBadRequest)
			log.Println(err.Error())
			return
		} else {
			httpJsonParser.SetJson(result, w, http.StatusOK)
			log.Println("Search Request Accepted")
			return

		}

	} else {
		httpJsonParser.SetJson(httpError.ErrorResponse{Error: "Invalid request method"}, w, http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}
}

type Server struct {
	Port string
}

func (s *Server) StartServer(service *service.Service) error {
	http.HandleFunc("/books", DispatcherWrapper(ServiceInjector(service, BookHandler)))
	http.HandleFunc("/books/", DispatcherWrapper(ServiceInjector(service, BookHandlerWithPath)))

	http.HandleFunc("/authors", DispatcherWrapper(ServiceInjector(service, AuthorHandler)))
	http.HandleFunc("/authors/", DispatcherWrapper(ServiceInjector(service, AuthorHandlerWithPath)))

	http.HandleFunc("/customers", DispatcherWrapper(ServiceInjector(service, CustomerHandler)))
	http.HandleFunc("/customers/", DispatcherWrapper(ServiceInjector(service, CustomerHandlerWithPath)))

	http.HandleFunc("/orders", DispatcherWrapper(ServiceInjector(service, OrderHandler)))
	http.HandleFunc("/orders/", DispatcherWrapper(ServiceInjector(service, OrderHandlerWithPath)))

	http.HandleFunc("/salesReport", DispatcherWrapper(ServiceInjector(service, SalesReportHandler)))

	err := http.ListenAndServe(":"+s.Port, nil)

	if err != nil {
		return fmt.Errorf("error serving : %w", err)
	}
	return nil

}

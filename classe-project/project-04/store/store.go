package store

import (
	"time"

	"um6p.ma/project-04/store/authorStore"
	"um6p.ma/project-04/store/bookStore"
	"um6p.ma/project-04/store/customerStore"
	"um6p.ma/project-04/store/orderStore"
	"um6p.ma/project-04/store/salesReportStore"
	"um6p.ma/project-04/store/storeJsonParser"
)

type Store struct {
	BookStore        bookStore.BookStore
	AuthorStore      authorStore.AuthorStore
	CustomerStore    customerStore.CustomerStore
	OrderStore       orderStore.OrderStore
	SalesReportStore salesReportStore.SalesReportStore
	Filepath         string
	Schedule         time.Duration
}

func (s *Store) Load() error {
	err := storeJsonParser.LoadJson(&s, s.Filepath)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Save() error {
	err := storeJsonParser.SaveJson(s, s.Filepath)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) SnapShot() error {

	s.BookStore.Lock()
	s.AuthorStore.Lock()
	s.CustomerStore.Lock()
	s.OrderStore.Lock()
	s.SalesReportStore.Lock()
	defer s.BookStore.Unlock()
	defer s.AuthorStore.Unlock()
	defer s.CustomerStore.Unlock()
	defer s.OrderStore.Unlock()
	defer s.SalesReportStore.Unlock()

	err := storeJsonParser.SaveJson(s, s.Filepath)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) StartSchedule() {
	go func() {
		for {
			time.Sleep(s.Schedule)
			s.SnapShot()
		}

	}()
}

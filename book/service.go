package book

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, dto CreateBookDto) (*Book, error)
	GetByID(ctx context.Context, id int) (*Book, error)
	Update(ctx context.Context, id int, dto UpdateBookDto) (*Book, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter BookFilter) ([]Book, error)
}

type BookService struct {
	repo Repository
}

func NewService(repo Repository) *BookService {
	return &BookService{repo}
}

func (s *BookService) Create(ctx context.Context, dto CreateBookDto) (*Book, error) {

	if err := dto.ValidateForCreate(); err != nil {
		return nil, err
	}

	book := &Book{
		Name:       dto.Name,
		Author:     dto.Author,
		TotalPages: dto.TotalPages,
		Read:       false,
	}

	if err := s.repo.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) Update(ctx context.Context, id int, dto UpdateBookDto) (*Book, error) {
	if err := dto.ValidateForUpdate(); err != nil {
		return nil, err
	}

	book, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		book.Name = *dto.Name
	}

	if dto.Author != nil {
		book.Author = *dto.Author
	}

	if dto.TotalPages != nil {
		book.TotalPages = *dto.TotalPages
	}

	if dto.Read != nil {
		if *dto.Read {
			book.ReadBook()
		} else {
			book.UnreadBook()
		}
	}

	err2 := s.repo.Update(ctx, book)

	return book, err2
}

func (s *BookService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *BookService) GetByID(ctx context.Context, id int) (*Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BookService) List(ctx context.Context, filter BookFilter) ([]Book, error) {
	return s.repo.List(ctx, filter)
}

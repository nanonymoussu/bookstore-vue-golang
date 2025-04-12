package handler

import (
	"context"

	"github.com/nanonymoussu/bookstore/internal/service"
	pb "github.com/nanonymoussu/bookstore/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookstoreServer struct {
	pb.UnimplementedBookstoreServiceServer
	bookService service.BookService
}

func NewBookstoreServer(bookService service.BookService) *BookstoreServer {
	return &BookstoreServer{bookService: bookService}
}

func (s *BookstoreServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	page := int(req.Page)
	pageSize := int(req.PageSize)

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	books, total, err := s.bookService.ListBooks(page, pageSize)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list books: %v", err)
	}

	pbBooks := make([]*pb.Book, len(books))
	for i, book := range books {
		pbBooks[i] = &pb.Book{
			Id:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Isbn:      book.ISBN,
			Price:     book.Price,
			Stock:     book.Stock,
			CreatedAt: timestamppb.New(book.CreatedAt),
			UpdatedAt: timestamppb.New(book.UpdatedAt),
		}
	}

	return &pb.ListBooksResponse{
		Books:      pbBooks,
		TotalCount: int32(total),
	}, nil
}

func (s *BookstoreServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	book, err := s.bookService.GetBook(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "book not found: %v", err)
	}

	return &pb.Book{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Isbn:      book.ISBN,
		Price:     book.Price,
		Stock:     book.Stock,
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
	}, nil
}

func (s *BookstoreServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
	book, err := s.bookService.CreateBook(req.Title, req.Author, req.Isbn, req.Price, req.Stock)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create book: %v", err)
	}

	return &pb.Book{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Isbn:      book.ISBN,
		Price:     book.Price,
		Stock:     book.Stock,
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
	}, nil
}

func (s *BookstoreServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.Book, error) {
	book, err := s.bookService.UpdateBook(req.Id, req.Title, req.Author, req.Isbn, req.Price, req.Stock)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update book: %v", err)
	}

	return &pb.Book{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Isbn:      book.ISBN,
		Price:     book.Price,
		Stock:     book.Stock,
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
	}, nil
}

func (s *BookstoreServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	err := s.bookService.DeleteBook(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete book: %v", err)
	}

	return &emptypb.Empty{}, nil
}

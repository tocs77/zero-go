package grpcapi

import (
	"context"
	"database/sql"
	"log"
	"maillist/src/mdb"
	pb "maillist/src/proto"
	"net"
	"time"

	"google.golang.org/grpc"
)

type MailServer struct {
	pb.UnimplementedMailingListServiceServer
	db *sql.DB
}

func pbEntryToMdbEntry(entry *pb.EmailEntry) *mdb.EmailEntry {
	t := time.Unix(entry.ConfirmedAt, 0)
	return &mdb.EmailEntry{
		Id:          entry.Id,
		Email:       entry.Email,
		ConfirmedAt: &t,
		OptOut:      entry.OptOut,
	}
}

func mdbEntryToPbEntry(entry *mdb.EmailEntry) *pb.EmailEntry {
	t := entry.ConfirmedAt.Unix()
	return &pb.EmailEntry{
		Id:          entry.Id,
		Email:       entry.Email,
		ConfirmedAt: t,
		OptOut:      entry.OptOut,
	}
}

func emailResponse(db *sql.DB, email string) (*pb.EmailResponse, error) {
	entry, err := mdb.GetEmail(db, email)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return &pb.EmailResponse{}, nil
	}

	return &pb.EmailResponse{Email: mdbEntryToPbEntry(entry)}, nil

}

func (s *MailServer) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC get email %v\n", req)
	return emailResponse(s.db, req.EmailAddr)
}

func (s *MailServer) GetEmailBatch(ctx context.Context, req *pb.GetEmailBatchRequest) (*pb.GetEmailBatchResponse, error) {
	log.Printf("gRPC get email batch %v\n", req)
	params := mdb.GetEmailBatchQueryParams{
		Count: int(req.Count),
		Page:  int(req.Page),
	}
	mdbEmails, err := mdb.GetEmailBatch(s.db, &params)
	if err != nil {
		return nil, err
	}

	pbEmails := make([]*pb.EmailEntry, 0, len(mdbEmails))
	for _, entry := range mdbEmails {
		pbEmails = append(pbEmails, mdbEntryToPbEntry(entry))
	}
	return &pb.GetEmailBatchResponse{EmailEntries: pbEmails}, nil
}

func (s *MailServer) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC create email %v\n", req)
	if err := mdb.CreateEmail(s.db, req.EmailAddr); err != nil {
		return nil, err
	}
	return emailResponse(s.db, req.EmailAddr)
}

func (s *MailServer) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC update email %v\n", req)
	entry := pbEntryToMdbEntry(req.EmailEntry)
	if err := mdb.UpdateEmail(s.db, entry); err != nil {
		return nil, err
	}
	return emailResponse(s.db, entry.Email)
}

func (s *MailServer) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC delete email %v\n", req)
	if err := mdb.DeleteEmail(s.db, req.EmailAddr); err != nil {
		return nil, err
	}
	return emailResponse(s.db, req.EmailAddr)
}

func Serve(db *sql.DB, bind string) {
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("gRPC failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mailServer := MailServer{db: db}

	pb.RegisterMailingListServiceServer(grpcServer, &mailServer)
	log.Printf("gRPC server started on %s\n", bind)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}

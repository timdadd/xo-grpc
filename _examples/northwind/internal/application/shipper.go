package application

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	models "northwind/internal/models"
	pb "northwind/proto/shipper"
	typespb "northwind/proto/typespb"
)

type ShipperService struct {
	pb.UnimplementedShipperServer
	db     *sql.DB
	logger *zap.Logger
}

func NewShipperService(logger *zap.Logger, db *sql.DB) *ShipperService {
	return &ShipperService{logger: logger, db: db}
}

func (s *ShipperService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.ShipperByShipperID(ctx, s.db, int16(req.ShipperID))
	if err != nil {
		return
	}

	err = m.Delete(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *ShipperService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Shipper
	m.CompanyName = req.GetCompanyName()
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	m.ShipperID = int16(req.GetShipperID())

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.ShipperID))

	return
}

func (s *ShipperService) ShipperByShipperID(ctx context.Context, req *pb.ShipperByShipperIDRequest) (res *typespb.Shipper, err error) {

	shipperID := int16(req.GetShipperID())

	result, err := models.ShipperByShipperID(ctx, s.db, shipperID)
	if err != nil {
		return
	}

	res = new(typespb.Shipper)
	res.ShipperID = int32(result.ShipperID)
	res.CompanyName = result.CompanyName
	if result.Phone.Valid {
		res.Phone = wrapperspb.String(result.Phone.String)
	}

	return
}

func (s *ShipperService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.ShipperByShipperID(ctx, s.db, int16(req.ShipperID))
	if err != nil {
		return
	}
	m.CompanyName = req.GetCompanyName()
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	m.ShipperID = int16(req.GetShipperID())

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *ShipperService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Shipper
	m.CompanyName = req.GetCompanyName()
	if v := req.GetPhone(); v != nil {
		m.Phone = sql.NullString{Valid: true, String: v.Value}
	}
	m.ShipperID = int16(req.GetShipperID())

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

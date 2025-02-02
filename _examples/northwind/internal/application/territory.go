package application

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	models "northwind/internal/models"
	pb "northwind/proto/territory"
	typespb "northwind/proto/typespb"
)

type TerritoryService struct {
	pb.UnimplementedTerritoryServer
	db     *sql.DB
	logger *zap.Logger
}

func NewTerritoryService(logger *zap.Logger, db *sql.DB) *TerritoryService {
	return &TerritoryService{logger: logger, db: db}
}

func (s *TerritoryService) Delete(ctx context.Context, req *pb.DeleteRequest) (res *emptypb.Empty, err error) {
	m, err := models.TerritoryByTerritoryID(ctx, s.db, req.TerritoryID)
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

func (s *TerritoryService) Insert(ctx context.Context, req *pb.InsertRequest) (res *emptypb.Empty, err error) {
	var m models.Territory
	m.RegionID = int16(req.GetRegionID())
	m.TerritoryDescription = req.GetTerritoryDescription()
	m.TerritoryID = req.GetTerritoryID()

	err = m.Insert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	err = sendResourceLocation(ctx, fmt.Sprintf("/%v", m.TerritoryID))

	return
}

func (s *TerritoryService) Region(ctx context.Context, req *pb.RegionRequest) (res *typespb.Region, err error) {
	m, err := models.TerritoryByTerritoryID(ctx, s.db, req.TerritoryID)
	if err != nil {
		return
	}

	result, err := m.Region(ctx, s.db)
	if err != nil {
		return
	}

	res = new(typespb.Region)
	res.RegionID = int32(result.RegionID)
	res.RegionDescription = result.RegionDescription

	return
}

func (s *TerritoryService) TerritoryByTerritoryID(ctx context.Context, req *pb.TerritoryByTerritoryIDRequest) (res *typespb.Territory, err error) {

	territoryID := req.GetTerritoryID()

	result, err := models.TerritoryByTerritoryID(ctx, s.db, territoryID)
	if err != nil {
		return
	}

	res = new(typespb.Territory)
	res.TerritoryID = result.TerritoryID
	res.TerritoryDescription = result.TerritoryDescription
	res.RegionID = int32(result.RegionID)

	return
}

func (s *TerritoryService) Update(ctx context.Context, req *pb.UpdateRequest) (res *emptypb.Empty, err error) {
	m, err := models.TerritoryByTerritoryID(ctx, s.db, req.TerritoryID)
	if err != nil {
		return
	}
	m.RegionID = int16(req.GetRegionID())
	m.TerritoryDescription = req.GetTerritoryDescription()
	m.TerritoryID = req.GetTerritoryID()

	err = m.Update(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

func (s *TerritoryService) Upsert(ctx context.Context, req *pb.UpsertRequest) (res *emptypb.Empty, err error) {
	var m models.Territory
	m.RegionID = int16(req.GetRegionID())
	m.TerritoryDescription = req.GetTerritoryDescription()
	m.TerritoryID = req.GetTerritoryID()

	err = m.Upsert(ctx, s.db)
	if err != nil {
		return
	}

	res = new(emptypb.Empty)

	return
}

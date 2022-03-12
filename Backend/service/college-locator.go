package service

import (
	"CollegeLocator/database"
	"CollegeLocator/pb"
	"CollegeLocator/structure"
	"context"
	"fmt"
)

type DataServer struct {
	postgresDb *database.PostgresDatabase
}

func NewDataServer() *DataServer {
	return &DataServer{
		postgresDb: database.InitPostgresDataBase(),
	}
}

func (dataServer *DataServer) InitApp(ctx context.Context, request *pb.InitRequest) (*pb.InitResponse, error) {
	rangeData, err := dataServer.postgresDb.GetFeesAndCutoff(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.InitResponse{
		MaxRangeFees: rangeData.MAXFees,
		MinRangeFees: rangeData.MINFees,
		MaxRangeCutoff: rangeData.MAXCutoff,
		MinRangeCutoff: rangeData.MINCutoff,
	}, nil
}

func (dataServer *DataServer) Search(request *pb.SearchRequest, stream pb.CollegeLocatorService_SearchServer) error {
	err := dataServer.postgresDb.SearchCollege(stream.Context(), request,
		func(data *structure.CollegeData) error {
			fmt.Println(data)
			err := stream.Send(&pb.SearchResponse{
				Image:     data.Images[0],
				Name:      data.Name,
				Address:   data.Address,
				CollegeId: data.CollegeId,
				Location: &pb.Location{
					Longitude: data.Longitude,
					Latitude:  data.Latitude,
				},
			})
			return err
		},
	)
	return err
}

func (dataServer *DataServer) Details(ctx context.Context, request *pb.DetailsRequest) (*pb.DetailsResponse, error) {
	data, err := dataServer.postgresDb.GetCollegeById(ctx, request.GetCollegeId())
	if data == nil {
		return nil, fmt.Errorf("No Value Found ..!!")
	}
	return &pb.DetailsResponse{
		CollegeId: data.CollegeId,
		Name:      data.Name,
		Address:   data.Address,
		Images:    data.Images,
		Email:     data.Email,
		ContactNo: data.ContactNo,
		Website:   data.Website,
		VrImage:   data.VrImages,
		Location: &pb.Location{
			Longitude: data.Longitude,
			Latitude:  data.Latitude,
		},
		Hostel:        data.Hostel,
		Cutoff:        data.Cutoff,
		Fees:          data.Fees,
		InstituteType: data.InstituteType,
		Details:       data.Details,
		Courses:       data.Courses,
		Deemed:        data.Deemed,
	}, err
}

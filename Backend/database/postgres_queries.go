package database

import (
	"CollegeLocator/pb"
	"CollegeLocator/structure"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

func (postgresDb *PostgresDatabase) CreateCollege(ctx context.Context, collegeData *structure.CollegeData) error {
	postgresDb.mutex.Lock()
	defer postgresDb.mutex.Unlock()

	tx := postgresDb.Db.MustBegin()

	fmt.Println("Executing Query Now")
	images, err := collegeData.Images.Value()
	if err != nil {
		return err
	}
	vrImages, err := collegeData.VrImages.Value()
	if err != nil {
		return err
	}
	courses, err := collegeData.Courses.Value()
	if err != nil {
		return err
	}
	rows := tx.MustExecContext(ctx,
		"insert into college_data (college_id, name, address, images, email, contact_no, website, vr_image, longitude, latitude, hostel, cutoff, fees, institute_type, details, courses, state, deemed, course_type, document_tokens) "+
			"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, to_tsvector($20))",
		collegeData.CollegeId, collegeData.Name, collegeData.Address, images, collegeData.Email, collegeData.ContactNo, collegeData.Website, vrImages, collegeData.Longitude, collegeData.Latitude, collegeData.Hostel, collegeData.Cutoff, collegeData.Fees, collegeData.InstituteType, collegeData.Details, courses, collegeData.State, collegeData.Deemed, collegeData.CourseType, collegeData.Name+" "+collegeData.CourseType+" "+collegeData.Details+" "+collegeData.Address+" "+collegeData.State)

	affected, err := rows.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "Unable To Get Affected Rows", err)
	}

	if affected <= 0 {
		return status.Errorf(codes.Unknown, "No Rows Get Affected")
	}

	err = tx.Commit()
	if err != nil {
		return status.Errorf(codes.Internal, "Unable To Commit", err)
	}

	fmt.Print("Created")
	return nil
}

func (postgresDb *PostgresDatabase) SearchCollege(ctx context.Context, searchRequest *pb.SearchRequest, send func(data *structure.CollegeData) error) error {
	postgresDb.mutex.Lock()
	defer postgresDb.mutex.Unlock()

	filters := ""

	var rows *sqlx.Rows
	var err error

	if searchRequest.GetHostelIsNull() == true {
		filters += "and hostel=" + strconv.FormatBool(searchRequest.GetHostel()) + " "
	}
	if searchRequest.GetCutoffNull() == true {
		filters += " and cutoff between 0" + string(searchRequest.GetCutoff()) + " "
	}
	if searchRequest.GetFeesNull() == true {
		filters += " and fees between 0" + string(searchRequest.GetFees()) + " "
	}
	if searchRequest.GetInstituteTypeNull() == true {
		filters += " and institute_type=" + strconv.FormatBool(searchRequest.GetInstituteType()) + " "
	}
	if searchRequest.GetStateNull() == true {
		filters += " and state=" + searchRequest.GetState() + " "
	}
	if searchRequest.GetCourseTypeNull() == true {
		filters += " and course_type=" + searchRequest.GetCourseType() + " "
	}
	if searchRequest.GetDeemedNull() == true {
		filters += " and deemed=" + strconv.FormatBool(searchRequest.GetDeemed()) + " "
	}
	if searchRequest.GetSearchQueryIsNull() == true {
		filteredQuery := RemoveSpecialChars(searchRequest.GetSearchQuery())
		searchString := ""
		for _, i := range strings.Fields(filteredQuery) {
			searchString += "& " + i + " "
		}
		searchString = RemoveFirstChar(searchString)
		filters += " and document_tokens @@ to_tsquery('" + searchString + "') "
	}

	var distance int = 0
	switch searchRequest.GetDistance() {
	case pb.Distance_NEARBY:
		distance = 1000
		break
	case pb.Distance_MID_RANGE:
		distance = 10000
		break
	case pb.Distance_LONG_RANGE:
		distance = 20000
		break
	}

	if searchRequest.Location != nil || searchRequest.GetLocationNull() == true {
		query := fmt.Sprintf("SELECT *, point(%G, %G) <@> point(longitude, latitude)::point as distance FROM college_data WHERE (point(%G, %G) <@> point(longitude, latitude)) < %d and 1=1 %s ORDER BY distance LIMIT 20;", searchRequest.Location.Longitude, searchRequest.Location.Latitude, searchRequest.Location.Longitude, searchRequest.Location.Latitude, distance, filters)
		fmt.Println(query)
		rows, err = postgresDb.Db.QueryxContext(ctx, query)
	} else {
		query := fmt.Sprintf("SELECT * from college_data where 1=1 %s LIMIT 20;", filters)
		fmt.Println(query)
		rows, err = postgresDb.Db.QueryxContext(ctx, query)
	}

	if err != nil {
		panic(err)
		return err
	}

	var data structure.CollegeData
	for rows.Next() {
		err := rows.StructScan(&data)
		if err != nil {
			panic(err)
			//			log.Fatalln(err)
		}

		err = send(&data)
		if err != nil {
			panic(err)
			//			log.Fatalln(err)
		}
	}
	return nil

}

func (postgresDb *PostgresDatabase) GetCollegeById(ctx context.Context, collegeId string) (*structure.CollegeData, error) {
	postgresDb.mutex.Lock()
	defer postgresDb.mutex.Unlock()

	data := structure.CollegeData{}
	err := postgresDb.Db.GetContext(ctx, &data, "select * from college_data where college_id=$1", collegeId)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (postgresDb *PostgresDatabase) GetMaxFeesAndCutoff(ctx context.Context) (float32, float32, error) {
	postgresDb.mutex.Lock()
	defer postgresDb.mutex.Unlock()

	data := structure.CollegeData{}
	err := postgresDb.Db.GetContext(ctx, &data, "select max(fees) as fees, max(cutoff) as cutoff  from college_data;")
	if err != nil {
		return 0, 0, err
	}
	return data.Fees, data.Cutoff, nil

}

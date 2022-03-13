package structure

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
)

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type RangeData struct {
	MINFees   float32 `json:"min_fees" db:"min_fees"`
	MAXFees   float32 `json:"max_fees" db:"max_fees"`
	MINCutoff float32 `json:"min_cutoff" db:"min_cutoff"`
	MAXCutoff float32 `json:"max_cutoff" db:"max_cutoff"`
}

type CollegeData struct {
	CollegeId      string         `json:"college_id" db:"college_id"`
	Name           string         `json:"name" db:"name"`
	Address        string         `json:"address" db:"address"`
	Images         pq.StringArray `json:"images" db:"images"`
	Email          string         `json:"email" db:"email"`
	ContactNo      string         `json:"contact_no" db:"contact_no"`
	Website        string         `json:"website" db:"website"`
	VrImages       pq.StringArray `json:"vr_image" db:"vr_image"`
	Longitude      float32        `json:"longitude" db:"longitude"`
	Latitude       float32        `json:"latitude" db:"latitude"`
	Hostel         bool           `json:"hostel" db:"hostel"`
	Cutoff         float32        `json:"cutoff" db:"cutoff"`
	Fees           float32        `json:"fees" db:"fees"`
	InstituteType  bool           `json:"institute_type" db:"institute_type"`
	Details        string         `json:"details" db:"details"`
	Courses        pq.StringArray `json:"courses" db:"courses"`
	State          string         `json:"state" db:"state"`
	CourseType     string         `json:"course_type" db:"course_type"`
	Deemed         bool           `json:"deemed" db:"deemed"`
	DocumentTokens string         `json:"document_tokens" db:"document_tokens"`

	Distance float32 `json:"distance" db:"distance"`
}

func (college *CollegeData) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(college)
	return data, err
}

func UnmarshalCollegeData(data []byte, college *CollegeData) {
	err := json.Unmarshal(data, &college)
	if err != nil {
		fmt.Println(err)
	}
}

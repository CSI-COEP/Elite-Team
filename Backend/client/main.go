package main

import (
	"CollegeLocator/pb"
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Fali to load enviromental variables")
	}

	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dialing to server  : %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial server ", err)
	}

	server := pb.NewCollegeLocatorServiceClient(conn)
	ctx := context.TODO()
	initData, err := server.InitApp(ctx, &pb.InitRequest{
		Location: &pb.Location{
			Longitude: 21.246578216552734,
			Latitude:  75.2966537475586,
		},
	})
	if err != nil {
		log.Fatal("Init Request Error :  ", err)
	}

	for {
		res, err := initData.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(res.String())
		fmt.Println(res.GetCollegeId())
		fmt.Println(res.GetName())
		fmt.Println(res.GetAddress())
		fmt.Println(res.GetImage())
		fmt.Println(res.GetMaxRangeFees())
	}
}

// NP_LOCATION : 	18.50378107199465, 73.80971162151037

//
/////*
//	base := database.InitPostgresDataBase()
//	ctx := context.TODO()
//
//	err := base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:      "1",
//		Name:           "College Of Engineering Pune",
//		Address:        "College of Engg. Pune,Wellesely Rd,Shivajinagar,Pune-411 005.Maharashtra, INDIA.",
//		Images:         []string{"https://orchidfoundation.info/sites/default/files/2021-03/College%20of%20Engineering.jpeg",
//								"https://www.coep.org.in/sites/default/files/slideshow%20images/coepfinal1.jpg"},
//		Email:          "coep@coep.ac.in",
//		ContactNo:      "+912025507000",
//		Website:        "https://www.coep.org.in/",
//		VrImages:       []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:      21.247335363588483,
//		Latitude:       75.2948708561471,
//		Hostel:         true,
//		Cutoff:         99,
//		Fees:           900000,
//		InstituteType:  false,	// GOV = false OR PRIVATE = true
//		Details:        "College of Engineering, Pune (CoEP), chartered in 1854 is a nationally respected leader in technical education. The institute is distinguished by its commitment to finding solutions to the great predicaments of the day through advanced technology. The institute has a rich history and dedication to the pursuit of excellence. CoEP offers a unique learning experience across a spectrum of academic and social experiences. With a firm footing in truth and humanity, the institute gives you an understanding of both technical developments and the ethics that go with it. The curriculum is designed to enhance your academic experience through opportunities like internships, study abroad programmes and research facilities. The hallmark of CoEP education is its strong and widespread alumni network, support of the industry and the camaraderie that the institute shares with several foreign universities. The institute is consistently ranked amongst the top 20 technical colleges in India and its alumni have contributed a lion’s share in development of national infrastructure.",
//		Courses:        []string{"COMPUTER ENGINEERING", "MECHANICAL ENGINEERING", "ELECTRONICS AND TELECOMMUNICATION ENGINEERING", "CIVIL ENGINEERING", "PRODUCTION ENGINEERING"},
//		State:          "Maharashtra",
//		CourseType: 	"ENGINEERING",
//		Deemed:         false,
//	})
//	if err != nil {
//		fmt.Println("Insert 1 : ", err)
//	}
//
//	err = base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:     "2",
//		Name:          "SCTR'S Pune Institute of Computer Technology",
//		Address:       "Survey No. 27, Near Trimurti Chowk, Dhankawadi, Pune-411043, Maharashtra (India).",
//		Images:        []string{"https://cache.careers360.mobi/media/presets/720X480/colleges/social-media/media-gallery/3137/2018/8/8/Pune-Institute-of-Computer-Technology-Pune.png",
//								"https://www.campustimespune.com/wp-content/uploads/2015/03/PICT-College-Building-Pune.jpg"},
//		Email:         "pict@pic.ac.in",
//		ContactNo:     "+912024371101",
//		Website:       "https://pict.edu/",
//		VrImages:      []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:     18.50454141602843,
//		Latitude:      73.80925940687749,
//		Hostel:        true,
//		Cutoff:        90,
//		Fees:          1000000,
//		InstituteType: true,
//		Details:       "PICT believes in value-based quality education in the field of Information and Communication Technology (ICT) and we are constantly endeavouring to achieve higher levels of technical excellence by offering under-graduate (UG) programmes in Computer Engineering (CE), Electronics & Telecommunication Engineering (E&TC), and Information Technology (IT); and post-graduate (PG) programmes in Computer Engineering, Data Science, Electronics & Communication (Wireless Communication Technology), and Information Technology. It is also Savitribai Phule Pune University (SPPU) recognized Research Center in Computer Engineering and Electronics & Telecommunication Engineering.",
//		Courses:       []string{"COMPUTER ENGINEERING", "ELECTRONICS AND TELECOMMUNICATION ENGINEERING", "INFORMATION TECHNOLOGY"},
//		State:         "Maharashtra",
//		CourseType: 	"ENGINEERING",
//		Deemed:        false,
//	})
//	if err != nil {
//		fmt.Println("Insert 2 : ", err)
//	}
//
//	err = base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:     "3",
//		Name:          "Pimpri Chinchwad Education Trust's Pimpri Chinchwad College Of Engineering",
//		Address:       "Sector -26, Pradhikaran, Nigdi,Near Akurdi Railway Station, Pune - 411 044.",
//		Images:        []string{"https://images.collegedunia.com/public/college_data/images/appImage/14564_PCCE_New.jpg",
//								"https://images.shiksha.com/mediadata/images/1600151627phpOVW3t7.png"},
//		Email:         "pccoeadmin@gmail.com",
//		ContactNo:     "+919960984347 ",
//		Website:       "http://www.pccoepune.com/",
//		VrImages:      []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:     18.505892711545275,
//		Latitude:      73.80510927697429,
//		Hostel:        true,
//		Cutoff:        85,
//		Fees:          1010000,
//		InstituteType: true,
//		Details:       "PCCoE is functioning proactively to provide the best professional environment to engineering and management students in terms of academics, industry oriented trainings, sports, co-curricular & extracurricular activities, cultural activities, various competitions, etc to create true aesthetically sensitive, socially committed and technologically competent engineers and management professionals.\n",
//		Courses:        []string{"COMPUTER ENGINEERING", "MECHANICAL ENGINEERING", "ELECTRONICS AND TELECOMMUNICATION ENGINEERING", "CIVIL ENGINEERING", "PRODUCTION ENGINEERING"},
//		State:         "Maharashtra",
//		CourseType: 	"ENGINEERING",
//		Deemed:        true,
//	})
//	if err != nil {
//		fmt.Println("Insert 3 : ", err)
//	}
//
//	err = base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:     "4",
//		Name:          "Government Polytechnic, Pune",
//		Address:       "Ganeshkhind, University Road,Shivajinagar,Pune-411016, Maharashtra, India.",
//		Images:        []string{"https://gppune.ac.in/images/slider1n2.jpg",
//								"https://www.gpawasari.ac.in/images/Two.JPG"},
//		Email:         "gpune@gmail.com",
//		ContactNo:     "+912025559200",
//		Website:       "https://gppune.ac.in/",
//		VrImages:      []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:     18.506530703142143,
//		Latitude:      73.80947437050042,
//		Hostel:        true,
//		Cutoff:        90,
//		Fees:          60000,
//		InstituteType: false,
//		Details:       "The Government Polytechnic, Pune is one of the prominent institutes among polytechnics in Maharashtra. The institute was established in the year 1957, under the academic control of Maharashtra State Board of Technical Education, Maharashtra State. The institute was awarded academic autonomy in May 1994.",
//		Courses:        []string{"DIPLOMA COMPUTER ENGINEERING", "DIPLOMA MECHANICAL ENGINEERING", "DIPLOMA ELECTRONICS AND TELECOMMUNICATION ENGINEERING", "DIPLOMA CIVIL ENGINEERING", "DIPLOMA PRODUCTION ENGINEERING"},
//		State:         "Maharashtra",
//		CourseType: 	"POLYTECHNIC",
//		Deemed:        true,
//	})
//	if err != nil {
//		fmt.Println("Insert 4 : ", err)
//	}
//
//	err = base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:     "5",
//		Name:          "Pimpri Chinchwad Polytechnic College",
//		Address:       "Sector No. 26, Pradhikaran, Nigdi, Pune - 411 044",
//		Images:        []string{"http://pcet.org.in/images/PCET-Institutes-PCP.png",
//								"https://images.shiksha.com/mediadata/images/1556012345phpr1ztG3.png"},
//		Email:         "pcpolytechnic@gmail.com",
//		ContactNo:     "+912027654156",
//		Website:       "https://www.pcpolytechnic.com",
//		VrImages:      []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:     21.246408199486382,
//		Latitude:      75.28935073013615,
//		Hostel:        false,
//		Cutoff:        90,
//		Fees:          50000,
//		InstituteType: true,
//		Courses:        []string{"DIPLOMA COMPUTER ENGINEERING", "DIPLOMA MECHANICAL ENGINEERING", "DIPLOMA ELECTRONICS AND TELECOMMUNICATION ENGINEERING", "DIPLOMA CIVIL ENGINEERING", "DIPLOMA PRODUCTION ENGINEERING"},
//		Details:       "The trust is composed group of highly active, devoted persons of high social standing. They have joined together to establish, develop and run a large technical and allied education institutes comprising various useful educational courses. This suburb being the largest Industrial area in Asia needed enormous technically skilled manpower.\n\nWith these dual objectives of public and industrial needs trust invited reputed citizens from this region to participate in this Local and National endeavour and established Pimpri Chinchwad Education Trust. The trust then introduced Diploma Courses in Engineering by starting the Pimpri Chinchwad Polytechnic College in Pune From the year 1990.",
//		State:         "Maharashtra",
//		CourseType: 	"POLYTECHNIC",
//		Deemed:        false,
//	})
//	if err != nil {
//		fmt.Println("Insert 5 : ", err)
//	}
//
//	err = base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:     "6",
//		Name:          "AFMC Pune - Armed Forces Medical College",
//		Address:       "Southern Command, Solapur - Pune Hwy, near Race Course, Wanowrie, Pune, Maharashtra 411040",
//		Images:        []string{"https://afmc.nic.in/assests/img/afmc-whyus-background.png",
//								"https://afmc.nic.in/assests/img/afmc-student-background.jpg"},
//		Email:         "afmc@afmc.nic.in",
//		ContactNo:     "+912026334230",
//		Website:       "https://www.afmc.nic.in/",
//		VrImages:      []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:     18.504772344329982,
//		Latitude:      73.81530799111948,
//		Hostel:        false,
//		Cutoff:        90,
//		Fees:          10000000,
//		InstituteType: false,
//		Courses:        []string{"MBBS", "NURSING", "PARA MEDICAL", "PG DIPLOMA", "SUPER SPECIALITY"},
//		Details:       "Armed Forces Medical College was established on 01 May 1948. The institution conducts training of medical undergraduates and post-graduates, nursing undergraduates and post-graduates, dental postgraduates and paramedical staff.  Patient care forms an integral part of the training curriculum and the affiliated hospitals utilise the expertise available at AFMC.  The institution is responsible for providing the entire pool of specialists and super specialists to Armed Forces by giving them in-service training.  Thus AFMC forms the backbone of high quality professional medical care being provided to the Armed Forces.",
//		State:         "Maharashtra",
//		CourseType:    "MEDICAL",
//		Deemed:        false,
//	})
//	if err != nil {
//		fmt.Println("Insert 6 : ", err)
//	}
//
//	err = base.CreateCollege(ctx, &structure.CollegeData{
//		CollegeId:     "7",
//		Name:          "Bharati Vidyapeeth Deemed University Medical College",
//		Address:       "Medical College Road, Pune - Satara Rd, Dhankawadi, Pune, Maharashtra 411043",
//		Images:        []string{"https://content3.jdmagicbox.com/comp/pune/w9/020pxx20.xx20.100913150650.f3w9/catalogue/medical-college-pune-dhankawadi-pune-medical-colleges-4fk9mv2.png",
//								"https://mcpune.bharatividyapeeth.edu/media/images/MCPune_Home-1_191118.jpg"},
//		Email:         "afmc@afmc.nic.in",
//		ContactNo:     "+912026334230",
//		Website:       "https://mcpune.bharatividyapeeth.edu/",
//		VrImages:      []string{"https://github.com/mr-shitij/localHost-With-Python/blob/master/test.jpg?raw=true"},
//		Longitude:     21.240154077085013,
//		Latitude:      75.29583882538645,
//		Hostel:        true,
//		Cutoff:        95,
//		Fees:          10000000,
//		InstituteType: false,
//		Courses:        []string{"MBBS", "NURSING", "PARA MEDICAL", "PG DIPLOMA", "SUPER SPECIALITY"},
//		Details:       "Bharati Vidyapeeth (Deemed to be University) Medical College, Pune was established in February 1989 and is renowned for its academic excellence and infrastructural facilities. It was initially affiliated to the University of Pune till it became a constituent unit of Bharati Vidyapeeth (Deemed to be University) in April 1996.",
//		State:         "Maharashtra",
//		CourseType:    "MEDICAL",
//		Deemed:        true,
//	})
//	if err != nil {
//		fmt.Println("Insert 7 : ", err)
//	}
//

//var val int32 = 10
//deem := false
//ser := "chopda"
//dist := pb.Distance_LONG_RANGE
//request := pb.SearchRequest{
//	Location: &pb.Location{
//		Longitude: 21.25054762117774,
//		Latitude:  75.31465684097351,
//	},
//	SearchQuery:   &ser,
//	Hostel:        nil,
//	Cutoff:        nil,
//	Distance:      &dist,
//	Fees:          nil,
//	InstituteType: nil,
//	CourseId:      nil,
//	State:         nil,
//	Deemed:        &deem,
//}
//
//err = base.SearchCollege(ctx, &request,
//	func(data *structure.CollegeData) error {
//		fmt.Println(data)
//		return nil
//	})

//	//*/
//}
/*
	base := database.InitRedisDataBase()
	err := base.CreateSchema()
	if err != nil {
		fmt.Println("Schema : ", err)
	}

	ctx := context.TODO()

	err = base.SETCollege(ctx, &structure.CollegeData{
		CollegeId: "1",
		Name:      "College Of Engineering Pune",
		Address:   "Shivaji Nagar",
		Images:    []string{"https://image1.com/image", "https://image2.com/image"},
		Email:     "coep@coep.ac.in",
		ContactNo: "1234567890",
		Website:   "www.coep.in",
		VrImages:  []string{"url1", "url2"},
		Location: structure.Location{
			Longitude: 0,
			Latitude:  0,
		},
		Hostel:        true,
		Cutoff:        99,
		Fees:          100000,
		InstituteType: pb.InstituteType_GOV,
		Details:       "Some INFO",
		CourseId:      3,
		Courses: structure.Courses{
			EngineeringCourses: []pb.EngineeringCourses{pb.EngineeringCourses_COMPUTER_ENG, pb.EngineeringCourses_MECHANICAL_ENG},
			MedicalCourses:     nil,
			DiplomaCourses:     []pb.DiplomaCourses{pb.DiplomaCourses_COMPUTER_DIP, pb.DiplomaCourses_MECHANICAL_DIP},
		},
	})
	if err != nil {
		fmt.Println("Insert 1 : ", err)
	}

	err = base.SETCollege(ctx, &structure.CollegeData{
		CollegeId: "2",
		Name:      "Smt Sharchchandrika suresh patil institute of technology chopda",
		Address:   "Some thing",
		Images:    []string{"https://image1.com/image", "https://image2.com/image"},
		Email:     "ssp@ssp.ac.in",
		ContactNo: "1234567890",
		Website:   "www.ssp.in",
		VrImages:  []string{"url1", "url2"},
		Location: structure.Location{
			Longitude: 0,
			Latitude:  0,
		},
		Hostel:        true,
		Cutoff:        80,
		Fees:          25,
		InstituteType: pb.InstituteType_GOV,
		Details:       "Some INFO",
		CourseId:      2,
		Courses: structure.Courses{
			EngineeringCourses: nil,
			MedicalCourses:     nil,
			DiplomaCourses:     []pb.DiplomaCourses{pb.DiplomaCourses_COMPUTER_DIP, pb.DiplomaCourses_MECHANICAL_DIP},
		},
	})
	if err != nil {
		fmt.Println("Insert 2 : ", err)
	}

	fmt.Println("Complete ..!!")
	doc, err := base.SearchWithFilters(&pb.SearchRequest{
		Location:           nil,
		SearchQuery:        nil,
		Hostel:             true,
		EngineeringCourses: nil,
		MedicalCourses:     nil,
		DiplomaCourses:     nil,
		Cutoff:             nil,
		Distance:           nil,
		Fees:               nil,
		InstituteType:      0,
		CourseId:           nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Complete ..!!", doc)


*/

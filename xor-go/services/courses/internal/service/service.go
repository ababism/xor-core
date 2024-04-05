package service

import (
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.CoursesService = &CoursesService{}

type CoursesService struct {
	course     adapters.CourseRepository
	courseEdit adapters.CourseRepository
	lesson     adapters.LessonRepository
	lessonEdit adapters.LessonRepository

	student adapters.StudentRepository
	teacher adapters.TeacherRepository

	publication adapters.PublicationRequestRepository

	financesClient adapters.FinancesClient
}

func NewCoursesService(
	courseRepository adapters.CourseRepository,
	courseEditRepository adapters.CourseRepository,
	lessonRepository adapters.LessonRepository,
	lessonEditRepository adapters.LessonRepository,
	teacherRepository adapters.TeacherRepository,
	studentRepository adapters.StudentRepository,
	publicationRequestRepository adapters.PublicationRequestRepository,
	financesClient adapters.FinancesClient) adapters.CoursesService {

	return &CoursesService{
		course:         courseRepository,
		courseEdit:     courseEditRepository,
		lesson:         lessonRepository,
		lessonEdit:     lessonEditRepository,
		student:        studentRepository,
		teacher:        teacherRepository,
		publication:    publicationRequestRepository,
		financesClient: financesClient,
	}
}

//func (s *CoursesService) GetTripByID(ctx context.Context, driverId uuid.UUID, tripId uuid.UUID) (*domain.Trip, error) {
//	logger := zapctx.Logger(ctx)
//
//	tr := global.Tracer(domain.ServiceName)
//	newCtx, span := tr.Start(ctx, "driver/service.GetTripByID")
//	defer span.End()
//
//	// err if trip driver != nil and driver != driverId
//	trip, err := s.r.GetTripByID(newCtx, tripId)
//	if err != nil {
//		logger.Error("driver-service: get trip from repository")
//		return nil, err
//	}
//	if trip.DriverId != nil && *trip.DriverId != driverId.String() {
//		return nil, errors.Wrap(domain.ErrAccessDenied, "trip driver id does not match passed id")
//	}
//	return trip, err
//}
//
//// InsertTrip inserts a trip
//func (s *CoursesService) InsertTrip(ctx context.Context, trip domain.Trip) error {
//	log := zapctx.Logger(ctx)
//
//	tr := global.Tracer(domain.ServiceName)
//	newCtx, span := tr.Start(ctx, "driver/service.InsertTrip")
//	defer span.End()
//
//	err := s.r.InsertTrip(newCtx, trip)
//	if err != nil {
//		log.Error("driver-service: insert trip in repository", zap.Error(err))
//		return err
//	}
//
//	return nil
//}

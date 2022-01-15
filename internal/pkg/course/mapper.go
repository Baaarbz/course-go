package course

import domain "barbz.dev/course-go/internal"

func mapListCourseToDto(courses []domain.Course) []DTO {
	coursesDto := make([]DTO, len(courses))
	for index, element := range courses {
		// _ is the index where we are
		// element is the element from courses for where we are
		course := mapCourseToDto(element)
		coursesDto[index] = course
	}
	return coursesDto
}

func mapCourseToDto(course domain.Course) DTO {
	return DTO{
		ID:          course.ID(),
		Name:        course.Name(),
		Description: course.Description(),
	}
}

func mapDtoToCourse(dto DTO) domain.Course {
	return domain.NewCourse(dto.ID, dto.Name, dto.Description)
}

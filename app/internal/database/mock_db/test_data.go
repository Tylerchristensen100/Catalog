package mock_db

import (
	"encoding/binary"
	"sort"
	"strings"

	"github.com/google/uuid"

	"catalog.tylerChristensen/internal/models"
)

func randInt() int32 {
	id, err := uuid.New().MarshalBinary()
	if err != nil {
		//handle error
	}
	return int32(binary.BigEndian.Uint32(id))
}

func GenerateTestCIPs() []models.Cip {
	cips := []models.Cip{
		{ID: randInt(), Cip: 12.34567, Name: "CIP 1", Description: "Description 1", Jobs: "Job 1, Job 2"},
		{ID: randInt(), Cip: 98.76543, Name: "CIP 2", Description: "Description 2", Jobs: "Job 3"},
		{ID: randInt(), Cip: 10.00000, Name: "CIP 3", Description: "Description 3", Jobs: "Job 4, Job 5, Job 6"},
		{ID: randInt(), Cip: 1.23456, Name: "CIP 4", Description: "Description 4", Jobs: "Job 7"},
	}

	sort.Slice(cips, func(i, j int) bool {
		return cips[i].ID < cips[j].ID
	})
	return cips
}

func GenerateTestCourses(majors []models.Major) []models.Course {
	courses := []models.Course{
		{ID: randInt(), Name: "Introduction to Programming", MajorCode: strings.ToLower(majors[0].Code), Major: majors[0], Code: "101", CreditHours: 3, CreatedBy: 1},
		{ID: randInt(), Name: "Data Structures and Algorithms", MajorCode: strings.ToLower(majors[1].Code), Major: majors[1], Code: "201", CreditHours: 4, CreatedBy: 10},
		{ID: randInt(), Name: "Calculus I", MajorCode: strings.ToLower(majors[2].Code), Major: majors[2], Code: "150", CreditHours: 4, CreatedBy: 2},
		{ID: randInt(), Name: "Linear Algebra", MajorCode: strings.ToLower(majors[3].Code), Major: majors[3], Code: "220", CreditHours: 3, CreatedBy: 2},
		{ID: randInt(), Name: "Principles of Economics", MajorCode: strings.ToLower(majors[4].Code), Major: majors[4], Code: "101", CreditHours: 3, CreatedBy: 4},
	}
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})
	return courses
}

func GenerateTestGradLevels() []models.GradLevel {
	levels := []models.GradLevel{
		{ID: randInt(), Level: ptr("Undergraduate")},
		{ID: randInt(), Level: ptr("Graduate")},
		{ID: randInt(), Level: ptr("Certificate")},
		{ID: randInt(), Level: ptr("Bootcamp")},
	}
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].ID < levels[j].ID
	})
	return levels
}

func GenerateTestMajors() []models.Major {
	majors := []models.Major{
		{ID: randInt(), Code: "CS", Name: "Computer Science"},
		{ID: randInt(), Code: "MBA", Name: "Master of Business Administration"},
		{ID: randInt(), Code: "DS", Name: "Data Science"},
		{ID: randInt(), Code: "WD", Name: "Web Development"},
		{ID: randInt(), Code: "IS", Name: "Information Systems"},
		{ID: randInt(), Code: "ECO", Name: "Economics"},
	}

	sort.Slice(majors, func(i, j int) bool {
		return majors[i].ID < majors[j].ID
	})
	return majors
}

func GenerateTestPrograms(gradLevels []models.GradLevel, schools []models.School) []models.Program {
	programs := []models.Program{
		{ID: randInt(),
			Name:        "Computer Science",
			GradLevelID: 1,
			GradLevel:   gradLevels[0],
			ProgramType: "Degree",
			School:      schools[0],
			MajorCode:   "CS",
			Online:      1,
			Campus:      0,
			Description: "A comprehensive CS program.",
			Cip:         11,
			CreatedBy:   1,
		},
		{ID: randInt(),
			Name:        "Master of Business Administration",
			GradLevel:   gradLevels[1],
			ProgramType: "Degree",
			School:      schools[1],
			MajorCode:   "MBA",
			Online:      0,
			Campus:      1,
			Description: "A challenging MBA program.",
			Cip:         52,
			CreatedBy:   2,
		},
		{ID: randInt(),
			Name:        "Data Science Certificate",
			GradLevel:   gradLevels[2],
			ProgramType: "Certificate",
			School:      schools[2],
			MajorCode:   "DS",
			Online:      1,
			Campus:      0,
			Description: "A focused data science program.",
			Cip:         27,
			CreatedBy:   1,
		},
		{ID: randInt(),
			Name:        "Web Development Bootcamp",
			GradLevel:   gradLevels[2],
			ProgramType: "Bootcamp",
			School:      schools[2],
			MajorCode:   "WD",
			Online:      1,
			Campus:      0,
			Description: "Intensive web dev training.",
			Cip:         15,
			CreatedBy:   1,
		},
	}
	sort.Slice(programs, func(i, j int) bool {
		return programs[i].Name < programs[j].Name
	})
	return programs
}

func GenerateTestUsers() []models.User {
	users := []models.User{
		{ID: randInt(), Username: "user1", Roles: "admin,faculty", ClientID: randInt()},
		{ID: randInt(), Username: "user2", Roles: "faculty", ClientID: randInt()},
		{ID: randInt(), Username: "user3", Roles: "faculty", ClientID: randInt()},
		{ID: randInt(), Username: "user4", Roles: "undefined", ClientID: randInt()},
		{ID: randInt(), Username: "user5", Roles: "admin", ClientID: randInt()},
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
	return users
}

func GenerateTestSchools() []models.School {
	schools := []models.School{
		{ID: randInt(), Code: "CET", Name: "College of Engineering & Technology"},
		{ID: randInt(), Code: "WSB", Name: "School of Business"},
		{ID: randInt(), Code: "SOA", Name: "School of the Arts"},
	}

	sort.Slice(schools, func(i, j int) bool {
		return schools[i].ID < schools[j].ID
	})
	return schools
}

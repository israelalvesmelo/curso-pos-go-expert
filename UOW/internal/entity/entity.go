package entity

type Category struct {
	ID      int
	Name    string
	Courses []int
}

func (c *Category) AddCourse(courseID int) {
	c.Courses = append(c.Courses, courseID)
}

type Course struct {
	ID         int
	Name       string
	CategoryID int
}

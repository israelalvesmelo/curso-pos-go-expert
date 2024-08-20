
# Documentation GraphQL

mutation createCategory {
  createCategory(input: {name: "Tecnologia", description: "Cursos de Tecnologia"})
  {
    id
    name
    description
  }
}

mutation createCourse {
  createCourse(input: {name: "Full Cycle 2", description:"The best", categoryId: "78b107f4-17cc-4a96-9078-fef53849b5c2" })
    {
      id
      name
      description
    }
}

query queryCaterogies {
  categories{
    id
    name
    description
    courses {
      id
      name
    }
  }
}

query queryCourses{
	courses{
  id
  name
  description
    category {
      id
      name
    }
	}  
}
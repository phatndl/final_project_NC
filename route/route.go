package route

import (
	"final_project/handler"

	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Public(e)
	Staff(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {
	g := e.Group("api/student/v1/staff")
	g.POST("/student", handler.AddStudent)
	g.PUT("/student", handler.UpdateStudent)
}

func Public(e *echo.Echo) {
	g := e.Group("api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test-add-number", handler.TestInsert)
	g.GET("/students", handler.GetAllStudents)
	g.PATCH("/student", handler.SearchStudentSimple)
}

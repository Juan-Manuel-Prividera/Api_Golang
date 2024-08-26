package main

import (
	"api/src/main/materias"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/materias", materias.GetMaterias)
	router.GET("/api/materias/:nivel", materias.BuscarMateriasPor)

	router.POST("/api/materia", materias.AgregarMateria)
	router.POST("/api/materia/delete/:id", materias.BorrarMateria)
	router.POST("/api/materia/update/:id", materias.ActualizarMateria)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Print(err)
	}
}

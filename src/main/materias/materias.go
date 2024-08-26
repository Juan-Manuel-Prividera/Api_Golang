package materias

import (
	"api/src/main/types"
	"api/src/main/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMaterias(c *gin.Context) {
	materias := utils.CrearMaterias()
	c.IndentedJSON(http.StatusOK, materias)
	return
}

func BuscarMateriasPor(c *gin.Context) {
	materias := utils.CrearMaterias()
	nivel, err := strconv.Atoi(c.Param("nivel"))
	if err != nil {
		fmt.Print(err)
	}

	var materiasFiltradas []*types.Materia
	for _, materia := range materias {
		if materia.Nivel == nivel {
			materiasFiltradas = append(materiasFiltradas, materia)
		}
	}

	c.IndentedJSON(http.StatusOK, materiasFiltradas)
	return
}

func AgregarMateria(c *gin.Context) {
	var materia types.Materia
	// Leo el cuerpo de la request
	if err := c.BindJSON(&materia); err != nil {
		fmt.Print(err)
		return
	}

	materias := utils.LeerJson()

	if utils.YaExisteEstaMateria(materia, materias) {
		return
	}

	utils.SetIdMateria(&materia)
	materias = append(materias, &materia)
	utils.EscribirArchivo(materias)

	c.JSON(http.StatusCreated, materia)

	return
}

func BorrarMateria(c *gin.Context) {
	idMateria, _ := strconv.Atoi(c.Param("id"))
	materias := utils.LeerJson()

	index := utils.EncontrarIdex(idMateria, materias)

	// Esto borra un elemento del slice
	materias = append(materias[:index], materias[index+1:]...)

	utils.EscribirArchivo(materias)
	c.JSON(http.StatusOK, materias)
	return
}

func ActualizarMateria(c *gin.Context) {
	idMateria, _ := strconv.Atoi(c.Param("id"))
	var materia types.Materia
	if err := c.BindJSON(&materia); err != nil {
		fmt.Print(err)
	}

	materias := utils.LeerJson()

	for _, materiaI := range materias {
		if materiaI.Id == idMateria {
			utils.UpdateMateria(materiaI, &materia)
		}
	}
	utils.EscribirArchivo(materias)
	c.JSON(http.StatusOK, materia)
}

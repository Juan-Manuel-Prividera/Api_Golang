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
	if materias == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Materias no encontradas", "message": "No se econtraron materias para ese nivel"})
	}
	//c.IndentedJSON(http.StatusOK, materias)
	c.IndentedJSON(http.StatusOK, gin.H{"materias": materias})
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
	if materiasFiltradas == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Materias no encontradas", "message": "No se econtraron materias para ese nivel"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"materias": materiasFiltradas})
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
	if materias == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Materias no encontradas", "message": "No se econtraron materias para ese nivel"})
	}
	//c.IndentedJSON(http.StatusOK, materias)
	c.IndentedJSON(http.StatusOK, gin.H{"materias": materias})
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
	if materiasFiltradas == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Materias no encontradas", "message": "No se econtraron materias para ese nivel"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"materias": materiasFiltradas})
	return
}

func BuscarMateriaPorId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Materia no econtrada, revise el id enviado"})
	}

	materias := utils.CrearMaterias()
	for _, materia := range materias {
		if materia.Id == id {
			c.IndentedJSON(http.StatusOK, gin.H{"materia": materia})
		}
	}
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

	c.IndentedJSON(http.StatusCreated, gin.H{"materia": materia})

	return
}

func BorrarMateria(c *gin.Context) {
	idMateria, _ := strconv.Atoi(c.Param("id"))
	materias := utils.LeerJson()

	index := utils.EncontrarIdex(idMateria, materias)

	// Esto borra un elemento del slice
	materias = append(materias[:index], materias[index+1:]...)

	utils.EscribirArchivo(materias)
	c.IndentedJSON(http.StatusOK, gin.H{"materia": materias})
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
	c.IndentedJSON(http.StatusOK, gin.H{"materia": materia})
}
	return
}

func BuscarMateriaPorId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Materia no econtrada, revise el id enviado"})
	}

	materias := utils.CrearMaterias()
	for _, materia := range materias {
		if materia.Id == id {
			c.IndentedJSON(http.StatusOK, gin.H{"materia": materia})
		}
	}
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

	c.IndentedJSON(http.StatusCreated, gin.H{"materia": materia})

	return
}

func BorrarMateria(c *gin.Context) {
	idMateria, _ := strconv.Atoi(c.Param("id"))
	materias := utils.LeerJson()

	index := utils.EncontrarIdex(idMateria, materias)

	// Esto borra un elemento del slice
	materias = append(materias[:index], materias[index+1:]...)

	utils.EscribirArchivo(materias)
	c.IndentedJSON(http.StatusOK, gin.H{"materia": materias})
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
	c.IndentedJSON(http.StatusOK, gin.H{"materia": materia})
}

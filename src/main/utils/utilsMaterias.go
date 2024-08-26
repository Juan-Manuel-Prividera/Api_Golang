package utils

import (
	"api/src/main/types"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func OpenMaterias() *os.File {
	file, err := os.Open("src/resources/materias.json")
	if err != nil {
		fmt.Print(err)
	}

	return file
}
func CrearMaterias() []*types.Materia {
	file := OpenMaterias()
	defer file.Close()

	materiasJson, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}

	var materias []*types.Materia
	err = json.Unmarshal(materiasJson, &materias)
	if err != nil {
		fmt.Print(err)
	}

	return materias
}

func LeerJson() []*types.Materia {
	file := OpenMaterias()
	defer file.Close()

	materiasJson, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}

	var materias []*types.Materia
	erro := json.Unmarshal(materiasJson, &materias)
	if erro != nil {
		fmt.Print(erro)
	}

	return materias
}

func EscribirArchivo(materias []*types.Materia) {
	contenido, _ := json.Marshal(materias)
	err := os.WriteFile("src/resources/materias.json", contenido, 0644)
	if err != nil {
		return
	}
}

func EncontrarIdex(id int, materias []*types.Materia) int {
	for i, materia := range materias {
		if materia.Id == id {
			return i
		}
	}
	return -1
}

func SetIdMateria(materia *types.Materia) {
	materias := LeerJson()
	maxId := 0
	for _, materiaI := range materias {
		if materiaI.Id >= maxId {
			maxId = materiaI.Id
		}
	}
	materia.Id = maxId + 1
}

func UpdateMateria(materiaI *types.Materia, materia *types.Materia) {
	if materia.Nombre != "" {
		materiaI.Nombre = materia.Nombre
	}
	if materia.Nivel != 0 {
		materiaI.Nivel = materia.Nivel
	}
	if materia.Especialidad != "" {
		materiaI.Especialidad = materia.Especialidad
	}
}

func YaExisteEstaMateria(materia types.Materia, materias []*types.Materia) bool {
	for _, materiaI := range materias {
		if materia.Id == materiaI.Id {
			return true
		}
	}
	return false
}

package lab12

import (
	"testing"

	"github.com/vipally/glab/lab12/filepath"
	"github.com/vipally/glab/lab12/pet"
)

func TestWalk(t *testing.T) {
	var itPet pet.Walker
	var itFilePath filepath.Walker
	println("call by pet.Walker:")
	itPet = &pet.Dog{}
	itPet.Walk()
	itPet = &pet.Cat{}
	itPet.Walk()
	itPet = &filepath.FilePath{}
	itPet.Walk()

	println("\ncall by filepath.Walker:")
	itFilePath = &pet.Dog{}
	itFilePath.Walk()
	itFilePath = &pet.Cat{}
	itFilePath.Walk()
	itFilePath = &filepath.FilePath{}
	itFilePath.Walk()

	println("\nThe strange is that pet.Warker and filepath.Walker has the same signiture but they are not the same one.")
	println("But Go treat them as the same one.")
	println(`
implements pet.Walker{
	*pet.Dog
	*pet.Cat
}

implements filepath.Walker{
	*filepath.FilePath
}
`)
}

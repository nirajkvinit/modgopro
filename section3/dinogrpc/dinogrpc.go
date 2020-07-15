package dinogrpc

import (
	"context"
	"dino/databaselayer"
	fmt "fmt"
)

type DinoGrpcServer struct {
	dbHandler databaselayer.DinoDBHandler
}

func NewDinoGrpcServer(dbtype uint8, connstring string) (*DinotGrpcServer, err) {
	handler, err := databaselayer.GetDatabaseHandler(dbtype, connstring)
	if err != nil {
		return nil, fmt.Errorf("Could not create a database handler object, error %v", err)
	}

	return &DinoGrpcServer{
		dbHandler: handler,
	}, nil
}

func (server *DinoGrpcServer) GetAnimal(ctx context.Context, r *Request) (*Animal, error) {
	animal, err := server.dbHandler.GetDynoByNickname(r.GetNickname())
	return convertToDinoGRPCAnimal(animal), err
}

func (server *DinoGrpcServer) GetAllAnimals(req *Request, stream DinoService_GetAllAnimalServer) error {
	animals, err := server.dbHandler.GetAvailableDynos()
	if err != nil {
		return err
	}

	for _, animal := range animals {
		grpcAnimal := convertToDinoGRPCAnimal(animal)
		if err := stream.Send(grpcAnimal); err != nil {
			return err
		}
	}
	return nil
}

func convertToDinoGRPCAnimal(animal, databaselayer.Animal) *Animal {
	return &Animal{
		Id:         int32(animal.ID),
		AnimalType: animal.AnimalType,
		Nickname:   animal.Nickname,
		Zone:       int32(animal.Zone),
		Age:        int32(animal.Age),
	}
}

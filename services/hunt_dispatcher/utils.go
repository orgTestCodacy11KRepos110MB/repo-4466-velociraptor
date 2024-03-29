package hunt_dispatcher

import (
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/services"
)

func GetArtifactSources(
	config_obj *config_proto.Config,
	artifact string) []string {
	result := []string{}
	manager, err := services.GetRepositoryManager(config_obj)
	if err != nil {
		return nil
	}

	repository, err := manager.GetGlobalRepository(config_obj)
	if err == nil {
		artifact_obj, pres := repository.Get(config_obj, artifact)
		if pres {
			for _, source := range artifact_obj.Sources {
				result = append(result, source.Name)
			}
		}
	}
	return result
}

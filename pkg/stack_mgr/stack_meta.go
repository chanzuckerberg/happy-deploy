package stack_mgr

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/chanzuckerberg/happy/pkg/backend"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type StackMeta struct {
	stackName string
	priority  int

	DataMap  map[string]string
	TagMap   map[string]string
	paramMap map[string]string
}

// Update the image tag with the given newTag, and set the priority randomly.
// To not collide, setting priority requirs knowing the the priority of all other
// stacks from the StackMgr
func (s *StackMeta) Update(
	newTag string,
	stackTags map[string]string,
	sliceName string,
	stackSvc *StackService,
) error {
	s.DataMap["imagetag"] = newTag
	s.DataMap["imagetags"] = ""

	if len(stackTags) > 0 {
		imageTagsJson, err := json.Marshal(stackTags)
		if err != nil {
			return errors.Wrap(err, "could not marshal json")
		}
		s.DataMap["imagetags"] = string(imageTagsJson)
	}

	s.DataMap["slice"] = sliceName

	now := strconv.FormatInt(time.Now().Unix(), 10)
	if _, ok := s.DataMap["created"]; !ok {
		s.DataMap["created"] = now
	}
	s.DataMap["updated"] = now

	if ownerVal, ok := s.DataMap["owner"]; !ok || ownerVal == "" {
		// TODO maybe move to util
		// NOTE this is an implicit dependency on the backend
		userIdBackend := backend.GetAwsSts(stackSvc.GetConfig())
		userName, err := userIdBackend.GetUserName()
		if err != nil {
			return err
		}

		s.DataMap["owner"] = userName
	}

	return s.setPriority(stackSvc)
}

func (s *StackMeta) setPriority(stackMgr *StackService) error {
	if s.priority != 0 {
		return nil
	}

	existingPrioirty := map[int]bool{}
	stacks, err := stackMgr.GetStacks()
	if err != nil {
		return err
	}

	for _, stack := range stacks {
		stackMeta, err := stack.Meta()
		if err != nil {
			return errors.Wrap(err, "failed to get stack information")
		}
		stackPriority := stackMeta.priority
		existingPrioirty[stackPriority] = true
	}
	// pick a random number between 1000 and 5000 that's not in use right now.
	priority := 0
	for priority == 0 {
		randInt := random.Random(1000, 5000)
		if _, ok := existingPrioirty[randInt]; !ok {
			priority = randInt
		}
	}
	s.DataMap["priority"] = strconv.Itoa(priority)
	log.WithField("priority", priority).Debug("Priority for workspace set")
	return nil
}

func (s *StackMeta) GetTags() map[string]string {
	out := map[string]string{}
	for k, v := range s.TagMap {
		out[v] = s.DataMap[k]
	}
	return out
}

func (s *StackMeta) GetParameters() map[string]string {
	out := map[string]string{}
	for k, v := range s.paramMap {
		out[v] = s.DataMap[k]
	}
	return out
}

func (s *StackMeta) Load(existingTags map[string]string) error {
	for shortTag, tagName := range s.TagMap {
		if _, ok := existingTags[tagName]; ok {
			s.DataMap[shortTag] = existingTags[tagName]
		} else if _, ok := s.DataMap[shortTag]; !ok {
			s.DataMap[shortTag] = ""
		}
	}
	return nil
}

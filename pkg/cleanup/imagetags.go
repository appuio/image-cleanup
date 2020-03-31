package cleanup

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/appuio/seiso/pkg/openshift"
	imagev1 "github.com/openshift/api/image/v1"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// MatchOption type defines how the tags should be matched
type MatchOption string

const (
	// MatchOptionExact for exact matches
	MatchOptionExact MatchOption = "exact"
	// MatchOptionPrefix for prefix matches
	MatchOptionPrefix MatchOption = "prefix"
)

// GetMatchingTags returns all image tags matching one of the provided git tags
func GetMatchingTags(gitTags, imageTags *[]string, matchOption MatchOption) []string {
	var matchingTags []string

	log.WithFields(log.Fields{
		"match":     matchOption,
		"imageTags": imageTags,
		"gitTags":   gitTags,
	}).Debug("Matching imageTags with gitTags")

	if len(*gitTags) == 0 || len(*imageTags) == 0 {
		return matchingTags
	}

	for _, gitTag := range *gitTags {
		for _, imageTag := range *imageTags {
			if match(imageTag, gitTag, matchOption) {
				matchingTags = append(matchingTags, imageTag)
				log.WithFields(log.Fields{
					"gitTag":   gitTag,
					"imageTag": imageTag,
				}).Debug("Found matching tag")
			}
		}
	}

	return matchingTags
}

// GetInactiveImageTags returns the tags without active tags (unsorted)
func GetInactiveImageTags(activeTags, allImageTags *[]string) []string {
	inactiveTags := funk.FilterString(*allImageTags, func(imageTag string) bool {
		return !funk.ContainsString(*activeTags, imageTag)
	})
	return inactiveTags
}

// FilterOrphanImageTags returns the tags that do not have any git commit match
func FilterOrphanImageTags(gitValues, imageTags *[]string, matchOption MatchOption) []string {

	log.WithFields(log.Fields{
		"imageTagsToFilter": imageTags,
		"gitTagsToFilter":   gitValues,
	}).Debug("Filtering image tags by commits...")

	orphans := funk.FilterString(*imageTags, func(imageTag string) bool {
		for _, gitValue := range *gitValues {
			if match(imageTag, gitValue, matchOption) {
				return false
			}
		}
		return true
	})
	return orphans
}

// FilterByRegex returns the tags that match the regexp
func FilterByRegex(imageTags *[]string, regexp *regexp.Regexp) []string {
	var matchedTags []string

	log.WithField("pattern:", regexp).Debug("Filtering image tags with regex...")

	for _, tag := range *imageTags {
		imageTagMatched := regexp.MatchString(tag)
		log.WithFields(log.Fields{
			"imageTag:": tag,
			"match:":    imageTagMatched,
		}).Debug("Matching image tag")
		if imageTagMatched {
			matchedTags = append(matchedTags, tag)
		}
	}
	return matchedTags
}

// LimitTags returns the tags which should not be kept by removing the first n tags
func LimitTags(tags *[]string, keep int) []string {
	if len(*tags) > keep {
		limitedTags := make([]string, len(*tags)-keep)
		copy(limitedTags, (*tags)[keep:])
		return limitedTags
	}

	return []string{}
}

// FilterImageTagsByTime returns the tags which are older than the specified time
func FilterImageTagsByTime(imageStreamObjectTags *[]imagev1.NamedTagEventList, olderThan time.Time) []string {
	var imageStreamTags []string

	for _, imageStreamTag := range *imageStreamObjectTags {
		lastUpdatedDate := imageStreamTag.Items[0].Created.Time
		for _, tagEvent := range imageStreamTag.Items {
			if lastUpdatedDate.Before(tagEvent.Created.Time) {
				lastUpdatedDate = tagEvent.Created.Time
			}
		}

		if lastUpdatedDate.Before(olderThan) {
			imageStreamTags = append(imageStreamTags, imageStreamTag.Tag)
		}
	}

	return imageStreamTags
}

func match(imageTag, value string, matchOption MatchOption) bool {
	switch matchOption {
	case MatchOptionPrefix:
		return strings.HasPrefix(imageTag, value)
	case MatchOptionExact:
		return imageTag == value
	}
	return false
}

// FilterActiveImageTags first gets all actively used image tags from imageStreamTags, then filters them out from matchingTags
func FilterActiveImageTags(namespace string, imageName string, imageStreamTags []string, matchingTags *[]string) ([]string, error) {
	activeImageStreamTags, err := openshift.GetActiveImageStreamTags(namespace, imageName, imageStreamTags)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve active image tags from %v/%v': %w", namespace, imageName, err)
	}

	log.WithField("activeTags", activeImageStreamTags).Debug("Found currently active image tags")
	return GetInactiveImageTags(&activeImageStreamTags, matchingTags), nil
}

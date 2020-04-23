package cmd

import (
	"fmt"
	"strings"

	"github.com/appuio/seiso/cfg"
	"github.com/appuio/seiso/pkg/git"
	"github.com/appuio/seiso/pkg/kubernetes"
	"github.com/appuio/seiso/pkg/openshift"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeleteImages deletes a list of image tags
func DeleteImages(imageTags []string, imageName string, namespace string, delete bool) {
	for _, inactiveTag := range imageTags {
		logEvent := log.WithFields(log.Fields{
			"namespace": namespace,
			"image":     imageName,
			"imageTag":  inactiveTag,
		})
		if delete {
			if err := openshift.DeleteImageStreamTag(namespace, openshift.BuildImageStreamTagName(imageName, inactiveTag)); err != nil {
				logEvent.WithError(err).Error("Failed to delete")
			} else {
				logEvent.Info("Deleted")
			}
		} else {
			logEvent.Info("Should be deleted")
		}
	}
}

// DeleteResources deletes a list of ConfigMaps or secrets
func DeleteResources(resources []cfg.KubernetesResource, delete bool, resourceSelectorFunc cfg.ResourceNamespaceSelector) {
	for _, resource := range resources {
		kind := resource.GetKind()
		name := resource.GetName()
		logEvent := log.WithFields(log.Fields{
			"namespace": resource.GetNamespace(),
			kind:        name,
		})
		if delete {
			if err := openshift.DeleteResource(name, resourceSelectorFunc); err != nil {
				logEvent.WithError(err).Error("Failed to delete")
			} else {
				logEvent.Info("Deleted")
			}
		} else {
			logEvent.Info("Should be deleted")
		}
	}
}

// PrintImageTags prints the given image tags line by line. In batch mode, only the tag name is printed, otherwise default
// log with info level
func PrintImageTags(imageTags []string) {
	if config.Log.Batch {
		for _, tag := range imageTags {
			fmt.Println(tag)
		}
	} else {
		for _, tag := range imageTags {
			log.WithField("imageTag", tag).Info("Found image tag candidate")
		}
	}
}

// PrintResources prints the given resource line by line. In batch mode, only the resource is printed, otherwise default
// log with info level
func PrintResources(resources []cfg.KubernetesResource) {
	if config.Log.Batch {
		for _, resource := range resources {
			fmt.Println(resource.GetKind() + ": " + resource.GetName())
		}
	} else {
		for _, resource := range resources {
			log.WithField(resource.GetKind(), resource.GetName()).Info("Found resource candidate")
		}
	}
}

// addCommonFlagsForGit sets up the delete flag, as well as the common git flags. Adding the flags to the root cmd would make those
// global, even for commands that do not need them, which might be overkill.
func addCommonFlagsForGit(cmd *cobra.Command, defaults *cfg.Configuration) {
	cmd.PersistentFlags().BoolP("delete", "d", defaults.Delete, "Confirm deletion of image tags.")
	cmd.PersistentFlags().BoolP("force", "f", defaults.Delete, "(deprecated) Confirm deletion. Alias for --delete")
	cmd.PersistentFlags().IntP("commit-limit", "l", defaults.Git.CommitLimit,
		"Only look at the first <l> commits to compare with tags. Use 0 (zero) for all commits. Limited effect if repo is a shallow clone.")
	cmd.PersistentFlags().StringP("repo-path", "p", defaults.Git.RepoPath, "Path to Git repository")
	cmd.PersistentFlags().BoolP("tags", "t", defaults.Git.Tag,
		"Instead of comparing commit history, it will compare git tags with the existing image tags, removing any image tags that do not match")
	cmd.PersistentFlags().String("sort", defaults.Git.SortCriteria,
		fmt.Sprintf("Sort git tags by criteria. Only effective with --tags. Allowed values: [%s, %s]", git.SortOptionVersion, git.SortOptionAlphabetic))
}

func listImages() error {
	ns, err := kubernetes.Namespace()
	if err != nil {
		return err
	}
	imageStreams, err := openshift.ListImageStreams(ns)
	if err != nil {
		return err
	}
	imageNames := []string{}
	for _, image := range imageStreams {
		imageNames = append(imageNames, image.Name)
	}
	log.WithFields(log.Fields{
		"\n - project":  ns,
		"\n - 📺 images": imageNames,
	}).Info("Please select an image. The following images are available:")
	return nil
}

func listConfigMaps(args []string) error {
	namespace, err := getNamespace(args)
	if err != nil {
		return err
	}

	configMaps, err := openshift.ListConfigMaps(namespace, metav1.ListOptions{})
	if err != nil {
		return err
	}

	configMapNames, labels := getNamesAndLabels(configMaps)

	log.WithFields(log.Fields{
		"\n - project":      namespace,
		"\n - 🔓 configMaps": configMapNames,
		"\n - 🎫 labels":     labels,
	}).Info("Please use labels to select ConfigMaps. The following ConfigMaps and Labels are available:")
	return nil
}

func listSecrets(args []string) error {
	namespace, err := getNamespace(args)
	if err != nil {
		return err
	}
	secrets, err := openshift.ListSecrets(namespace, metav1.ListOptions{})
	if err != nil {
		return err
	}

	secretNames, labels := getNamesAndLabels(secrets)
	log.WithFields(log.Fields{
		"\n - project":   namespace,
		"\n - 🔐 secrets": secretNames,
		"\n - 🎫 labels":  labels,
	}).Info("Please use labels to select Secrets. The following Secrets and Labels are available:")
	return nil
}

func getNamesAndLabels(resources []cfg.KubernetesResource) (resourceNames, labels []string) {
	for _, resource := range resources {
		resourceNames = append(resourceNames, resource.GetName())
		for key, element := range resource.GetLabels() {
			label := key + "=" + element
			if !funk.ContainsString(labels, label) {
				labels = append(labels, label)
			}
		}
	}

	return resourceNames, labels
}

//GetListOptions returns a ListOption object based on labels
func getListOptions(labels []string) metav1.ListOptions {
	labelSelector := fmt.Sprintf(strings.Join(labels, ","))
	return metav1.ListOptions{
		LabelSelector: labelSelector,
	}
}

func getNamespace(args []string) (string, error) {
	if len(args) == 0 {
		namespace, err := kubernetes.Namespace()
		if err != nil {
			return "", err
		}
		return namespace, err
	} else {
		return args[0], nil
	}
}

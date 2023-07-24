package githubconnector

import (
	"context"
	"strings"

	"github.com/google/go-github/v53/github"
)

func (client *GithubConnector) GetRelease(org, project, version string) (*Release, error) {

	// I don't like this, but we have to add it somewhere. Bury it here so everything else can
	// work on the basis of versions instead of tags.
	versionTag := "v" + version

	ghRelease, _, err := client.github.Repositories.GetReleaseByTag(context.Background(), org, project, versionTag)

	if err != nil {
		return nil, err
	}

	return &Release{
		Tag:     *ghRelease.TagName,
		Version: strings.Replace(*ghRelease.TagName, "v", "", 1),
		Assets:  getAssetsForRelease(ghRelease),
	}, nil

}

func (client *GithubConnector) GetReleases(org string, project string) ([]*Release, error) {

	happyReleases := make([]*Release, 0)

	// Only up to 1000 results are returned by the API
	for page := 1; page < 2; page++ {
		releases, _, err := client.github.Repositories.ListReleases(context.TODO(), org, project,
			&github.ListOptions{
				Page:    page,
				PerPage: 100,
			})

		if err != nil {
			return nil, err
		}

		if len(releases) == 0 {
			break
		}

		for _, release := range releases {

			if strings.HasPrefix(*release.TagName, "v") {
				happyReleases = append(happyReleases, &Release{
					Tag:     *release.TagName,
					Version: tagToVersion(*release.TagName),
					Assets:  getAssetsForRelease(release),
				})
			}
		}
	}

	return happyReleases, nil

}

func getAssetsForRelease(release *github.RepositoryRelease) []ReleaseAsset {
	assets := make([]ReleaseAsset, 5)

	for _, asset := range release.Assets {

		assets = append(assets, ReleaseAsset{
			Name:         asset.GetName(),
			Component:    nameToComponent(asset.GetName()),
			OS:           nameToOS(asset.GetName()),
			Architecture: nameToArchitecture(asset.GetName()),
			URL:          asset.GetBrowserDownloadURL(),
			FileType:     asset.GetContentType(),
		})

	}

	return assets
}

func tagToVersion(tag string) string {
	return strings.Replace(tag, "v", "", 1)
}
func nameToComponent(name string) string {
	return strings.Split(name, "_")[0]
}

func nameToArchitecture(name string) string {
	parts := strings.Split(name, "_")

	if len(parts) < 4 {
		return ""
	}

	archAndExtension := strings.Split(name, "_")[3]
	return strings.Split(archAndExtension, ".")[0]
}

func nameToOS(label string) string {
	os := strings.Split(label, "_")[2]

	if os != "checksums.txt" {
		return os
	}

	return ""
}

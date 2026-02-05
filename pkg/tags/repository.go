package tags

import (
	"bytes"
	"cv-landing-cli/pkg/client"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TagsClient struct {
	Base *client.BaseClient
}

func (h *TagsClient) Get(filter TagFilter) ([]Tag, error) {
	apiLink, err := h.buildGetQuery(filter)
	if err != nil {
		return []Tag{}, err
	}
	resp, err := h.Base.HTTPClient.Get(apiLink)
	if err != nil {
		return []Tag{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Tag{}, errors.New("http is not OK")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Tag{}, err
	}
	var tags []Tag
	err = json.Unmarshal(bodyBytes, &tags)
	if err != nil {
		return []Tag{}, err
	}
	return tags, nil
}

func (h *TagsClient) buildGetQuery(filter TagFilter) (string, error) {
	var query strings.Builder
	apiLink, err := h.Base.Resolve("tags_read", "tags/")
	if err != nil {
		return "", err
	}
	query.WriteString(apiLink)
	if filter.ActivityID == nil {
		if filter.TagType != nil {
			fmt.Fprintf(&query, "%s/", *filter.TagType)
		}
	} else {
		if filter.TagType == nil {
			fmt.Fprintf(&query, "%d/", *filter.ActivityID)
		} else {
			fmt.Fprintf(&query, "%d/%s/", *filter.ActivityID, *filter.TagType)
		}
	}
	return query.String(), nil
}

func (h *TagsClient) Add(item Tag) (Tag, error) {
	result, err := json.Marshal(item)
	if err != nil {
		return Tag{}, err
	}
	apiLink, err := h.Base.Resolve("tags_write", "activity/")
	if err != nil {
		return Tag{}, err
	}
	resp, err := h.Base.HTTPClient.Post(apiLink, "application/json", bytes.NewReader(result))
	if err != nil {
		return Tag{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return Tag{}, errors.New("http is not ok")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Tag{}, err
	}
	var insertedTag Tag
	err = json.Unmarshal(bodyBytes, &insertedTag)
	if err != nil {
		return Tag{}, err
	}
	return insertedTag, nil
}

func (h *TagsClient) Remove(id int) error {
	apiLink, err := h.Base.Resolve("tags_remove", fmt.Sprintf("tags/%d/", id))
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", apiLink, nil)
	if err != nil {
		return err
	}
	resp, err := h.Base.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("http is not ok")
	}
	return nil
}

package tags

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TagsClient struct {
	Client  *http.Client
	ApiLink string
}

func (h *TagsClient) Get(filter TagFilter) ([]Tag, error) {
	resp, err := h.Client.Get(h.buildGetQuery(filter))
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

func (h *TagsClient) buildGetQuery(filter TagFilter) string {
	var query strings.Builder
	query.WriteString(h.ApiLink)
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
	return query.String()
}

func (h *TagsClient) Add(item Tag) (Tag, error) {
	result, err := json.Marshal(item)
	if err != nil {
		return Tag{}, err
	}
	reader := bytes.NewReader(result)
	resp, err := h.Client.Post(h.ApiLink, "application/json", reader)
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

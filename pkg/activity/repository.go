package activity

import (
	"bytes"
	"cv-landing-cli/pkg/client"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ActivityClient struct {
	Base *client.BaseClient
}

func (h *ActivityClient) GetAllOfType(activityType string) ([]Activity, error) {
	apiLink, err := h.Base.Resolve("activity_read", fmt.Sprintf("activity/%s/", activityType))
	if err != nil {
		return []Activity{}, err
	}
	resp, err := h.Base.HTTPClient.Get(apiLink)
	if err != nil {
		return []Activity{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Activity{}, errors.New("http is not OK")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Activity{}, err
	}
	var activities []Activity
	err = json.Unmarshal(bodyBytes, &activities)
	if err != nil {
		return []Activity{}, err
	}
	return activities, nil
}

func (h *ActivityClient) Add(item Activity) (Activity, error) {
	result, err := json.Marshal(item)
	if err != nil {
		return Activity{}, err
	}
	reader := bytes.NewReader(result)
	apiLink, err := h.Base.Resolve("activity_write", "activity/")
	if err != nil {
		return Activity{}, err
	}
	resp, err := h.Base.HTTPClient.Post(apiLink, "application/json", reader)
	if err != nil {
		return Activity{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return Activity{}, errors.New("http is not ok")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Activity{}, err
	}
	var insertedActivity Activity
	err = json.Unmarshal(bodyBytes, &insertedActivity)
	if err != nil {
		return Activity{}, err
	}
	return insertedActivity, nil
}

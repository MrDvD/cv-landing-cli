package activity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ActivityClient struct {
	Client  *http.Client
	ApiLink string
}

func (h *ActivityClient) GetAll() ([]Activity, error) {
	return h.getGeneric(nil)
}

func (h *ActivityClient) GetAllOfType(activityType string) ([]Activity, error) {
	return h.getGeneric(&activityType)
}

func (h *ActivityClient) getGeneric(activityType *string) ([]Activity, error) {
	var apiLink string
	if activityType == nil {
		apiLink = h.ApiLink
	} else {
		apiLink = fmt.Sprintf("%s%s", h.ApiLink, *activityType)
	}
	resp, err := h.Client.Get(apiLink)
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
	resp, err := h.Client.Post(h.ApiLink, "application/json", reader)
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

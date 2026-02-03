package attachments

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AttachmentClient struct {
	Client  *http.Client
	ApiLink string
}

func (h *AttachmentClient) Get(activityId int) ([]Attachment, error) {
	resp, err := h.Client.Get(fmt.Sprintf("%s%d/", h.ApiLink, activityId))
	if err != nil {
		return []Attachment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Attachment{}, errors.New("http is not OK")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Attachment{}, err
	}
	var attachments []Attachment
	err = json.Unmarshal(bodyBytes, &attachments)
	if err != nil {
		return []Attachment{}, err
	}
	return attachments, nil
}

func (h *AttachmentClient) Add(item Attachment) (Attachment, error) {
	result, err := json.Marshal(item)
	if err != nil {
		return Attachment{}, err
	}
	reader := bytes.NewReader(result)
	resp, err := h.Client.Post(h.ApiLink, "application/json", reader)
	if err != nil {
		return Attachment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return Attachment{}, errors.New("http is not ok")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Attachment{}, err
	}
	var insertedAttachment Attachment
	err = json.Unmarshal(bodyBytes, &insertedAttachment)
	if err != nil {
		return Attachment{}, err
	}
	return insertedAttachment, nil
}

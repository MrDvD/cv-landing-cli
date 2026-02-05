package attachments

import (
	"bytes"
	"cv-landing-cli/pkg/client"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AttachmentClient struct {
	Base *client.BaseClient
}

func (h *AttachmentClient) Get(activityId int) ([]Attachment, error) {
	apiLink, err := h.Base.Resolve("attachments_read", fmt.Sprintf("attachments/%d/", activityId))
	if err != nil {
		return []Attachment{}, err
	}
	resp, err := h.Base.HTTPClient.Get(apiLink)
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
	apiLink, err := h.Base.Resolve("attachments_write", "attachments/")
	if err != nil {
		return Attachment{}, err
	}
	resp, err := h.Base.HTTPClient.Post(apiLink, "application/json", bytes.NewReader(result))
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

func (h *AttachmentClient) Remove(id int) error {
	apiLink, err := h.Base.Resolve("attachments_remove", fmt.Sprintf("attachments/%d/", id))
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

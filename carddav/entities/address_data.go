package entities

import (
	"encoding/xml"
	"github.com/mikelcom/caldav-go/icalendar"
	"github.com/mikelcom/caldav-go/icalendar/components"
	"github.com/mikelcom/caldav-go/utils"
	"strings"
)

type AddressData struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:carddav address-data"`
	Prop    []*CProp `xml:",omitempty"`
	Content string   `xml:",chardata"`
}

func (c *AddressData) Card() (*components.Card, error) {
	content := strings.TrimSpace(c.Content)
	if content == "" {
		return nil, utils.NewError(c.Card, "no calendar data to decode", c, nil)
	}

	cal := new(components.Card)
	err := icalendar.Unmarshal(content, cal)
	if err != nil {
		return nil, utils.NewError(c.Card, "decoding calendar data failed", c, err)
	}

	return cal, nil
}

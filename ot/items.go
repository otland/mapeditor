package ot

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"unicode/utf8"
)

type CharsetISO88591 struct {
	r   io.ByteReader
	buf *bytes.Buffer
}

func NewCharsetISO88591(r io.Reader) *CharsetISO88591 {
	return &CharsetISO88591{
		r.(io.ByteReader),
		bytes.NewBuffer(make([]byte, 0, utf8.UTFMax)),
	}
}

func (cs *CharsetISO88591) ReadByte() (b byte, err error) {
	if cs.buf.Len() == 0 {
		if r, err := cs.r.ReadByte(); err != nil {
			return 0, err
		} else if r < utf8.RuneSelf {
			return r, nil
		} else {
			cs.buf.WriteRune(rune(r))
		}
	}
	return cs.buf.ReadByte()
}

func (cs *CharsetISO88591) Read(p []byte) (int, error) {
	return 0, nil
}

func ISO88591Reader(_ string, input io.Reader) (io.Reader, error) {
	return NewCharsetISO88591(input), nil
}

// TODO: support fromid/toid
type XMLItem struct {
	Id     uint16 `xml:"id,attr"`
	FromId uint16 `xml:"fromid,attr"`
	ToId   uint16 `xml:"toid,attr"`
	Name   string `xml:"name,attr"`
}

type XMLItems struct {
	XMLName xml.Name  `xml:"items"`
	Items   []XMLItem `xml:"item"`
}

type ItemLoader struct {
	items map[uint16]string
}

func (itemLoader *ItemLoader) LoadXML(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	items := XMLItems{}

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = ISO88591Reader
	if err := decoder.Decode(&items); err != nil {
		return err
	}

	itemLoader.items = make(map[uint16]string)
	for _, v := range items.Items {
		if v.Id != 0 {
			itemLoader.items[v.Id] = v.Name
			continue
		}

		for id := v.FromId; id <= v.ToId; id++ {
			itemLoader.items[id] = v.Name
		}
	}
	return nil
}

func (itemLoader *ItemLoader) GetItemName(id uint16) (string, error) {
	if name, ok := itemLoader.items[id]; ok {
		return name, nil
	}
	return "", errors.New("item not found")
}

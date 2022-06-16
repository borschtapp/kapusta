package microdata

import (
	"log"
)

type Microdata struct {
	Items []*Item `json:"items"`
}

// addItem adds the item to the items list.
func (m *Microdata) addItem(item *Item) {
	m.Items = append(m.Items, item)
}

// GetFirstOfType returns the first item of the given type.
func (m *Microdata) GetFirstOfType(itemType ...string) *Item {
	for _, item := range m.Items {
		for _, t1 := range item.Types {
			for _, t2 := range itemType {
				if t1 == t2 {
					return item
				}
			}
		}
	}

	return nil
}

type ValueList []interface{}

type PropertyMap map[string]ValueList

type Item struct {
	Types      []string    `json:"type"`
	Properties PropertyMap `json:"properties"`
	ID         string      `json:"id,omitempty"`
}

// NewItem returns a new Item.
func NewItem() *Item {
	return &Item{
		Types:      make([]string, 0),
		Properties: make(PropertyMap, 0),
	}
}

// addType adds the value to the types list.
func (i *Item) addType(value string) {
	i.Types = append(i.Types, value)
}

// addProperty adds the property, value pair to the properties map. It appends to any existing property.
func (i *Item) addProperty(property string, value interface{}) {
	i.Properties[property] = append(i.Properties[property], value)
}

// addItem adds the property, value pair to the properties map. It appends to any existing property.
func (i *Item) addItem(property string, value *Item) {
	i.Properties[property] = append(i.Properties[property], value)
}

func (i *Item) IsOfType(itemType ...string) bool {
	for _, t1 := range i.Types {
		for _, t2 := range itemType {
			if t1 == t2 {
				return true
			}
		}
	}
	return false
}

func (i *Item) GetProperty(property string) (val interface{}, ok bool) {
	if arr, ok := i.GetProperties(property); ok {
		if len(arr) > 1 {
			log.Printf("Probably unexpected behaviour, more values of '%s' available", property)
		}

		return arr[0], true
	}
	return
}

func (i *Item) GetProperties(property string) (arr []interface{}, ok bool) {
	for _, v := range i.Properties[property] {
		arr = append(arr, v)
	}
	return arr, len(arr) > 0
}

func (i *Item) GetNestedItem(property string) (val *Item, ok bool) {
	if data, ok := i.GetNested(property); ok {
		if len(data.Items) > 1 {
			log.Printf("Probably unexpected behaviour, more values of '%s' available", property)
		}

		return data.Items[0], true
	}
	return
}

func (i *Item) GetNested(property string) (data Microdata, ok bool) {
	var arr []*Item
	for _, v := range i.Properties[property] {
		switch v.(type) {
		case *Item:
			arr = append(arr, v.(*Item))
		}
	}
	return Microdata{Items: arr}, len(arr) > 0
}

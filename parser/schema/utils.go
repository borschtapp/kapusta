package schema

import (
	"log"
	"time"

	duration "github.com/channelmeter/iso8601duration"

	"borscht.app/kapusta/microdata"
	"borscht.app/kapusta/model"
	"borscht.app/kapusta/utils"
)

func getStringOrChild(val interface{}, child ...string) (string, bool) {
	switch val.(type) {
	case string:
		return val.(string), true
	case *microdata.Item:
		item := val.(*microdata.Item)
		if text, ok := getPropertyString(item, child...); ok {
			return text, true
		} else {
			log.Printf("none of expected properties [%v] exists in (%v)\n", child, item)
		}
	default:
		log.Printf("unable to process `%s`, unexpected type `%T`\n", val, val)
	}

	return "", false
}

func getPropertyStringOrChild(item *microdata.Item, key string, child ...string) (string, bool) {
	if val, ok := item.GetProperty(key); ok {
		return getStringOrChild(val, child...)
	}

	return "", false
}

func getPropertyString(item *microdata.Item, key ...string) (string, bool) {
	if val, ok := item.GetProperty(key...); ok {
		if text, ok := val.(string); ok {
			return text, true
		} else {
			log.Printf("unable to retrieve `string` value of `%s` in (%v)\n", key, item)
		}
	}

	return "", false
}

func getPropertyInt(item *microdata.Item, key ...string) (int, bool) {
	if val, ok := item.GetProperty(key...); ok {
		return utils.FindInt(val), true
	}

	return 0, false
}

func getPropertyFloat(item *microdata.Item, key ...string) (float32, bool) {
	if val, ok := item.GetProperty(key...); ok {
		return utils.FindFloat32(val), true
	}

	return 0, false
}

func getPropertyDuration(item *microdata.Item, key string) (time.Duration, bool) {
	if val, ok := getPropertyStringOrChild(item, key, "maxValue", "minValue"); ok {
		d, err := duration.FromString(val)
		if err != nil {
			log.Printf("unable to parse duration `%s` in (%v)\n", val, item)
		} else {
			return d.ToDuration(), true
		}
	}

	return 0, false
}

func parseInstructionSteps(item *microdata.Item) model.Step {
	var instr model.Step
	if val, ok := getPropertyString(item, "text", "description"); ok {
		instr.Text = val
	}
	if val, ok := getPropertyString(item, "name"); ok && val != instr.Text {
		instr.Name = val
	}
	if val, ok := getPropertyStringOrChild(item, "image", "url"); ok {
		instr.Image = val
	}
	if val, ok := getPropertyString(item, "url"); ok {
		instr.Url = val
	}

	return instr
}

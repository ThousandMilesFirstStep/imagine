package domain

func setExportOptions(image Image, conf map[string]interface{}) error {
	exportOptions := &ExportOptions{}

	quality := conf["quality"]
	if quality != nil {
		exportOptions.Quality = int(quality.(int64))
	}

	format := conf["format"]
	if format != nil {
		exportOptions.Format = format.(string)
	}

	strip := conf["strip"]
	if strip != nil {
		exportOptions.Strip = strip.(bool)
	}

	interlace := conf["interlace"]
	if interlace != nil {
		exportOptions.Interlace = interlace.(bool)
	}

	lossless := conf["lossless"]
	if lossless != nil {
		exportOptions.Lossless = lossless.(bool)
	}

	effort := conf["effort"]
	if effort != nil {
		exportOptions.Effort = effort.(int)
	}

	image.SetExportOptions(exportOptions)

	return nil
}

func watermark(image Image, conf map[string]interface{}) error {
	position := TopRight
	if conf["position"] != nil {
		position = Position(conf["position"].(string))
	}

	padding := 0
	if conf["padding"] != nil {
		padding = int(conf["padding"].(int64))
	}

	return image.Watermark(conf["image"].(string), position, padding)
}

func strip(image Image, conf map[string]interface{}) error {
	return image.Strip()
}

func autorotate(image Image, conf map[string]interface{}) error {
	return image.AutoRotate()
}

func thumbnail(image Image, conf map[string]interface{}) error {
	width := int(conf["width"].(int64))
	height := int(conf["height"].(int64))
	inset := conf["inset"].(bool)

	return image.Thumbnail(width, height, inset)
}

func flatten(image Image, conf map[string]interface{}) error {
	configColor := conf["backgroundColor"].(string)

	color, err := NewColorFromHex(configColor)
	if err != nil {
		return err
	}

	return image.Flatten(color)
}

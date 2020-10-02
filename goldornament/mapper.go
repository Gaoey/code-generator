package goldornament

func goldornamentRequestToModel(req Request) models.GoldOrnamentAsset {
	result := models.GoldOrnamentAsset{
		Name: req.Name,
	}

	return result
}

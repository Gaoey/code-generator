package test

func testRequestToModel(req Request) models.Test {
	result := models.Test{
		Name: req.Name,
	}

	return result
}

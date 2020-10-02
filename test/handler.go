package test

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAll(c echo.Context) error {
	var (
		offsetParam = c.QueryParam("offset")
		limitParam  = c.QueryParam("limit")
	)
	offset, _ := strconv.Atoi(offsetParam)
	limit, _ := strconv.Atoi(limitParam)
	if limit == 0 {
		limit = 10
	}

	userId, err := jwt.GetUserID(c)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			models.NewTestErrorResponse(http.StatusNotFound, errs.MismatchToken, err.Error()))
	}

	res, err := h.service.GetAll(offset, limit, userId)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.NewTestErrorResponse(http.StatusInternalServerError, errs.InternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, models.NewTestResponse(http.StatusOK, res))
}

func (h *Handler) GetById(c echo.Context) error {
	_, err := jwt.GetUserID(c)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			models.NewTestErrorResponse(http.StatusNotFound, errs.MismatchToken, err.Error()))
	}

	id := c.Param("id")
	res, err := h.service.GetById(id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.NewTestErrorResponse(http.StatusInternalServerError, errs.InternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, models.NewTestResponse(http.StatusOK, &res))
}

func (h *Handler) Add(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.NewTestErrorResponse(http.StatusBadRequest, errs.BadRequestError, err.Error()))
	}

	userId, err := jwt.GetUserID(c)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			models.NewTestErrorResponse(http.StatusNotFound, errs.MismatchToken, err.Error()))
	}

	response, err := h.service.Add(req, userId)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.NewTestErrorResponse(http.StatusInternalServerError, errs.InternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, models.NewTestResponse(http.StatusOK, response))
}

func (h *Handler) Update(c echo.Context) error {
	id := c.Param("id")
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.NewTestErrorResponse(http.StatusBadRequest, errs.BadRequestError, err.Error()))
	}

	_, err := jwt.GetUserID(c)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			models.NewTestErrorResponse(http.StatusNotFound, errs.MismatchToken, err.Error()))
	}

	err = h.service.Update(id, req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.NewTestErrorResponse(http.StatusInternalServerError, errs.InternalServerError, err.Error()))
	}

	return c.JSON(http.StatusNoContent, models.NewTestResponse(http.StatusOK, nil))
}

func (h *Handler) Remove(c echo.Context) error {
	_, err := jwt.GetUserID(c)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			models.NewTestErrorResponse(http.StatusNotFound, errs.MismatchToken, err.Error()))
	}

	id := c.Param("id")
	err = h.service.Remove(id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.NewTestErrorResponse(http.StatusInternalServerError, errs.InternalServerError, err.Error()))
	}

	return c.JSON(http.StatusNoContent, models.NewTestResponse(http.StatusOK, nil))
}

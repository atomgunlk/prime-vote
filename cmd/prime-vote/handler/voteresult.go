package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func (h *Handler) GetVoteResult(c *fiber.Ctx) error {
	req := new(model.GetVoteResultRequest)

	response, err := h.Deps.Repo.GetVoteResult(req)
	if err != nil {
		logger.WithError(err).Error("[GetVoteResult]: repo.GetVoteResult failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}

func (h *Handler) ExportVoteResult(c *fiber.Ctx) error {
	// req := new(model.ExportVoteResultRequest)
	response, err := h.Deps.Repo.GetVoteResult(&model.GetVoteResultRequest{})
	if err != nil {
		logger.WithError(err).Error("[ExportVoteResult]: repo.GetVoteResult failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	filepath, err := writeVoteResultToExcel(response.Items)
	if err != nil {
		logger.WithError(err).Error("[ExportVoteResult]: writeVoteResultToExcel failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	defer os.Remove(filepath)

	bkk, _ := time.LoadLocation("Asia/Bangkok")
	c.Attachment(fmt.Sprintf("vote-result-%s.xlsx", time.Now().In(bkk).Format("20060102")))
	return c.SendFile(filepath)
}

func writeVoteResultToExcel(results []model.VoteItem) (string, error) {
	xlsLocalFilePath := fmt.Sprintf("/tmp/%s.xlsx", uuid.NewString())
	err := os.MkdirAll(filepath.Dir(xlsLocalFilePath), 0755)
	if err != nil {
		return "", err
	}
	// Gen XLS File
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// Header
	headers := map[string]string{
		"A1": "ID",
		"B1": "Name",
		"C1": "Description",
		"D1": "Vote count",
	}
	for k, v := range headers {
		f.SetCellValue(sheetName, k, v)
	}

	// Data
	var data map[string]interface{}
	row := 2
	for _, result := range results {
		var cell []string
		for j, c := 0, 'A'; c <= 'Z'; j, c = j+1, c+1 {
			cell = append(cell, fmt.Sprintf("%c%d", c, row))
		}

		data = map[string]interface{}{
			cell[0]: result.ID,
			cell[1]: result.Name,
			cell[2]: result.Description,
			cell[3]: result.VoteCount,
		}
		for k, v := range data {
			f.SetCellValue(sheetName, k, v)
		}
		row = row + 1
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs(xlsLocalFilePath); err != nil {
		return "", err
	}
	return xlsLocalFilePath, nil
}

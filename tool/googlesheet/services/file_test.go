package services

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/getcouragenow/bootstrap/tool/googlesheet/services/config"
)

var (
	sheet         = "TEST SHEET"
	absPath       = ""
	csvFilePathA  = ""
	csvFilePathB  = ""
	csvFilePathC  = ""
	outputsFolder = "content"
)

var sheetA = config.Config{
	ID:        "1VvNEPpTzTJbYvSu9SRvc9BmeYVtSprQmmwqY3CAK7gM",
	URL:       "https://docs.google.com/spreadsheets/d/1VvNEPpTzTJbYvSu9SRvc9BmeYVtSprQmmwqY3CAK7gM",
	CSV:       "https://docs.google.com/spreadsheets/d/e/2PACX-1vSqNo3Li7zrKSqIRDbykSipeT4RtsyrAM0iL3WoWdr6XnPtOf9NsYkV2xYV5N0dR4uNM7YeZyXXiuTT/pub?output=csv",
	Merge:     "cell",
	OutDir:    "outputs/content/XXX/",
	FileName:  "",
	Extension: "",
}

var sheetB = config.Config{
	ID:        "1y6TdPrs-iUrHZ8XDv7PoCdBa5uimSHwmgNj4dhZs5Lc",
	URL:       "https://docs.google.com/spreadsheets/d/1y6TdPrs-iUrHZ8XDv7PoCdBa5uimSHwmgNj4dhZs5Lc",
	CSV:       "https://docs.google.com/spreadsheets/d/e/2PACX-1vQtmjheqdY7msmINYw4h8w4E-Sx1XLhtRTvBV4KohHumkF3JkyEFMhdJG2qyICYRujMxq5ZGiQ37RF1/pub?output=csv",
	Merge:     "row",
	OutDir:    "outputs/config/_default/",
	FileName:  "",
	Extension: "",
}

var sheetC = config.Config{
	ID:        "1jtFj-gEo6Gyyy1pPqY0eNtBDzD-1MlEoTZNTA-syO6g",
	URL:       "https://docs.google.com/spreadsheets/d/1jtFj-gEo6Gyyy1pPqY0eNtBDzD-1MlEoTZNTA-syO6g",
	CSV:       "https://docs.google.com/spreadsheets/d/e/2PACX-1vSbABjbU7HOT9kKDiNFCvsSpoxf0l9Mv5DKqoMWaHbGliSUVKAgSiyJcjbcRSQNhtpbM3s1RPQbxyH3/pub?output=csv",
	Merge:     "column",
	OutDir:    "outputs/i18n/",
	FileName:  "",
	Extension: ".toml",
}

func TestAll(t *testing.T) {
	download()
	getAbsPath()
	t.Run("Write JSON Files", testWriteLanguageFiles)
	t.Run("Write JSON dump Files", testWriteDataDumpFiles)
	t.Run("Write hugo files", testWriteHugoFiles)
	cleanTestData()
}

// Test write languages in json file
func testWriteLanguageFiles(t *testing.T) {
	jsonDirPath := path.Join(absPath, outputsFolder)
	err := WriteLanguageFiles(csvFilePathA, jsonDirPath, sheet)

	if err != nil {
		t.Error("Error to Write files ====>", err)
		return
	}
}

// Test write dump data in json file
func testWriteDataDumpFiles(t *testing.T) {
	jsonDirPath := path.Join(absPath, outputsFolder)
	err := WriteDataDumpFiles(csvFilePathA, jsonDirPath, "dump")

	if err != nil {
		t.Error("Error to Write files ====>", err)
		return
	}
}

// Test generate hugo content
func testWriteHugoFiles(t *testing.T) {

	//
	if err := WriteHugoFiles(csvFilePathA, sheetA, "", ""); err != nil {
		t.Error("Error ====>", csvFilePathA, err)
	}

	if err := WriteHugoFiles(csvFilePathB, sheetB, "", ""); err != nil {
		t.Error("Error ====>", csvFilePathB, err)
	}

	if err := WriteHugoFiles(csvFilePathC, sheetC, "", ""); err != nil {
		t.Error("Error ====>", csvFilePathC, err)
	}
}

// download csv files
func download() {
	csvFilePathA = path.Join(absPath, "sheet_a.csv")
	csvFilePathB = path.Join(absPath, "sheet_b.csv")
	csvFilePathC = path.Join(absPath, "sheet_c.csv")

	if err := Download(sheetA.CSV, csvFilePathA, 5000, sheet, true); err != nil {
		log.Fatal("Error to download:", err)
	}

	if err := Download(sheetB.CSV, csvFilePathB, 5000, sheet, true); err != nil {
		log.Fatal("Error to download:", err)
	}

	if err := Download(sheetC.CSV, csvFilePathC, 5000, sheet, true); err != nil {
		log.Fatal("Error to download:", err)
	}
}

func getAbsPath() {
	absPath, err := GetAbsoluteFilePath("", sheet)
	if err != nil {
		log.Fatal("Error to get Abs path:", err)
	}
	absPath = absPath
}

func cleanTestData() {
	os.RemoveAll(path.Join(absPath, outputsFolder))
	os.Remove(path.Join(absPath, "sheet_a.csv"))
	os.Remove(path.Join(absPath, "sheet_b.csv"))
	os.Remove(path.Join(absPath, "sheet_c.csv"))
}

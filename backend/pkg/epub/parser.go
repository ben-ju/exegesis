package epub

import (
	"log"
	"path/filepath"

	"github.com/ben-ju/exegesis/internal/domain"
	"github.com/ben-ju/exegesis/internal/repository"
	"github.com/taylorskalyo/goreader/epub"
)

func ParseEPUB(filePath string, repo repository.BibleRepository) error {
	rc, err := epub.OpenReader(filePath)
	if err != nil {
		return err
	}
	defer rc.Close()

	book := rc.Rootfiles[0]
	version := filepath.Base(filePath)
	version = version[:len(version)-len(filepath.Ext(version))]

	bibleVersion := domain.BibleVersion{
		Title:    book.Title,
		Version:  version,
		FilePath: filePath,
	}

	if err := repo.AddVersion(bibleVersion); err != nil {
		return err
	}

	// Parse EPUB content and add verses
	// This is a simplified example. You'll need to implement the actual parsing logic.
	for _, item := range book.Manifest.Items {
		if item.MediaType == "application/xhtml+xml" {
			content, err := item.Open()
			if err != nil {
				return err
			}
			log.Printf("book content : %v\n", content)
			// Parse the content and extract verses
			// For each verse:
			// repo.AddVerse(version, domain.Verse{OsisID: "...", Content: "..."})
		}
	}

	return nil
}

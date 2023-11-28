package valueObject

import "testing"

func TestUnixFilePath(t *testing.T) {
	t.Run("ValidUnixFilePath", func(t *testing.T) {
		validUnixFilePaths := []string{
			"/",
			"/root",
			"/root/",
			"/home/sandbox/file.php",
			"/home/sandbox/file",
			"/home/sandbox/file with space.php",
			"/home/100sandbox/file.php",
			"/home/100sandbox/Imagem - Sem Título.jpg",
			"/home/100sandbox/Imagem - Sem Título & BW.jpg",
			"/home/100sandbox/Imagem - Sem Título # BW.jpg",
			"/home/@directory/file.gif",
			"file.php",
		}
		for _, filePath := range validUnixFilePaths {
			_, err := NewUnixFilePath(filePath)
			if err != nil {
				t.Errorf("Expected no error for %s, got %v", filePath, err)
			}
		}
	})

	t.Run("ValidUnixFilePath", func(t *testing.T) {
		invalidUnixFilePaths := []string{
			"",
			"/home/user/file.php?blabla",
			"/home/sandbox/domains/@<php52.sandbox.ntorga.com>",
			"../file.php",
			"./file.php",
			"/home/../file.php",
		}
		for _, filePath := range invalidUnixFilePaths {
			_, err := NewUnixFilePath(filePath)
			if err == nil {
				t.Errorf("Expected error for %s, got nil", filePath)
			}
		}
	})
}

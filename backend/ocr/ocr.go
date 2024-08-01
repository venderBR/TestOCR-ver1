package ocr

import (
    "fmt"
    "os/exec"
)

func RunOCR(pdfPath string) (string, error) {
    cmd := exec.Command("python", "C:\\Users\\msi01\\Desktop\\TestOCR\\scripts\\ocr_script.py", pdfPath)
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to run OCR script: %v", err)
    }
    return string(output), nil
}

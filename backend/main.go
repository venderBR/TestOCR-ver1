package main

import (
    "backend/ocr"
    "fmt"
    "log"
    "net/http"
)

func ocrHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
    pdfPath := "C:/Users/msi01/Desktop/TestOCR/non-text-searchable.pdf"

    // Run OCR
    ocrOutput, err := ocr.RunOCR(pdfPath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    

    // // Run Spell Check
    // spellCheckedOutput, err := ocr.SpellCheck(ocrOutput)
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }

    fmt.Fprintf(w, "OCR and Spell Checked Output: %s", ocrOutput)
}

func main() {
    http.HandleFunc("/ocr", ocrHandler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

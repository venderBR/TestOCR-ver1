import fitz  # PyMuPDF
import pytesseract
from PIL import Image
import io

def pdf_to_text(pdf_path):
    try:
        # Open the PDF file
        pdf_document = fitz.open(pdf_path)
        text = ""

        # Process each page
        for page_number in range(len(pdf_document)):
            print(f"Processing page {page_number + 1}...")
            page = pdf_document.load_page(page_number)
            pix = page.get_pixmap()
            
            # Convert the Pixmap to PIL Image
            img = Image.open(io.BytesIO(pix.tobytes()))
            
            # Perform OCR
            text += pytesseract.image_to_string(img)

        return text

    except Exception as e:
        print(f"An error occurred: {e}")
        return None

if __name__ == "__main__":
    pdf_path = 'C:/Users/msi01/Desktop/TestOCR/non-text-searchable.pdf'  # Replace with your PDF file path
    text = pdf_to_text(pdf_path)
    
    if text:
        print(text)
        # with open('output.txt', 'w', encoding='utf-8') as f:
        #     f.write(text)
        # print("Text extraction complete. Check 'output.txt'.")
    else:
        print("Failed to extract text from the PDF.")

import os
import re
from typing import Optional
import psycopg2

# from database import add_verses
from enums.books import BOOKS_FR
from ebooklib import epub
from bs4 import BeautifulSoup
from dotenv import load_dotenv
import warnings

from tests.books_coverage import test_book_coverage
# Ignorer les UserWarnings spécifiques d'ebooklib
warnings.filterwarnings("ignore", category=UserWarning, module="ebooklib.epub")

# Ignorer les FutureWarnings spécifiques d'ebooklib
warnings.filterwarnings("ignore", category=FutureWarning, module="ebooklib.epub")

# Ignorer les warnings XML/HTML de BeautifulSoup
warnings.filterwarnings("ignore", category=UserWarning, module="html.parser")

# Load environment variables from the parent directory
dotenv_path = os.path.join(os.path.dirname(os.path.dirname(__file__)), '.env')
load_dotenv(dotenv_path)

books_pattern = "|".join(re.escape(book) for book in BOOKS_FR)
header_re = re.compile(r"^(?P<book>" + books_pattern + r")\s*(?P<chapter>\d+)?\s*$")
verse_re = re.compile(r"^(\d+)$")
def get_db_connection():
    return psycopg2.connect(
        host=os.getenv("POSTGRES_HOST"),
        database=os.getenv("POSTGRES_DB"),
        user=os.getenv("POSTGRES_USER"),
        password=os.getenv("POSTGRES_PASSWORD")
    )

def parse_epub_metadata(book: epub.EpubBook) -> dict:
    """Extract and structure EPUB metadata"""
    metadata = {
        "title": safe_filename(book.get_metadata('DC', 'title')[0][0]) if book.get_metadata('DC', 'title') else None,
        "author": book.get_metadata('DC', 'creator')[0][0] if book.get_metadata('DC', 'creator') else "Unknown Author",
        "language": book.get_metadata('DC', 'language')[0][0] if book.get_metadata('DC', 'language') else "Unknown Language",
        "publisher": book.get_metadata('DC', 'publisher')[0][0] if book.get_metadata('DC', 'publisher') else "Unknown Publisher",
        "date": book.get_metadata('DC', 'date')[0][0] if book.get_metadata('DC', 'date') else "Unknown Date"
    }
    return metadata
def safe_filename(title: str) -> str:
    """Sanitize filename by removing invalid characters"""
    return re.sub(r'[\\/*?:"<>|]', "", title).strip().replace(" ", "_")
def flatten_text_sections(book, metadata):
    print("flattening")
    """
    Lit un fichier EPUB et crée un unique fichier texte
    contenant le contenu brut de tous les documents dans l'ordre de lecture,
    SANS balises HTML (uniquement le texte).
    """
    # Get paths from environment variable
    resources_path = "resources/"
    # Create flattened directory
    base_dir = os.path.dirname(resources_path)
    flatten_dir = os.path.join(base_dir, "flattened")
    os.makedirs(flatten_dir, exist_ok=True)

    # Create output filename
    original_filename = os.path.basename(resources_path)
    output_path = os.path.join(flatten_dir, f"{metadata["title"]}_flatten.txt")

    # Process book content
    all_text_parts = []
    for itemref in book.spine:
        item_id = itemref[0]
        item = book.get_item_with_id(item_id)

        if item and item.get_name():
            try:
                html_content = item.get_content().decode("utf-8", errors="ignore")
                soup = BeautifulSoup(html_content, "html.parser")
                text_content = soup.get_text(separator="\n")
                all_text_parts.append(text_content.strip())
            except Exception as e:
                print(f"Error processing {item.get_name()}: {str(e)}")

    # Write output file
    with open(output_path, "w", encoding="utf-8") as f:
        f.write("\n\n".join(all_text_parts))

    return output_path


def parse(flattened_book_path):
    """
    Parse l'EPUB de la Bible et retourne une liste de tuples (book, chapter, verse, text).
    Le parser conserve l'état courant et concatène les lignes non numérotées aux versets précédents.
    """
    verses = []
    current_book = None
    current_chapter = None
    current_verse = None
    current_text = ""

    with open(flattened_book_path, "r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            header_match = header_re.match(line)
            if header_match:
                if current_verse is not None and current_text:
                    verses.append((current_book, current_chapter, current_verse, current_text.strip()))
                    current_verse = None
                    current_text = ""
                current_book = header_match.group("book")
                current_chapter = header_match.group("chapter") or "1"
                continue

            verse_match = verse_re.match(line)
            if verse_match:
                if current_verse is not None:
                    verses.append((current_book, current_chapter, current_verse, current_text.strip()))
                current_verse = verse_match.group(1)
                current_text = ""
                continue

            if current_text:
                current_text += " " + line
            else:
                current_text = line

    if current_verse is not None and current_text:
        verses.append((current_book, current_chapter, current_verse, current_text.strip()))

    return verses

def open_epub(epub_path) -> Optional[epub.EpubBook] | Exception:
    """Charge et retourne l'objet ebooklib du fichier EPUB."""
    print("open epub")
    try:
        book = epub.read_epub(epub_path)
        print("epub opened")
        return book
    except Exception as e:
        print("error in epub opening")
        print(e)
        return e


def main(book_path):
    try:
        book = open_epub(book_path)
        if not book:
            return
        metadata = parse_epub_metadata(book)
        flattened_path = flatten_text_sections(book, metadata)
        verses = parse(flattened_path)
        test_book_coverage(verses)
    except Exception as e:
        print(e)


if __name__ == "__main__":
    main("./resources/lsg.epub")


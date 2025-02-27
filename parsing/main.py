
import os
import re
from typing import Optional

from database import add_verses
from enums.parser import BOOKS_FR
from ebooklib import epub
from bs4 import BeautifulSoup


books_pattern = "|".join(re.escape(book) for book in BOOKS_FR)
header_re = re.compile(r"^(?P<book>" + books_pattern + r")\s*(?P<chapter>\d+)?\s*$")
verse_re = re.compile(r"^(\d+)$")

def parse_epub_metadata(book: epub.EpubBook) -> dict[str, dict[str, str]]:
    """
    Exemple d'usage: ouvre l'EPUB, récupère les sections,
    renvoie une structure de base (metadata + list de (title, text)).
    """
    title = book.get_metadata('DC', 'title')
    authors = book.get_metadata('DC', 'creator')
    language = book.get_metadata('DC', 'language')
    publisher = book.get_metadata('DC', 'publisher')
    date = book.get_metadata('DC', 'date')

    metadata = {}
    if title:
        metadata["title"] = title
    if authors:
        metadata["authors"] = authors
    if language:
        metadata["language"] = language
    if publisher:
        metadata["publisher"] = publisher
    if date:
        metadata["date"] = date

    return {
        "metadata": metadata,
    }
def flatten_text_sections(book):
    """
    Lit un fichier EPUB et crée un unique fichier texte
    contenant le contenu brut de tous les documents dans l'ordre de lecture,
    SANS balises HTML (uniquement le texte).
    """
    # TODO : Add specific book flattening by removing indexing [0]
    base_name = os.path.splitext(os.path.basename(os.getenv("RESOURCES_PATH")))[0]
    flatten_dir = os.path.join(os.path.dirname(os.getenv("RESOURCES_PATH"), "flatten"))
    os.makedirs(flatten_dir, exist_ok=True)
    output_path = os.path.join(flatten_dir, f"{base_name}_flatten.txt")

    all_text_parts = []
    for itemref in book.spine:
        item_id = itemref[0]
        item = book.get_item_with_id(item_id)

        if item is not None and item.get_name():
            html_content = item.get_content().decode("utf-8", errors="ignore")
            soup = BeautifulSoup(html_content, "html.parser")
            text_content = soup.get_text(separator="\n")
            all_text_parts.append(text_content)

    with open(output_path, "w", encoding="utf-8") as f:
        f.write("\n".join(all_text_parts))
        f.write("\n")

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

def open_epub(epub_path) -> Optional[epub.EpubBook]:
    """Charge et retourne l'objet ebooklib du fichier EPUB."""
    try:
        book = epub.read_epub(epub_path)
        return book
    except Exception as e:
        return e

def main(book_path = ""):
    book = open_epub(book_path)
    if not book:
        return
    # TODO : insert metadata to database
    metadata = parse_epub_metadata(book)
    flattened_path = flatten_text_sections(book)
    verses = parse(flattened_path)
    add_verses(verses)

if __name__ == "__main__":
    main()

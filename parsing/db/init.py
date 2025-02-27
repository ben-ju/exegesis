import sqlite3
import os

from enums.parser import FRENCH_TO_ENGLISH, SEED_BOOKS
from dotenv import load_dotenv
load_dotenv()

def init_db():
    """Initialise la base de données SQLite avec le schéma défini."""
    conn = sqlite3.connect(os.getenv("DATABASE"))
    cursor = conn.cursor()

    cursor.executescript("""
    -- Table Category
    CREATE TABLE IF NOT EXISTS category (
        id            INTEGER PRIMARY KEY AUTOINCREMENT,
        title         TEXT    NOT NULL,
        desc          TEXT,
        abbreviation  TEXT
    );

    -- Table bible_books
    CREATE TABLE IF NOT EXISTS bible_books (
        id                  INTEGER PRIMARY KEY AUTOINCREMENT,
        title               TEXT NOT NULL,
        abbreviation        TEXT NOT NULL,
        is_deuterocanonical BOOLEAN DEFAULT 0,
        is_old_testament    BOOLEAN DEFAULT 0,
        is_new_testament    BOOLEAN DEFAULT 0
    );

    -- Table books
    CREATE TABLE IF NOT EXISTS books (
        id           INTEGER PRIMARY KEY AUTOINCREMENT,
        title        TEXT    NOT NULL,
        abbreviation TEXT,
        language     TEXT,
        authors      TEXT,
        cover        TEXT,
        category_id  INTEGER,
        FOREIGN KEY (category_id) REFERENCES category(id)
    );

    -- Table chapters
    CREATE TABLE IF NOT EXISTS chapters (
        id              INTEGER PRIMARY KEY AUTOINCREMENT,
        bible_book_id   INTEGER NOT NULL,
        number          INTEGER NOT NULL,
        is_ambiguous    BOOLEAN DEFAULT 0,
        FOREIGN KEY (bible_book_id) REFERENCES bible_books(id)
    );

    -- Table verses
    CREATE TABLE IF NOT EXISTS verses (
        id           INTEGER PRIMARY KEY AUTOINCREMENT,
        chapter_id   INTEGER NOT NULL,
        number       INTEGER NOT NULL,
        is_ambiguous BOOLEAN DEFAULT 0,
        FOREIGN KEY (chapter_id) REFERENCES chapters(id)
    );

    -- Table contents
    CREATE TABLE IF NOT EXISTS contents (
        id             INTEGER PRIMARY KEY AUTOINCREMENT,
        book_id        INTEGER NOT NULL,
        start_verse_id INTEGER,
        end_verse_id   INTEGER,
        text           TEXT,
        FOREIGN KEY (book_id)        REFERENCES books(id),
        FOREIGN KEY (start_verse_id) REFERENCES verses(id),
        FOREIGN KEY (end_verse_id)   REFERENCES verses(id)
    );
    """)

    for book in SEED_BOOKS:
        cursor.execute("""
            INSERT INTO bible_books (title, abbreviation, is_deuterocanonical, is_old_testament, is_new_testament)
            VALUES (?, ?, ?, ?, ?)
        """, (
            book["title"],
            book["abbreviation"],
            book["is_deuterocanonical"],
            book["is_old_testament"],
            book["is_new_testament"]
        ))


    conn.commit()
    conn.close()

def populate_bible_version(verses,
                           bible_version_title,
                           bible_version_abbreviation,
                           language, authors="Unknown", cover=None, category_id=None):
    """
    Inserts a new Bible version (e.g. LSG) into the 'books' table and for each parsed verse
    (a tuple: (french_book, chapter, verse, text)), it creates or reuses the canonical chapter
    and verse records (linked to bible_books) and inserts the text into 'contents' linked to this Bible version.
    """
    conn = sqlite3.connect(os.getenv("DATABASE"))
    cursor = conn.cursor()

    cursor.execute("SELECT id FROM books WHERE title = ?", (bible_version_title,))
    row = cursor.fetchone()
    if row:
        bible_version_id = row[0]
    else:
        cursor.execute("""
            INSERT INTO books (title, abbreviation, language, authors, cover, category_id)
            VALUES (?, ?, ?, ?, ?, ?)
        """, (bible_version_title, bible_version_abbreviation, language, authors, cover, category_id))
        bible_version_id = cursor.lastrowid

    for french_book, chapter_num, verse_num, text in verses:
        if french_book not in FRENCH_TO_ENGLISH:
            continue
        english_title = FRENCH_TO_ENGLISH[french_book]

        cursor.execute("SELECT id FROM bible_books WHERE title = ?", (english_title,))
        row = cursor.fetchone()
        if not row:
            continue
        bible_book_id = row[0]

        cursor.execute("""
            SELECT id FROM chapters
            WHERE bible_book_id = ? AND number = ?
        """, (bible_book_id, chapter_num))
        row = cursor.fetchone()
        if row:
            chapter_id = row[0]
        else:
            cursor.execute("""
                INSERT INTO chapters (bible_book_id, number)
                VALUES (?, ?)
            """, (bible_book_id, chapter_num))
            chapter_id = cursor.lastrowid

        cursor.execute("""
            SELECT id FROM verses
            WHERE chapter_id = ? AND number = ?
        """, (chapter_id, verse_num))
        row = cursor.fetchone()
        if row:
            verse_id = row[0]
        else:
            cursor.execute("""
                INSERT INTO verses (chapter_id, number)
                VALUES (?, ?)
            """, (chapter_id, verse_num))
            verse_id = cursor.lastrowid

        cursor.execute("""
            INSERT INTO contents (book_id, start_verse_id, end_verse_id, text)
            VALUES (?, ?, ?, ?)
        """, (bible_version_id, verse_id, verse_id, text))

    conn.commit()
    conn.close()


def get_verse_contents(book_title, chapter_number, verse_number):
    """
    Retrieve the content of every Bible version that contains the given reference.

    Parameters:
      - book_title: Canonical Bible book title in English (e.g. "Colossians").
      - chapter_number: Chapter number (e.g. 3) as int or string.
      - verse_number: Verse number (e.g. 16) as int or string.

    Returns:
      A list of dictionaries where each dictionary contains:
         - 'bible_version': The title of the Bible version (from the 'books' table)
         - 'text': The verse text stored in the 'contents' table.
    """
    conn = sqlite3.connect(os.getenv("DATABASE"))
    conn.row_factory = sqlite3.Row  
    cursor = conn.cursor()

    cursor.execute("SELECT id FROM bible_books WHERE title = ?", (book_title,))
    row = cursor.fetchone()
    if not row:
        conn.close()
        return []
    bible_book_id = row["id"]

    cursor.execute("SELECT id FROM chapters WHERE bible_book_id = ? AND number = ?", (bible_book_id, chapter_number))
    row = cursor.fetchone()
    if not row:
        conn.close()
        return []
    chapter_id = row["id"]

    cursor.execute("SELECT id FROM verses WHERE chapter_id = ? AND number = ?", (chapter_id, verse_number))
    row = cursor.fetchone()
    if not row:
        conn.close()
        return []
    verse_id = row["id"]

    query = """
        SELECT b.title AS bible_version, c.text
        FROM contents c
        JOIN books b ON c.book_id = b.id
        WHERE c.start_verse_id = ?
    """
    cursor.execute(query, (verse_id,))
    results = cursor.fetchall()
    conn.close()

    return [dict(row) for row in results]
if __name__ == "__main__":
    init_db()

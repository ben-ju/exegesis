import os
import sqlite3
from enums.parser import FRENCH_TO_ENGLISH
from dotenv import load_dotenv
load_dotenv()

def add_verses(verses,
                           bible_version_title,
                           bible_version_abbreviation,
                           language, authors="Unknown", cover=None, category_id=None):
    """
    Inserts a new Bible version (e.g. LSG) into the 'books' table and for each parsed verse
    (a tuple: (french_book, chapter, verse, text)), it creates or reuses the canonical chapter
    and verse records (linked to bible_books) and inserts the text into 'contents' linked to this Bible version.
    """
    conn = sqlite3.connect(os.getenv(("DATABASE")))
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

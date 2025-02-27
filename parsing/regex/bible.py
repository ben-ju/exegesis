import re
from enums.parser import BOOKS_FR

books_pattern = "|".join(re.escape(book) for book in BOOKS_FR)
header_re = re.compile(r"^(?P<book>" + books_pattern + r")\s*(?P<chapter>\d+)?\s*$")
verse_re = re.compile(r"^(\d+)$")

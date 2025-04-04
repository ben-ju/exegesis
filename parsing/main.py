import os
import re
# import psycopg2
import warnings
import importlib.util
import sys

from typing import Optional
from enums.books import BOOKS_FR
from ebooklib import epub
from dotenv import load_dotenv
# from tests.books_coverage import test_book_coverage

# Ignorer les UserWarnings spécifiques d'ebooklib
warnings.filterwarnings("ignore", category=UserWarning, module="ebooklib.epub")
# Ignorer les FutureWarnings spécifiques d'ebooklib
warnings.filterwarnings("ignore", category=FutureWarning, module="ebooklib.epub")
# Ignorer les warnings XML/HTML de BeautifulSoup
warnings.filterwarnings("ignore", category=UserWarning, module="html.parser")

# Load environment variables from the parent directory
dotenv_path = os.path.join(os.path.dirname(os.path.dirname(__file__)), '.env')
load_dotenv(dotenv_path)


# TODO : ADD ENV FOR PARSERS
RESOURCES_PATH = os.getenv("RESOURCES_PATH") or "resources/"
FLATTENED_PATH=os.getenv("FLATTENED_PATH") or "resources/flattened/"
PARSERS_PATH=os.getenv("PARSER_PATH") or "parsers/"

books_pattern = "|".join(re.escape(book) for book in BOOKS_FR)
header_re = re.compile(r"^(?P<book>" + books_pattern + r")\s*(?P<chapter>\d+)?\s*$")
verse_re = re.compile(r"^(\d+)$")

def load_parser(parser_filename):
    """
    Dynamically load a parser module from its filename.
    Args:
        parser_filename (str): The filename of the parser (with .py extension)
    Returns:
        module: The imported parser module
    """
    # Get the full path to the parser file
    parser_path = os.path.join(PARSERS_PATH, parser_filename)
    # Extract module name without extension
    module_name = os.path.splitext(parser_filename)[0]
    # Load the module specification
    spec = importlib.util.spec_from_file_location(module_name, parser_path)
    if spec is None:
        raise ImportError(f"Could not find the parser: {parser_path}")
    # Create the module and initialize it
    parser_module = importlib.util.module_from_spec(spec)
    sys.modules[module_name] = parser_module
    spec.loader.exec_module(parser_module)
    return parser_module

# Create the flattened directory inside the resources folder containing already flattened resources
def create_flattened_dir():
    # Get paths from environment variable
    resources_path = "resources/"
    # Create flattened directory
    base_dir = os.path.dirname(resources_path)
    flatten_dir = os.path.join(base_dir, "flattened")
    os.makedirs(flatten_dir, exist_ok=True)

def safe_filename(title: str) -> str:
    return re.sub(r'[\\/*?:"<>|]', "", title).strip().replace(" ", "_")

def open_epub(resource_filename) -> epub.EpubBook:
    try:
        if not os.path.exists(resource_filename):
            raise FileNotFoundError(f"Resource file not found : {resource_filename}.")
        print(f"[INFO] Opening book : {resource_filename}.\n")
        book = epub.read_epub(resource_filename)
        if book is None:
            raise epub.EpubException(500, f"Couldn't load the book with the filename : {resource_filename}")
        return book
    except Exception:
        raise

def user_resources_selection() -> str:
    try:
        if not os.path.exists(RESOURCES_PATH):
            raise FileNotFoundError(f"Resources directory not found at {RESOURCES_PATH}")
        resources = [k for k in os.listdir(RESOURCES_PATH) if 'flattened' not in k]
        if not resources:
            raise ValueError("No resources available. Please add EPUB files to the resources directory.")
        print("\n[INFO] ----- AVAILABLE RESOURCES -----")
        for idx, resource in enumerate(resources):
            print(f"[RESULT] [{idx}] -> {resource}")
        try:
            resource_index = int(input("\n[INPUT] Select the resource number that you would like to parse\n >>> "))
        except ValueError:
            raise ValueError("Invalid input. Please enter a number.")
        if resource_index < 0 or resource_index >= len(resources):
            raise IndexError(f"Invalid selection. Please choose a number between 0 and {len(resources)-1}")
        resource_filename = resources[resource_index]
        print(f"[INFO] You selected: {resource_filename}")
        return resource_filename
    except (FileNotFoundError, ValueError, IndexError):
        raise
def user_parsers_selection() -> str:
    try:
        if not os.path.exists(PARSERS_PATH):
            raise FileNotFoundError(f"Parsers directory not found at {PARSERS_PATH}")
        parsers = [k for k in os.listdir(PARSERS_PATH) if "__pycache__" not in k]
        if not parsers:
            raise ValueError("No parsers available. Please add EPUB files to the resources directory.")
        print("\n[INFO] ----- AVAILABLE PARSERS -----")
        for idx, parser in enumerate(parsers):
            print(f"[{idx}] -> {parser}")
        try:
            parser_index = int(input("\n[INPUT] Select the parser number with with you would like to parser the selected resource.\n >>> "))
        except ValueError:
            raise ValueError("Invalid input. Please enter a number.")
        if parser_index < 0 or parser_index >= len(parsers):
            raise IndexError(f"Invalid selection. Please choose a number between 0 and {len(parsers)-1}")
        parser_filename = parsers[parser_index]
        print(f"[INFO] You selected: {parser_filename}")
        return parser_filename
    except (FileNotFoundError, ValueError, IndexError):
        raise

def load_selected_parser(parser_filename, book):
    parser_module = load_parser(parser_filename)
    # Check if the parser has a parse_epub method
    if hasattr(parser_module, 'parse_epub'):
        result = parser_module.parse_epub(book)
        print(f"Parsing completed with {parser_filename}")
        print(f"Parsed {len(result) if result else 0} items")
        return result
    else:
        print(f"Error: {parser_filename} doesn't have a parse_epub method")

def main():
    try:
        create_flattened_dir()
        resource_filename = user_resources_selection()
        parser_filename = user_parsers_selection()
        book = open_epub(os.path.join(RESOURCES_PATH, resource_filename))
        # we're not doing anything with the result for now
        result = load_selected_parser(parser_filename, book)
    except Exception as e:
        print(f"Unexpected error: {e}")

if __name__ == "__main__":
    main()

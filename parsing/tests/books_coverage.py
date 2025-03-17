
from dotenv import load_dotenv
from enums.books import BOOKS_EN, BOOKS_FR, SEED_BOOKS, FRENCH_TO_ENGLISH

load_dotenv()


def test_book_coverage(test_verses):
    """
    Tests book coverage, validating start/end, missing books, and chapter count.
    Each element of test_verses is a tuple (book, chapter, verse, text).
    """

    lang = input("Choisir la langue [fr/en] (défaut=en): ").strip().lower() or 'en'
    book_list = BOOKS_EN if lang == 'en' else BOOKS_FR
    start_books = ["Genesis"] if lang == 'en' else ["Genèse"]
    end_books = ["Revelation"] if lang == 'en' else ["Apocalypse"]

    # Récupère la liste (ensemble) des livres détectés depuis test_verses
    # (par ex. "Ephésiens", "Éphésiens", "Ephesiens", etc.)
    detected_books = {v[0] for v in test_verses}

    # Premier et dernier livre détectés
    first_book = test_verses[0][0] if test_verses else None
    last_book = test_verses[-1][0] if test_verses else None

    # Vérification du début et de la fin
    valid_start = first_book in start_books
    valid_end = last_book in end_books

    missing_canonical = set()
    missing_deuterocanonical = set()

    # Parcourt chaque livre "officiel" dans book_list
    for book in book_list:
        # Convertit le livre "officiel" en anglais s'il est français
        if lang == 'fr':
            # S'il existe un mapping, on l'utilise, sinon on garde le nom d'origine
            title_en = FRENCH_TO_ENGLISH.get(book, book)
        else:
            # Déjà en anglais
            title_en = book

        # On vérifie si ce livre (en anglais) est détecté
        is_detected = False

        # Parcourt tous les livres réellement détectés par le parsing
        for detected_book in detected_books:
            # Convertit le livre détecté en anglais si on est en français
            detected_book_en = FRENCH_TO_ENGLISH.get(detected_book, detected_book)

            # Compare les versions anglaises
            if detected_book_en == title_en:
                is_detected = True
                break

        # Vérifie si ce livre est deutérocanonique ou non (basé sur SEED_BOOKS)
        is_deuterocanonical = False
        for seed_book in SEED_BOOKS:
            if seed_book["title"] == title_en and seed_book["is_deuterocanonical"] == 1:
                is_deuterocanonical = True
                break

        # S'il n'a pas été détecté, on l'ajoute aux livres manquants
        if not is_detected:
            if is_deuterocanonical:
                missing_deuterocanonical.add(book)
            else:
                missing_canonical.add(book)

    # Calcul du nombre de chapitres par livre détecté
    chapters_per_book = {}
    for book, chapter, verse, text in test_verses:
        if book not in chapters_per_book:
            chapters_per_book[book] = set()
        chapters_per_book[book].add(chapter)

    # ---- Affichage des résultats ----
    print(f"\n=== TEST COMPLET ({lang.upper()}) ===")

    # Affiche la liste des livres détectés
    print(f"Livres détectés ({len(detected_books)}/{len(book_list)}) :")
    print(", ".join(sorted(detected_books)))

    # Affiche les livres canoniques manquants
    print(f"\nLivres canoniques manquants ({len(missing_canonical)}) :")
    if missing_canonical:
        print("⚠️ " + ", ".join(sorted(missing_canonical)))
    else:
        print("✅ Aucun livre canonique manquant !")

    # Affiche les livres deutérocanoniques manquants
    print(f"\nLivres deutérocanoniques manquants ({len(missing_deuterocanonical)}) :")
    if missing_deuterocanonical:
        print("⚠️ " + ", ".join(sorted(missing_deuterocanonical)))
    else:
        print("✅ Aucun livre deutérocanonique manquant !")

    # Validation du début et de la fin
    print(f"\nValidation du début : {'✅' if valid_start else '❌'} {first_book}")
    print(f"Validation de la fin : {'✅' if valid_end else '❌'} {last_book}")

    # Affiche le nombre de chapitres par livre détecté
    print("\nNombre de chapitres par livre :")
    for book in book_list:
        # On reste dans la langue d'affichage, ici "book" est l'élément "officiel"
        chapters = chapters_per_book.get(book, set())
        nb_chapters = len(chapters)
        print(f"{book}: {nb_chapters} chapitre{'s' if nb_chapters > 1 else ''}")

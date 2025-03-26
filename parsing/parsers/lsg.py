
def parse(book_path):

def flatten_text_sections(book):
    print("flattening")
    """
    Lit un fichier EPUB et crée un unique fichier texte
    contenant le contenu brut de tous les documents dans l'ordre de lecture,
    SANS balises HTML (uniquement le texte).
    """

    # Create output filename
    original_filename = os.path.basename("resources")
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


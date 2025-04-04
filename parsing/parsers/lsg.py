import os
from bs4 import BeautifulSoup

def flatten_text_sections(book, flatten_dir, metadata):
    print("Flattening...")
    print(metadata)
    title_value = "unknown"
    if "title" in metadata and metadata["title"]:
        title_list = metadata["title"]
        if isinstance(title_list, list) and len(title_list) > 0:
            title_value = title_list[0][0]
    sanitized_title = "".join(c if c.isalnum() or c in " -_." else "_" for c in title_value)
    output_path = os.path.join(flatten_dir, f"{sanitized_title}_flattened.txt")

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

    os.makedirs(flatten_dir, exist_ok=True)
    with open(output_path, "w", encoding="utf-8") as f:
        f.write("\n\n".join(all_text_parts))

    print(f"Created flattened file: {output_path}")
    return output_path

def parse_epub(book, flatten_dir, metadata):
    flatten_text_sections(book, flatten_dir, metadata)
    print("Inside parsing LSG")

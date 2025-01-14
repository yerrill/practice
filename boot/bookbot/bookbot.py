import sys

def word_count(book_text: str) -> int:
    return len(book_text.split())

def character_count(book_text: str) -> list[tuple[str, int]]:
    lower_text = book_text.lower()
    chars = {}
    
    for ch in lower_text:
        if ch in chars:
            chars[ch] += 1
        else:
            chars[ch] = 1
    
    chars = filter(lambda x: x[0].isalpha(), chars.items())
    chars = list(chars)
    chars.sort(reverse=True, key=lambda x: x[1])
    
    return chars

def main(path_to_file):
    with open(path_to_file) as f:
        book_contents = f.read()
    
    print(f"--- {path_to_file} ---")
    print(f"Word Count: {word_count(book_contents)}")
    
    for item in character_count(book_contents):
        print(f"Letter '{item[0]}' was found {item[1]} times")
    
main(sys.argv[1])
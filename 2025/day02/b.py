from typing import List

file = open('input.txt')
# file = open('prova.txt')

lines: List[str] = file.readlines()[0].strip().split(',')
totale = 0
expected = 4174379265

def split_into_equal_parts(s: str, parts: int) -> List[str]:
    """Split `s` into `parts` substrings of equal length.

    Raises ValueError if `s` cannot be evenly divided into `parts` parts.
    """
    if parts < 1:
        raise ValueError("parts must be >= 1")
    if len(s) % parts != 0:
        raise ValueError("String length is not divisible by number of parts")
    size = len(s) // parts
    return [s[i * size:(i + 1) * size] for i in range(parts)]


def all_equal_substrings(s: str, parts: int) -> bool:
    """Return True if splitting `s` into `parts` equal parts yields all identical substrings."""
    try:
        parts_list = split_into_equal_parts(s, parts)
    except ValueError:
        return False
    first = parts_list[0]
    return all(p == first for p in parts_list)


def find_all_equal_partitions(s: str) -> List[int]:
    """Return a list of `parts` values (>1) such that `s` can be split into that many equal identical parts.

    Example: '1212' -> [2] (since '12' == '12') ; '1111' -> [2,4] ('11'=='11', and '1'=='1'=='1'=='1')
    """
    res: List[int] = []
    n = len(s)
    for parts in range(2, n + 1):
        if n % parts != 0:
            continue
        if all_equal_substrings(s, parts):
            res.append(parts)
    return res


if __name__ == '__main__':
    # Example usage: for every index in the input ranges, print any equal partitions found
    for line in lines:
        start, stop = map(int, line.strip().split('-'))
        for index in range(start, stop + 1):
            s = str(index)
            parts = find_all_equal_partitions(s)
            if parts:
                totale += index
    print("Totale:", totale)




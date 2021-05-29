// Package funks implements basic functions that I use frequently to save time on replicating code across different projects.
package funks
 
var AsciiChars = []string{
"!",
"#",
"$",
"",
"&",
"",
"(",
")",
"*",
"+",
",",
"-",
"",
"/",
":",
";",
"<",
"=",
">",
"?",
"@",
"\\",
"]",
"^",
"_",
"`",
"{",
"|",
"}",
"~",
"Ç",
"ü",
"é",
"â",
"ä",
"à",
"å",
"ç",
"ê",
"ë",
"è",
"ï",
"î",
"ì",
"Ä",
"Å",
"É",
"æ",
"Æ",
"ô",
"ö",
"ò",
"û",
"ù",
"ÿ",
"Ö",
"Ü",
"ø",
"£",
"Ø",
"×",
"ƒ",
"á",
"í",
"ó",
"ú",
"ñ",
"Ñ",
"ª",
"º",
"¿",
"®",
"¬",
"½",
"¼",
"¡",
"«",
"»",
"░",
"▒",
"▓",
"│",
"┤",
"Á",
"Â",
"À",
"©",
"╣",
"║",
"╗",
"╝",
"¢",
"¥",
"┐",
"└",
"┴",
"┬",
"├",
"─",
"┼",
"ã",
"Ã",
"╚",
"╔",
"╩",
"╦",
"╠",
"═",
"╬",
"¤",
"ð",
"Ð",
"Ê",
"Ë",
"È",
"ı",
"Í",
"Î",
"Ï",
"┘",
"┌",
"█",
"▄",
"¦",
"Ì",
"▀",
"Ó",
"ß",
"Ô",
"Ò",
"õ",
"Õ",
"µ",
"þ",
"Þ",
"Ú",
"Û",
"Ù",
"ý",
"Ý",
"¯",
"´",
"¬",
"±",
"‗",
"¾",
"¶",
"§",
"÷",
"¸",
"°",
"¨",
"•",
"¹",
"³",
"²",
"■",
" ",
}

var AsciiCharsMap = map[int]string{
    0: "!",
    1: "#",
    2: "$",
    3: "",
    4: "&",
    5: "",
    6: "(",
    7: ")",
    8: "*",
    9: "+",
    10: ",",
    11: "-",
    12: "",
    13: "/",
    14: ":",
    15: ";",
    16: "<",
    17: "=",
    18: ">",
    19: "?",
    20: "@",
    21: "\\",
    22: "]",
    23: "^",
    24: "_",
    25: "`",
    26: "{",
    27: "|",
    28: "}",
    29: "~",
    30: "Ç",
    31: "ü",
    32: "é",
    33: "â",
    34: "ä",
    35: "à",
    36: "å",
    37: "ç",
    38: "ê",
    39: "ë",
    40: "è",
    41: "ï",
    42: "î",
    43: "ì",
    44: "Ä",
    45: "Å",
    46: "É",
    47: "æ",
    48: "Æ",
    49: "ô",
    50: "ö",
    51: "ò",
    52: "û",
    53: "ù",
    54: "ÿ",
    55: "Ö",
    56: "Ü",
    57: "ø",
    58: "£",
    59: "Ø",
    60: "×",
    61: "ƒ",
    62: "á",
    63: "í",
    64: "ó",
    65: "ú",
    66: "ñ",
    67: "Ñ",
    68: "ª",
    69: "º",
    70: "¿",
    71: "®",
    72: "¬",
    73: "½",
    74: "¼",
    75: "¡",
    76: "«",
    77: "»",
    78: "░",
    79: "▒",
    80: "▓",
    81: "│",
    82: "┤",
    83: "Á",
    84: "Â",
    85: "À",
    86: "©",
    87: "╣",
    88: "║",
    89: "╗",
    90: "╝",
    91: "¢",
    92: "¥",
    93: "┐",
    94: "└",
    95: "┴",
    96: "┬",
    97: "├",
    98: "─",
    99: "┼",
    100: "ã",
    101: "Ã",
    102: "╚",
    103: "╔",
    104: "╩",
    105: "╦",
    106: "╠",
    107: "═",
    108: "╬",
    109: "¤",
    110: "ð",
    111: "Ð",
    112: "Ê",
    113: "Ë",
    114: "È",
    115: "ı",
    116: "Í",
    117: "Î",
    118: "Ï",
    119: "┘",
    120: "┌",
    121: "█",
    122: "▄",
    123: "¦",
    124: "Ì",
    125: "▀",
    126: "Ó",
    127: "ß",
    128: "Ô",
    129: "Ò",
    130: "õ",
    131: "Õ",
    132: "µ",
    133: "þ",
    134: "Þ",
    135: "Ú",
    136: "Û",
    137: "Ù",
    138: "ý",
    139: "Ý",
    140: "¯",
    141: "´",
    142: "¬",
    143: "±",
    144: "‗",
    145: "¾",
    146: "¶",
    147: "§",
    148: "÷",
    149: "¸",
    150: "°",
    151: "¨",
    152: "•",
    153: "¹",
    154: "³",
    155: "²",
    156: "■",
    157: " ",
}
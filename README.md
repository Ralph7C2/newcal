# newcal

A Go library and CLI tool for working with ["The New Calendar"](https://thenewcalendar.com) - a reformed calendar system featuring 5 seasons and 9-day weeks.

## The New Calendar System

The New Calendar reimagines how we track time:

- **5 Seasons**: Winter, Spring, Summer, Autumn, and Fall - each exactly 73 days long
- **9-Day Weeks**: Named after the planets (Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto)
- **Mid-Season Days**: Day 37 of each season is a special "Mid Season" day that doesn't belong to any week
- **Leap Days**: Leap years have a special "Leap Day" between Winter and Spring (occurring on day 70 of the year)
- **365/366 Days**: Standard years have 365 days (5 × 73), leap years add one extra day

### Calendar Structure

Each season has 73 days divided into:
- **Early Half**: Days 1-36 (4 complete weeks)
- **Mid-Season Day**: Day 37 (standalone day)
- **Late Half**: Days 38-73 (4 complete weeks)

## Installation

### As a CLI tool

```bash
go install github.com/ralph7c2/newcal/cmd/cli@latest
```

### As a library

```bash
go get github.com/ralph7c2/newcal
```

## Usage

### CLI Tool

#### Display today's date in the New Calendar

```bash
newcal
```

Output:
```
Today is Saturn, 15 Early Winter, 2025
*-------------------------------------------------------------------------------*
|                              Early Winter                                   |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
|  Mer  ||  Ven  ||  Ear  ||  Mar  ||  Jup  ||  Sat  ||  Ura  ||  Nep  ||  Plu  |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
|     1 ||     2 ||     3 ||     4 ||     5 ||     6 ||     7 ||     8 ||     9 |
|    10 ||    11 ||    12 ||    13 ||    14 ||  *15* ||    16 ||    17 ||    18 |
|    19 ||    20 ||    21 ||    22 ||    23 ||    24 ||    25 ||    26 ||    27 |
|    28 ||    29 ||    30 ||    31 ||    32 ||    33 ||    34 ||    35 ||    36 |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
```

#### Convert a Gregorian date

```bash
newcal --gregorian 2024-07-04
# or
newcal -g 2024-07-04
```

Output:
```
2024-07-04 is Mars, 50 Late Summer, 2024
*-------------------------------------------------------------------------------*
|                                  Late Summer                                  |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
|  Mer  ||  Ven  ||  Ear  ||  Mar  ||  Jup  ||  Sat  ||  Ura  ||  Nep  ||  Plu  |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
|    38 ||    39 ||    40 ||    41 ||    42 ||    43 ||    44 ||    45 ||    46 |
|    47 ||    48 ||    49 ||  *50* ||    51 ||    52 ||    53 ||    54 ||    55 |
|    56 ||    57 ||    58 ||    59 ||    60 ||    61 ||    62 ||    63 ||    64 |
|    65 ||    66 ||    67 ||    68 ||    69 ||    70 ||    71 ||    72 ||    73 |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
```

#### Convert a Unix timestamp

```bash
newcal --unix 1234567890
# or
newcal -u 1234567890
```

Output:
```
1234567890 is Pluto, 55 Late Winter, 2009
*-------------------------------------------------------------------------------*
|                                  Late Winter                                  |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
|  Mer  ||  Ven  ||  Ear  ||  Mar  ||  Jup  ||  Sat  ||  Ura  ||  Nep  ||  Plu  |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
|    38 ||    39 ||    40 ||    41 ||    42 ||    43 ||    44 ||    45 ||    46 |
|    47 ||    48 ||    49 ||    50 ||    51 ||    52 ||    53 ||    54 ||  *55* |
|    56 ||    57 ||    58 ||    59 ||    60 ||    61 ||    62 ||    63 ||    64 |
|    65 ||    66 ||    67 ||    68 ||    69 ||    70 ||    71 ||    72 ||    73 |
*-------==-------==-------==-------==-------==-------==-------==-------==-------*
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

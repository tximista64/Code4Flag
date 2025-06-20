#!/usr/bin/python3

def make_flag(s: str) -> str:
    a = s[5]  # char at index 5
    _b = s[2]  # char at index 2, as string

    for s_ in range(len(s)):
        # b = _b.substring(_b.length() - s_) + _b.substring(s_)
        # En python, ça devient
        b = _b[-s_:] + _b[s_:]

        if s_ >= 3:
            _b2 = _b + s[s_ - 3]
        else:
            _b2 = _b + s[len(s) - (3 - s_)]

        if s_ >= len(_b2):
            _b = _b2 + s[s_ - len(_b2)]
        elif len(s) >= len(_b2) - s_:
            _b = _b2 + s[len(s) - (len(_b2) - s_)]
        else:
            _b = _b2 + s[len(s) - ((len(_b2) - s_) - len(s))]

        idx = (((len(s) + len(_b)) * s_) + len(_b)) % len(b)
        a = a + b[idx]

    return a[0:2] + s[3] + a[3] + '0' + a[5:7]


if __name__ == "__main__":
    seed = "somestring??"
    flag = make_flag(seed)
    print(f"Flag pour seed '{seed}': {flag}")


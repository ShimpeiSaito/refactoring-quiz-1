TAX_RATE: float = 1.1


def calc_amount_price(price: int, count: int):
    return price * count * TAX_RATE

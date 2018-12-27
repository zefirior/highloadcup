from collections import defaultdict
from enum import Enum


dict_container = defaultdict(set)


class ParamsBase(Enum):
    def __init__(self, code, field, rank):
        self.code = code
        self.field = field
        self.rank = rank
        container = self.container()
        container.add(code)

    @classmethod
    def container(cls) -> set:
        return dict_container[cls]

    @classmethod
    def exists_param(cls, code):
        container = cls.container()
        return code in container


class AccField(Enum):
    SEX = "sex"
    EMAIL = "email"
    STATUS = "status"
    FNAME = "fname"
    SNAME = "sname"
    PHONE = "phone"
    COUNTRY = "country"
    INTERESTS = "interests"
    LIKES = "likes"
    PREMIUM = "premium"


class ApiParam(ParamsBase):
    SEX_EQ =             ("sex_eq",             AccField.SEX,       1)
    EMAIL_DOMAIN =       ("email_domain",       AccField.EMAIL,     1)
    EMAIL_LT =           ("email_lt",           AccField.EMAIL,     1)
    EMAIL_GT =           ("email_gt",           AccField.EMAIL,     1)
    STATUS_EQ =          ("status_eq",          AccField.STATUS,    1)
    STATUS_NEQ =         ("status_neq",         AccField.STATUS,    1)
    FNAME_EQ =           ("fname_eq",           AccField.FNAME,     1)
    FNAME_ANY =          ("fname_any",          AccField.FNAME,     1)
    FNAME_NULL =         ("fname_null",         AccField.FNAME,     1)
    SNAME_EQ =           ("sname_eq",           AccField.SNAME,     10)
    SNAME_STARTS =       ("sname_starts",       AccField.SNAME,     1)
    SNAME_NULL =         ("sname_null",         AccField.SNAME,     1)
    PHONE_CODE =         ("phone_code",         AccField.PHONE,     1)
    PHONE_NULL =         ("phone_null",         AccField.PHONE,     1)
    COUNTRY_EQ =         ("country_eq",         AccField.COUNTRY,   9)
    COUNTRY_NULL =       ("country_null",       AccField.COUNTRY,   1)
    INTERESTS_CONTAINS = ("interests_contains", AccField.INTERESTS, 1)
    INTERESTS_ANY =      ("interests_any",      AccField.INTERESTS, 1)
    LIKES_CONTAINS =     ("likes_contains",     AccField.LIKES,     1)
    PREMIUM_NOW =        ("premium_now",        AccField.PREMIUM,   1)
    PREMIUM_NULL =       ("premium_null",       AccField.PREMIUM,   1)


def check_filter_param(param_codes: list):
    for code in param_codes:
        if not ApiParam.exists_param(code):
            return False
    return True
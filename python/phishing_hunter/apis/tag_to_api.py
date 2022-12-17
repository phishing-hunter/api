import typing_extensions

from phishing_hunter.apis.tags import TagValues
from phishing_hunter.apis.tags.hunting_api import HuntingApi
from phishing_hunter.apis.tags.observation_api import ObservationApi
from phishing_hunter.apis.tags.scanner_api import ScannerApi
from phishing_hunter.apis.tags.settings_api import SettingsApi
from phishing_hunter.apis.tags.user_api import UserApi

TagToApi = typing_extensions.TypedDict(
    'TagToApi',
    {
        TagValues.HUNTING: HuntingApi,
        TagValues.OBSERVATION: ObservationApi,
        TagValues.SCANNER: ScannerApi,
        TagValues.SETTINGS: SettingsApi,
        TagValues.USER: UserApi,
    }
)

tag_to_api = TagToApi(
    {
        TagValues.HUNTING: HuntingApi,
        TagValues.OBSERVATION: ObservationApi,
        TagValues.SCANNER: ScannerApi,
        TagValues.SETTINGS: SettingsApi,
        TagValues.USER: UserApi,
    }
)

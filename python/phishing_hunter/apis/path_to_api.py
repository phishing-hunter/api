import typing_extensions

from phishing_hunter.paths import PathValues
from phishing_hunter.apis.paths.scanner import Scanner
from phishing_hunter.apis.paths.scanner_add import ScannerAdd
from phishing_hunter.apis.paths.scanner_remove import ScannerRemove
from phishing_hunter.apis.paths.hunting import Hunting
from phishing_hunter.apis.paths.observation import Observation
from phishing_hunter.apis.paths.observation_add import ObservationAdd
from phishing_hunter.apis.paths.observation_remove import ObservationRemove
from phishing_hunter.apis.paths.notify import Notify
from phishing_hunter.apis.paths.apikeys import Apikeys
from phishing_hunter.apis.paths.users_info import UsersInfo
from phishing_hunter.apis.paths.apikey import Apikey
from phishing_hunter.apis.paths.apikey_create import ApikeyCreate
from phishing_hunter.apis.paths.apikey_delete import ApikeyDelete

PathToApi = typing_extensions.TypedDict(
    'PathToApi',
    {
        PathValues.SCANNER: Scanner,
        PathValues.SCANNER_ADD: ScannerAdd,
        PathValues.SCANNER_REMOVE: ScannerRemove,
        PathValues.HUNTING: Hunting,
        PathValues.OBSERVATION: Observation,
        PathValues.OBSERVATION_ADD: ObservationAdd,
        PathValues.OBSERVATION_REMOVE: ObservationRemove,
        PathValues.NOTIFY: Notify,
        PathValues.APIKEYS: Apikeys,
        PathValues.USERS_INFO: UsersInfo,
        PathValues.APIKEY: Apikey,
        PathValues.APIKEY_CREATE: ApikeyCreate,
        PathValues.APIKEY_DELETE: ApikeyDelete,
    }
)

path_to_api = PathToApi(
    {
        PathValues.SCANNER: Scanner,
        PathValues.SCANNER_ADD: ScannerAdd,
        PathValues.SCANNER_REMOVE: ScannerRemove,
        PathValues.HUNTING: Hunting,
        PathValues.OBSERVATION: Observation,
        PathValues.OBSERVATION_ADD: ObservationAdd,
        PathValues.OBSERVATION_REMOVE: ObservationRemove,
        PathValues.NOTIFY: Notify,
        PathValues.APIKEYS: Apikeys,
        PathValues.USERS_INFO: UsersInfo,
        PathValues.APIKEY: Apikey,
        PathValues.APIKEY_CREATE: ApikeyCreate,
        PathValues.APIKEY_DELETE: ApikeyDelete,
    }
)

package ptttype

import "github.com/Ptt-official-app/go-pttbbs/types"

func config() {
	serviceModeStr := setStringConfig("SERVICE_MODE", "DEV")
	SERVICE_MODE = stringToServiceMode(serviceModeStr)

	//////////
	// make.conf
	//////////
	BBSHOME = setStringConfig("BBSHOME", BBSHOME)

	//////////
	// pttbbs.conf
	//////////
	BBSNAME = setStringConfig("BBSNAME", BBSNAME)
	BBSENAME = setStringConfig("BBSENAME", BBSENAME)
	MYHOSTNAME = setStringConfig("MYHOSTNAME", MYHOSTNAME)
	MYIP = setStringConfig("MYIP", MYIP)

	QUERY_ARTICLE_URL = setBoolConfig("QUERY_ARTICLE_URL", QUERY_ARTICLE_URL)
	URL_PREFIX = setStringConfig("URL_PREFIX", URL_PREFIX)

	BN_ARTDSN = setStringConfig("BN_ARTDSN", BN_ARTDSN)

	BN_BBSMOVIE = setStringConfig("BN_BBSMOVIE", BN_BBSMOVIE)

	BN_WHOAMI = setStringConfig("BN_WHOAMI", BN_WHOAMI)

	IS_BN_FIVECHESS_LOG_INFERRED = setBoolConfig("IS_BN_FIVECHESS_LOG_INFERRED", IS_BN_FIVECHESS_LOG_INFERRED)
	BN_FIVECHESS_LOG = setStringConfig("BN_FIVECHESS_LOG", BN_FIVECHESS_LOG)
	IS_BN_CCHESS_LOG_INFERRED = setBoolConfig("IS_BN_CCHESS_LOG_INFERRED", IS_BN_CCHESS_LOG_INFERRED)
	BN_CCHESS_LOG = setStringConfig("BN_CCHESS_LOG", BN_CCHESS_LOG)

	BN_NOTE_AGGCHKDIR = setStringConfig("BN_NOTE_AGGCHKDIR", BN_NOTE_AGGCHKDIR)

	EDITPOST_SMARTMERGE = setBoolConfig("EDITPOST_SMARTMERGE", EDITPOST_SMARTMERGE)

	MULTI_WELCOME_LOGIN = setBoolConfig("MULTI_WELCOME_LOGIN", MULTI_WELCOME_LOGIN)

	ALL_REEDIT_LOG = setBoolConfig("ALL_REEDIT_LOG", ALL_REEDIT_LOG)

	BMCHS = setBoolConfig("BMCHS", BMCHS)

	OUTJOBSPOOL = setBoolConfig("OUTJOBSPOOL", OUTJOBSPOOL)

	NO_GAMBLE = setBoolConfig("NO_GAMBLE", NO_GAMBLE)

	DYMAX_ACTIVE = setBoolConfig("DYMAX_ACTIVE", DYMAX_ACTIVE)

	CPULIMIT_PER_DAY = setIntConfig("CPULIMIT_PER_DAY", CPULIMIT_PER_DAY)

	DEBUGSLEEP = setBoolConfig("DEBUGSLEEP", DEBUGSLEEP)

	DEBUG_FWDADDRERR = setBoolConfig("DEBUG_FWDADDRERR", DEBUG_FWDADDRERR)

	CRITICAL_MEMORY = setBoolConfig("CRITICAL_MEMORY", CRITICAL_MEMORY)

	PRE_FORK = setIntConfig("PRE_FORK", PRE_FORK)

	CONVERT = setBoolConfig("CONVERT", CONVERT)

	COLORDATE = setBoolConfig("COLORDATE", COLORDATE)

	HAVE_USERAGREEMENT = setStringConfig("HAVE_USERAGREEMENT", HAVE_USERAGREEMENT)
	HAVE_USERAGREEMENT_VERSION = setStringConfig("HAVE_USERAGREEMENT_VERSION", HAVE_USERAGREEMENT_VERSION)
	HAVE_USERAGREEMENT_ACCEPTABLE = setStringConfig("HAVE_USERAGREEMENT_ACCEPTABLE", HAVE_USERAGREEMENT_ACCEPTABLE)

	DBCSAWARE = setBoolConfig("DBCSAWARE", DBCSAWARE)

	GUEST_DEFAULT_DBCS_NOINTRESC = setBoolConfig("GUEST_DEFAULT_DBCS_NOINTRESC", GUEST_DEFAULT_DBCS_NOINTRESC)

	USE_PMORE = setBoolConfig("USE_PMORE", USE_PMORE)
	USE_RFORK = setBoolConfig("USE_RFORK", USE_RFORK)
	USE_HUGETLB = setBoolConfig("USE_HUGETLB", USE_HUGETLB)

	SHMALIGNEDSIZE = setIntConfig("SHMALIGNEDSIZE", SHMALIGNEDSIZE)

	USE_COOLDOWN = setBoolConfig("USE_COOLDOWN", USE_COOLDOWN)

	SAFE_ARTICLE_DELETE = setBoolConfig("SAFE_ARTICLE_DELETE", SAFE_ARTICLE_DELETE)

	NOKILLWATERBALL = setBoolConfig("NOKILLWATERBALL", NOKILLWATERBALL)

	NO_SYSOP_ACCOUNT = setBoolConfig("NO_SYSOP_ACCOUNT", NO_SYSOP_ACCOUNT)

	PLAY_ANGEL = setBoolConfig("PLAY_ANGEL", PLAY_ANGEL)

	OLDRECOMMEND = setBoolConfig("OLDRECOMMEND", OLDRECOMMEND)

	GUESTRECOMMEND = setBoolConfig("GUESTRECOMMEND", GUESTRECOMMEND)

	FASTRECMD_LIMIT = setIntConfig("FASTRECMD_LIMIT", FASTRECMD_LIMIT)

	USE_AUTOCPLOG = setBoolConfig("USE_AUTOCPLOG", USE_AUTOCPLOG)

	DEFAULT_AUTOCPLOG = setBoolConfig("DEFAULT_AUTOCPLOG", DEFAULT_AUTOCPLOG)

	TIMET64 = setBoolConfig("TIMET64", TIMET64)

	UTMPD = setBoolConfig("UTMPD", UTMPD)
	UTMPD_ADDR = setStringConfig("UTMPD_ADDR", UTMPD_ADDR)

	NOFLOODING = setBoolConfig("NOFLOODING", NOFLOODING)

	FROMD = setBoolConfig("FROMD", FROMD)

	NO_GUEST_ACCOUNT_REG = setBoolConfig("NO_GUEST_ACCOUNT_REG", NO_GUEST_ACCOUNT_REG)

	EMAILDB_LIMIT = setIntConfig("EMAILDB_LIMIT", EMAILDB_LIMIT)

	USE_REG_CAPTCHA = setBoolConfig("USE_REG_CAPTCHA", USE_REG_CAPTCHA)
	USE_POST_CAPTCHA_FOR_NOREG = setBoolConfig("USE_POST_CAPTCHA_FOR_NOREG", USE_POST_CAPTCHA_FOR_NOREG)
	USE_REMOTE_CAPTCHA = setBoolConfig("USE_REMOTE_CAPTCHA", USE_REMOTE_CAPTCHA)
	CAPTCHA_INSERT_SERVER_ADDR = setStringConfig("CAPTCHA_INSERT_SERVER_ADDR", CAPTCHA_INSERT_SERVER_ADDR)
	IS_CAPTCHA_INSERT_HOST_INFERRED = setBoolConfig("IS_CAPTCHA_INSERT_HOST_INFERRED", IS_CAPTCHA_INSERT_HOST_INFERRED)
	CAPTCHA_INSERT_HOST = setStringConfig("CAPTCHA_INSERT_HOST", CAPTCHA_INSERT_HOST)
	CAPTCHA_INSERT_URI = setStringConfig("CAPTCHA_INSERT_URI", CAPTCHA_INSERT_URI)
	CAPTCHA_INSERT_SECRET = setStringConfig("CAPTCHA_INSERT_SECRET", CAPTCHA_INSERT_SECRET)
	CAPTCHA_URL_PREFIX = setStringConfig("CAPTCHA_URL_PREFIX", CAPTCHA_URL_PREFIX)
	CAPTCHA_CODE_LENGTH = setIntConfig("CAPTCHA_CODE_LENGTH", CAPTCHA_CODE_LENGTH)

	REQUIRE_SECURE_CONN_TO_REGISTER = setBoolConfig("REQUIRE_SECURE_CONN_TO_REGISTER", REQUIRE_SECURE_CONN_TO_REGISTER)
	REQUIRE_VERIFY_EMAIL_AT_REGISTER = setBoolConfig("REQUIRE_VERIFY_EMAIL_AT_REGISTER", REQUIRE_VERIFY_EMAIL_AT_REGISTER)

	INSCREEN = setStringConfig("INSCREEN", INSCREEN)

	//config.h
	BBSPROGPOSTFIX = setStringConfig("BBSPROGPOSTFIX", BBSPROGPOSTFIX)
	BAN_FILE = setStringConfig("BAN_FILE", BAN_FILE)
	LOAD_FILE = setStringConfig("LOAD_FILE", LOAD_FILE)

	BBSMNAME = setStringConfig("BBSMNAME", BBSMNAME)
	BBSMNAME2 = setStringConfig("BBSMNAME2", BBSMNAME2)

	IS_MONEYNAME_INFFERRED = setBoolConfig("IS_MONEYNAME_INFFERRED", IS_MONEYNAME_INFFERRED)
	MONEYNAME = setStringConfig("MONEYNAME", MONEYNAME)

	IS_AID_HOSTNAME_INFERRED = setBoolConfig("IS_AID_HOSTNAME_INFERRED", IS_AID_HOSTNAME_INFERRED)
	AID_HOSTNAME = setStringConfig("AID_HOSTNAME", AID_HOSTNAME)

	TITLE_COLOR = setColorConfig("TITLE_COLOR", TITLE_COLOR)

	HLP_CATEGORY_COLOR = setColorConfig("HLP_CATEGORY_COLOR", HLP_CATEGORY_COLOR)

	HLP_DESCRIPTION_COLOR = setColorConfig("HLP_DESCRIPTION_COLOR", HLP_DESCRIPTION_COLOR)

	HLP_KEYLIST_COLOR = setColorConfig("HLP_KEYLIST_COLOR", HLP_KEYLIST_COLOR)

	BBSUSER = setStringConfig("BBSUSER", BBSUSER)
	BBSUID = setIntConfig("BBSUID", BBSUID)
	BBSGID = setIntConfig("BBSGID", BBSGID)

	TAR_PATH = setStringConfig("TAR_PATH", TAR_PATH)
	MUTT_PATH = setStringConfig("MUTT_PATH", MUTT_PATH)

	DEFAULT_FOLDER_CREATE_PERM = setIntConfig("DEFAULT_FOLDER_CREATE_PERM", DEFAULT_FOLDER_CREATE_PERM)

	DEFAULT_FILE_CREATE_PERM = setIntConfig("DEFAULT_FILE_CREATE_PERM", DEFAULT_FILE_CREATE_PERM)

	SHM_KEY = types.Key_t(setIntConfig("SHM_KEY", int(SHM_KEY)))

	PASSWDSEM_KEY = setIntConfig("PASSWDSEM_KEY", PASSWDSEM_KEY)

	MEM_CHECK = setIntConfig("MEM_CHECK", MEM_CHECK)

	RELAY_SERVER_IP = setStringConfig("RELAY_SERVER_IP", RELAY_SERVER_IP)

	XCHATD_ADDR = setStringConfig("XCHATD_ADDR", XCHATD_ADDR)

	/////////////////////////////////////////////////////////////////////////////
	// Default Board Names 預設看板名稱

	BN_SECURITY = setStringConfig("BN_SECURITY", BN_SECURITY)
	BN_NOTE = setStringConfig("BN_NOTE", BN_NOTE)
	BN_RECORD = setStringConfig("BN_RECORD", BN_RECORD)

	BN_SYSOP = setStringConfig("BN_SYSOP", BN_SYSOP)
	BN_TEST = setStringConfig("BN_TEST", BN_TEST)
	BN_BUGREPORT = setStringConfig("BN_BUGREPORT", BN_BUGREPORT)
	BN_LAW = setStringConfig("BN_LAW", BN_LAW)
	BN_NEWBIE = setStringConfig("BN_NEWBIE", BN_NEWBIE)
	BN_ASKBOARD = setStringConfig("BN_ASKBOARD", BN_ASKBOARD)
	BN_FOREIGN = setStringConfig("BN_FOREIGN", BN_FOREIGN)
	BN_ID_PROBLEM = setStringConfig("BN_ID_PROBLEM", BN_ID_PROBLEM)
	BN_DELETED = setStringConfig("BN_DELETED", BN_DELETED)
	BN_JUNK = setStringConfig("BN_JUNK", BN_JUNK)

	BN_POLICELOG = setStringConfig("BN_POLICELOG", BN_POLICELOG)
	BN_UNANONYMOUS = setStringConfig("BN_UNANONYMOUS", BN_UNANONYMOUS)
	BN_NEWIDPOST = setStringConfig("BN_NEWIDPOST", BN_NEWIDPOST)
	BN_ALLPOST = setStringConfig("BN_ALLPOST", BN_ALLPOST)
	BN_ALLHIDPOST = setStringConfig("BN_ALLHIDPOST", BN_ALLHIDPOST)

	MAX_GUEST = setIntConfig("MAX_GUEST", MAX_GUEST)
	MAX_CPULOAD = setIntConfig("MAX_CPULOAD", MAX_CPULOAD)
	DEBUGSLEEP_SECONDS = setIntConfig("DEBUGSLEEP_SECONDS", DEBUGSLEEP_SECONDS)

	OVERLOADBLOCKFDS = setIntConfig("OVERLOADBLOCKFDS", OVERLOADBLOCKFDS)
	MAX_SWAPUSED = setDoubleConfig("MAX_SWAPUSED", MAX_SWAPUSED)

	TARQUEUE_TIME_STR = setStringConfig("TARQUEUE_TIME_STR", TARQUEUE_TIME_STR)

	/////////////////////////////////////////////////////////////////////////////
	// More system messages 系統訊息

	RECYCLE_BIN_NAME = setStringConfig("RECYCLE_BIN_NAME", RECYCLE_BIN_NAME)

	TIME_CAPSULE_NAME = setStringConfig("TIME_CAPSULE_NAME", TIME_CAPSULE_NAME)

	MAX_POST_MONEY = setIntConfig("MAX_POST_MONEY", MAX_POST_MONEY)

	MAX_GUEST_LIFE = setIntConfig("MAX_GUEST_LIFE", MAX_GUEST_LIFE)

	MAX_EDIT_LINE = setIntConfig("MAX_EDIT_LINE", MAX_EDIT_LINE)

	MAX_EDIT_LINE_LARGE = setIntConfig("MAX_EDIT_LINE_LARGE", MAX_EDIT_LINE_LARGE)

	MAX_LIFE = setIntConfig("MAX_LIFE", MAX_LIFE)

	KEEP_DAYS_REGGED = setIntConfig("KEEP_DAYS_REGGED", KEEP_DAYS_REGGED)

	KEEP_DAYS_UNREGGED = setIntConfig("KEEP_DAYS_UNREGGED", KEEP_DAYS_UNREGGED)

	THREAD_SEARCH_RANGE = setIntConfig("THREAD_SEARCH_RANGE", THREAD_SEARCH_RANGE)

	FOREIGN_REG = setBoolConfig("FOREIGN_REG", FOREIGN_REG)

	FOREIGN_REG_DAY = setIntConfig("FOREIGN_REG_DAY", FOREIGN_REG_DAY)

	FORCE_PROCESS_REGISTER_FORM = setIntConfig("FORCE_PROCESS_REGISTER_FORM", FORCE_PROCESS_REGISTER_FORM)

	HBFLexpire = setIntConfig("HBFLexpire", HBFLexpire)

	MAX_EXKEEPMAIL = setIntConfig("MAX_EXKEEPMAIL", MAX_EXKEEPMAIL)

	INNTIMEZONE = setStringConfig("INNTIMEZONE", INNTIMEZONE)

	ADD_EXMAILBOX = setIntConfig("ADD_EXMAILBOX", ADD_EXMAILBOX)

	BADPOST_CLEAR_DURATION = setIntConfig("BADPOST_CLEAR_DURATION", BADPOST_CLEAR_DURATION)

	BADPOST_MIN_CLEAR_DURATION = setIntConfig("BADPOST_MIN_CLEAR_DURATION", BADPOST_MIN_CLEAR_DURATION)

	MAX_CROSSNUM = setIntConfig("MAX_CROSSNUM", MAX_CROSSNUM)

	MAX_QUERYLINES = setIntConfig("MAX_QUERYLINES", MAX_QUERYLINES)

	MAX_LOGIN_INFO = setIntConfig("MAX_LOGIN_INFO", MAX_LOGIN_INFO)

	MAX_POST_INFO = setIntConfig("MAX_POST_INFO", MAX_POST_INFO)

	MAX_NAMELIST = setIntConfig("MAX_NAMELIST", MAX_NAMELIST)

	MAX_NOTE = setIntConfig("MAX_NOTE", MAX_NOTE)

	MAX_SIGLINES = setIntConfig("MAX_SIGLINES", MAX_SIGLINES)

	LOGINATTEMPTS = setIntConfig("LOGINATTEMPTS", LOGINATTEMPTS)

	MAX_KEEPMAIL = setIntConfig("MAX_KEEPMAIL", MAX_KEEPMAIL)

	MAX_KEEPMAIL_SOFTLIMIT = setIntConfig("MAX_KEEPMAIL_SOFTLIMIT", MAX_KEEPMAIL_SOFTLIMIT)

	MAX_KEEPMAIL_HARDLIMIT = setIntConfig("MAX_KEEPMAIL_HARDLIMIT", MAX_KEEPMAIL_HARDLIMIT)

	BADCIDCHARS = setStringConfig("BADCIDCHARS", BADCIDCHARS)

	MAX_ROOM = setIntConfig("MAX_ROOM", MAX_ROOM)

	MAXTAGS = setIntConfig("MAXTAGS", MAXTAGS)

	WRAPMARGIN = setIntConfig("WRAPMARGIN", WRAPMARGIN)

	/////////////////////////////////////////////////////////////////////////////
	// Logging 記錄設定
	LOG_CONF_KEYWORD = setBoolConfig("LOG_CONF_KEYWORD", LOG_CONF_KEYWORD)
	LOG_CONF_INTERNETMAIL = setBoolConfig("LOG_CONF_INTERNETMAIL", LOG_CONF_INTERNETMAIL)
	LOG_CONF_PUSH = setBoolConfig("LOG_CONF_PUSH", LOG_CONF_PUSH)
	LOG_CONF_EDIT_CALENDAR = setBoolConfig("LOG_CONF_EDIT_CALENDAR", LOG_CONF_EDIT_CALENDAR)
	LOG_CONF_POST = setBoolConfig("LOG_CONF_POST", LOG_CONF_POST)
	LOG_CONF_CRAWLER = setBoolConfig("LOG_CONF_CRAWLER", LOG_CONF_CRAWLER)
	LOG_CONF_CROSSPOST = setBoolConfig("LOG_CONF_CROSSPOST", LOG_CONF_CROSSPOST)
	LOG_CONF_BAD_REG_CODE = setBoolConfig("LOG_CONF_BAD_REG_CODE", LOG_CONF_BAD_REG_CODE)
	LOG_CONF_VALIDATE_REG = setBoolConfig("LOG_CONF_VALIDATE_REG", LOG_CONF_VALIDATE_REG)
	LOG_CONF_MASS_DELETE = setBoolConfig("LOG_CONF_MASS_DELETE", LOG_CONF_MASS_DELETE)
	LOG_CONF_OSONG_VERBOSE = setBoolConfig("LOG_CONF_OSONG_VERBOSE", LOG_CONF_OSONG_VERBOSE)
	LOG_CONF_EDIT_TITLE = setBoolConfig("LOG_CONF_EDIT_TITLE", LOG_CONF_EDIT_TITLE)

	/////////////////////////////////////////////////////////////////////////////
	// Default Configurations 預設參數
	LOGINASNEW = setBoolConfig("LOGINASNEW", LOGINASNEW)
	DOTIMEOUT = setBoolConfig("DOTIMEOUT", DOTIMEOUT)
	INTERNET_EMAIL = setBoolConfig("INTERNET_EMAIL", INTERNET_EMAIL)
	SHOWUID = setBoolConfig("SHOWUID", SHOWUID)
	HAVE_ANONYMOUS = setBoolConfig("HAVE_ANONYMOUS", HAVE_ANONYMOUS)
	HAVE_ORIGIN = setBoolConfig("HAVE_ORIGIN", HAVE_ORIGIN)
	USE_BSMTP = setBoolConfig("USE_BSMTP", USE_BSMTP)
	REJECT_FLOOD_POST = setBoolConfig("REJECT_FLOOD_POST", REJECT_FLOOD_POST)

	IDLE_TIMEOUT = setIntConfig("IDLE_TIMEOUT", IDLE_TIMEOUT)
	SHOW_IDLE_TIME = setBoolConfig("SHOW_IDLE_TIME", SHOW_IDLE_TIME)

	//////////
	// common.h
	//////////
	SZ_RECENTLOGIN = setIntConfig("SZ_RECENTLOGIN", SZ_RECENTLOGIN)
	SZ_RECENTPAY = setIntConfig("SZ_RECENTPAY", SZ_RECENTPAY)

	USE_EDIT_HISTORY = setBoolConfig("USE_EDIT_HISTORY", USE_EDIT_HISTORY)
	FN_SAFEDEL = setStringConfig("FN_SAFEDEL", FN_SAFEDEL)
	FN_SAFEDEL_PREFIX_LEN = setIntConfig("FN_SAFEDEL_PREFIX_LEN", FN_SAFEDEL_PREFIX_LEN)

	STR_SAFEDEL_TITLE = setStringConfig("STR_SAFEDEL_TITLE", STR_SAFEDEL_TITLE)

	SAFE_ARTICLE_DELETE_NUSER = setIntConfig("SAFE_ARTICLE_DELETE_NUSER", SAFE_ARTICLE_DELETE_NUSER)

	//////////
	// proto.h
	//////////
	USE_COMMENTD = setBoolConfig("USE_COMMENTD", USE_COMMENTD)
	USE_EMAILDB = setBoolConfig("USE_EMAILDB", USE_EMAILDB)
	USE_REGCHECKD = setBoolConfig("USE_REGCHECKD", USE_REGCHECKD)
	USE_VERIFYDB = setBoolConfig("USE_VERIFYDB", USE_VERIFYDB)

	//additional config
	IS_NEW_SHM = setBoolConfig("IS_NEW_SHM", IS_NEW_SHM)

}

package errorcode

var errorCode = map[string]APIError{
	/** 共同 **/
	"SUCCESS":                 {1, "SUCCESS"},                    // 呼叫API成功
	"PERMISSION_DENIED":       {403, "PERMISSION DENIED"},        // 權限不足
	"CREATE_DIR_ERROR":        {1001, "CREATE DIR ERROR"},        // 建立資料夾失敗
	"GET_UPLOAD_FILE_ERROR":   {1002, "GET UPLOAD FILE ERROR"},   // 取得上傳檔案失敗
	"CREATE_FILE_ERROR":       {1003, "CREATE FILE ERROR"},       // 建立檔案失敗
	"GET_UPLOAD_TYPE_ERROR":   {1004, "GET UPLOAD TYPE ERROR"},   // 取得上傳類型錯誤
	"JSON_MARSHAL_ERROR":      {1005, "JSON MARSHAL ERROR"},      // json encode 失敗
	"JSON_UNMARSHAL_ERROR":    {1006, "JSON UNMARSHAL ERROR"},    // json decode 失敗
	"CHANGE_PARAMS_TYPE_FAIL": {1007, "CHANGE PARAMS TYPE FAIL"}, // 資料轉型失敗
	"PARSE_TIME_ERROR":        {1008, "PARSE TIME ERROR"},        // 時間格式轉換錯誤
	"VAILDATE_PARAMS_FAIL":    {1009, "VAILDATE PARAMS FAIL"},    // 參數型態驗證失敗
	"IMAGES_TOO_LARGE":        {1010, "IMAGES TOO LARGE"},        // 檔案過大
	"BIND_PARAMS_FAIL":        {1011, "BIND PARAMS FAIL"},        // 組合參數失敗
	"CRYPTION_ERROR":          {1012, "CRYPTION ERROR"},          // 密碼加密錯誤

	/** DB 錯誤 **/
	"DB_CONNECT_ERROR":                     {2000, "DB CONNECT ERROR"},                     // DB連線失敗
	"CREATE_ADMIN_USER_ERROR":              {2001, "CREATE ADMIN USER ERROR"},              // 新增管理者失敗
	"GET_ADMIN_LIST_ERROR":                 {2002, "GET ADMIN LIST ERROR"},                 // 取管理者清單錯誤
	"DELETE_SESSION_ERROR":                 {2003, "DELETE SESSION ERROR"},                 // 登出時，清除DB session資料失敗
	"UPDATE_PASSWORD_ERROR":                {2004, "UPDATE PASSWORD ERROR"},                // 更改密碼失敗
	"UPDATE_ADMIN_DATA_ERROR":              {2005, "UPDATE ADMIN DATA ERROR"},              // 更新管理者帳號權限失敗
	"GET_USER_ACCOUNT_ERROR":               {2006, "GET USER ACCOUNT ERROR"},               // 查詢帳號是否存在失敗
	"DELETE_ADMIN_ERROR":                   {2007, "DELETE ADMIN ERROR"},                   // 刪除管理者帳號失敗
	"DELETE_EXPIRE_SESSION_ERROR":          {2008, "DELETE EXPIRE SESSION ERROR"},          // 清理DB過期session錯誤
	"DELETE_ROTATE_ERROR":                  {2009, "DELETE ROTATE ERROR"},                  // 刪除輪播圖失敗
	"CREATE_SESSION_ERROR":                 {2010, "CREATE SESSION ERROR"},                 // 建立session資料失敗
	"UPDATE_LOGIN_TIME_ERROR":              {2011, "UPDATE LOGIN TIME ERROR"},              // 更新登入時間失敗
	"GET_ROTATE_LIST_ERROR":                {2012, "GET ROTATE LIST ERROR"},                // 取得輪播圖失敗
	"GET_VIDEO_LIST_ERROR":                 {2013, "GET VIDEO LIST ERROR"},                 // 取影片清單錯誤
	"DELETE_STYLE_ERROR":                   {2014, "DELETE STYLE ERROR"},                   // 刪除風格資料錯誤
	"CREATE_MARQUEE_ERROR":                 {2015, "CREATE MARQUEE ERROR"},                 // 建立跑馬燈失敗
	"CHECK_MARQUEE_EXISTS_ERROR":           {2016, "CHECK MARQUEE EXISTS ERROR"},           // 確認跑馬燈ID是否存在失敗
	"CREATE_OR_UPDATE_MARQUEE_ERROR":       {2017, "CREATE OR UPDATE MARQUEE ERROR"},       // 新增/編輯跑馬燈失敗
	"CREATE_OR_UPDATE_VIDEO_ERROR":         {2018, "CREATE OR UPDATE VIDEO ERROR"},         // 新增/編輯宣除影片失敗
	"GET_CATEGORY_LIST_ERROR":              {2019, "GET CATEGORY LIST ERROR"},              // 取類別清單錯誤
	"GET_MARQUEE_LIST_ERROR":               {2020, "GET MARQUEE LIST ERROR"},               // 取跑馬燈清單錯誤
	"DELETE_MARQUEE_ERROR":                 {2021, "DELETE MARQUEE ERROR"},                 // 刪除跑馬燈錯誤
	"CREATE_OR_UPDATE_TAG_ERROR":           {2022, "CREATE OR UPDATE TAG ERROR"},           // 更改或新增標籤錯誤
	"DELETE_TAG_ERROR":                     {2023, "DELETE TAG ERRORR"},                    // 刪除標籤錯誤
	"GET_ACTIVE_LIST_ERROR":                {2024, "GET ACTIVE LIST ERROR"},                // 取活動清單錯誤
	"CHECK_VIDEO_EXISTS_ERROR":             {2025, "CHECK VIDEO EXISTS ERROR"},             // 確認宣傳影片ID是否存在失敗
	"CHECK_TAG_EXISTS_ERROR":               {2026, "CHECK TAG EXISTS ERROR"},               // 檢查標籤是否存在失敗
	"CREATE_STYLE_ERROR":                   {2027, "CREATE STYLE ERROR"},                   //  建立風個資料錯誤
	"DELETE_VIDEO_ERROR":                   {2028, "DELETE_VIDEO_ERROR"},                   //  刪除宣傳影片失敗
	"CHANGE_VIDEO_STATUS_ERROR":            {2029, "CHANGE VIDEO STATUS ERROR"},            //  宣傳影片狀態修改失敗
	"DELETE_NEWS_ERROR":                    {2030, "DELETE NEWS ERRORR"},                   // 刪除報導錯誤
	"DELETE_ACTIVE_ERROR":                  {2031, "DELETE ACTIVE ERROR"},                  // 刪除活動錯誤
	"CREATE_OR_UPDATE_STYLE_SETTING_ERROR": {2032, "CREATE OR UPDATE STYLE SETTING ERROR"}, // 調整風格設定錯誤
	"CHECK_ACTIVE_EXISTS_ERROR":            {2033, "CHECK ACTIVE EXISTS ERROR"},            // 檢查活動是否存在失敗
	"CHECK_NEWS_EXISTS_ERROR":              {2034, "CHECK NEWS EXISTS ERROR"},              // 檢查報導是否存在失敗
	"CREATE_OR_UPDATE_ACTIVE_ERROR":        {2035, "CREATE OR UPDATE ACTIVE ERROR"},        // 新增或更新活動失敗
	"CREATE_OR_UPDATE_ROTATE_ERROR":        {2036, "CREATE OR UPDATE ROTATE ERROR"},        // 新增或更新輪播圖失敗
	"CHECK_ROTATE_EXISTS_ERROR":            {2037, "CHECK ROTATE EXISTS ERROR"},            // 檢查輪播圖是否存在失敗
	"GET_STYLE_LIST_ERROR":                 {2038, "GET STYLE LIST ERROR"},                 // 取風格清單失敗
	"GET_STYLE_SETTING_ERROR":              {2039, "GET STYLE SETTING ERROR"},              // 取風格設定值失敗
	"CHANGE_STYLE_STATUS_ERROR":            {2040, "CHANGE STYLE STATUS ERROR"},            // 更改風格開關失敗
	"CHECK_CATEGORY_EXISTS_ERROR":          {2041, "CHECK CATEGORY EXISTS ERROR"},          // 檢查類別是否存在失敗
	"CREATE_OR_UPDATE_CATEGORY_ERROR":      {2042, "CREATE OR UPDATE CATEGORY ERROR"},      // 新增或更新類別失敗

	/** Redis 錯誤 **/
	"REDIS_CONNECT_ERROR":     {3000, "REDIS CONNECT ERROR"},     // Redis連線失敗
	"REDIS_INSERT_ERROR":      {3001, "REDIS INSERT ERROR"},      // Redis寫入失敗
	"REDIS_DELETE_ERROR":      {3002, "REDIS DELETE ERROR"},      // Redis刪除失敗
	"REDIS_APPEND_ERROR":      {3003, "REDIS APPEND ERROR"},      // Redis增加值失敗
	"REDIS_SET_EXPIRE_ERROR":  {3004, "REDIS SET EXPIRE ERROR"},  // Redis設定過期時間失敗
	"REDIS_CHECK_EXIST_ERROR": {3005, "REDIS CHECK EXIST ERROR"}, // 檢查Redis值是否存在時發生錯誤

	/** 呼叫 API 錯誤 **/
	"CURL_CREATE_FAIL":  {4000, "CURL CREATE FAIL"},  // CURL建立失敗
	"CURL_GET_FAIL":     {4001, "CURL GET FAIL"},     // CURL GET 失敗
	"CURL_POST_FAIL":    {4002, "CURL POST FAIL"},    // 取API失敗
	"API_CONNECT_ERROR": {4003, "API CONNECT ERROR"}, // 對外連線失敗
	"API_STATUS_ERROR":  {4004, "API STATUS ERROR"},  // 對外連線回傳code異常

	/** 其他 **/
	"AUTH_VAILDATE_FAIL":            {5000, "AUTH VAILDATE FAIL"},            // 登入驗證失敗
	"USER_OR_PASSWORD_ERROR":        {5001, "USER OR PASSWORD ERROR"},        // 用戶不存在
	"USER_ACCOUNT_DISABLE":          {5002, "USER ACCOUNT DISABLE"},          // 用戶帳號遭凍結
	"ACCOUNT_RULE_ERROR":            {5003, "ACCOUNT RULE ERROR"},            // 帳號內容不符合正則規則
	"PASSWORD_RULE_ERROR":           {5004, "PASSWORD RULE ERROR"},           // 密碼內容不符合正則規則
	"USER_IS_EXIST":                 {5005, "USER IS EXIST"},                 // 用戶已經存在，不可重複註冊
	"PASSWORD_CAN_NOT_BE_THE_SAME":  {5006, "PASSWORD CAN NOT BE THE SAME"},  // 更改自己密碼時，不可以輸入相同密碼
	"GET_TOPGAME_API_ERROR":         {5007, "GET TOPGAME API ERROR"},         // 取得CY topgame API錯誤
	"MARQUEE_ID_NOT_EXISTS":         {5008, "MARQUEE ID NOT EXISTS"},         // 跑馬燈不存在
	"VIDEO_ID_NOT_EXISTS":           {5009, "VIDEO ID NOT EXISTS"},           // 宣傳影片不存在
	"ROTATE_ID_NOT_EXISTS":          {5010, "ROTATE ID NOT EXISTS"},          // 輪播圖不存在
	"TAG_ID_NOT_EXISTS":             {5011, "TAG ID NOT EXISTS"},             //  標籤不存在
	"NEWS_ID_NOT_EXISTS":            {5012, "NEWS ID NOT EXISTS"},            // 報導不存在
	"ACTIVE_ID_NOT_EXISTS":          {5013, "ACTIVE ID NOT EXISTS"},          // 活動不存在
	"NO_DATA_BE_AFFECTED":           {5014, "NO DATA BE AFFECTED"},           //  沒有資料被異動
	"NO_MATCHE_DATA_FOUND":          {5015, "NO MATCHE DATA FOUND"},          // 沒有找到批配資料
	"SESSION_NOT_EXIST":             {5016, "SESSION NOT EXIST"},             // DB Session資料不存在
	"CATEGORY_ID_NOT_EXISTS":        {5018, "CATEGORY ID NOT EXISTS"},        // 類別不存在
	"GET_TOPREWARD_API_ERROR":       {5019, "GET TOPREWARD API ERROR"},       // 取得CY topreward API錯誤
	"GET_TOPLUCKY_API_ERROR":        {5020, "GET TOPLUCKY API ERROR"},        // 取得CY toplucky API錯誤
	"GET_TOPWIN_API_ERROR":          {5021, "GET TOPWIN API ERROR"},          // 取得CY topwin API錯誤
	"GET_GAME_NAME_LIST_API_ERROR":  {5022, "GET GAME NAME LIST API ERROR"},  // 取得CY game name list API錯誤
	"GET_MARQUEE_SET_API_ERROR":     {5023, "GET MARQUEE SET API ERROR"},     // 取得 跑馬燈贏分條件與倍率條件 失敗
	"SET_MARQUEE_SET_API_ERROR":     {5024, "SET MARQUEE SET API ERROR"},     // 設定 跑馬燈贏分條件與倍率條件 失敗
	"GET_AUTO_MARQUEE_API_ERROR":    {5025, "GET AUTO MARQUEE API ERROR"},    // 取跑馬燈自動清單失敗
	"GET_MARQUEE_GAME_API_ERROR":    {5026, "GET MARQUEE GAME API ERROR"},    // 取允許進跑馬燈的遊戲失敗
	"CREATE_MARQUEE_GAME_API_ERROR": {5027, "CREATE MARQUEE GAME API ERROR"}, // 新增允許進跑馬燈的遊戲失敗
	"DELETE_MARQUEE_GAME_API_ERROR": {5028, "DELETE MARQUEE GAME API ERROR"}, // 刪除允許進跑馬燈的遊戲失敗
	"GET_ORDER_DETAIL_API_ERROR":    {5029, "GET ORDER DETAIL API ERROR"},    // 取細單失敗
	"TOKEN_EXPIRED":                 {5030, "TOKEN_EXPIRED"},                 // round id token 過期
	"CAN_NOT_ENTER_PREVIOUS_DATE":   {5031, "CAN NOT ENTER PREVIOUS DATE"},   // 禁止輸入當天以前的日期
}

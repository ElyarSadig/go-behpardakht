package bankerrors

import "errors"

var (
	ErrInvalidCardNumber            = errors.New("شماره کارت نامعتبر است")
	ErrInsufficientFunds            = errors.New("موجودی کافی نیست")
	ErrIncorrectPassword            = errors.New("رمز نادرست است")
	ErrExceededPasswordAttempts     = errors.New("تعداد دفعات وارد کردن رمز بیش از حد مجاز است")
	ErrInvalidCard                  = errors.New("کارت نامعتبر است")
	ErrExceededWithdrawalAttempts   = errors.New("دفعات برداشت وجه بیش از حد مجاز است")
	ErrUserCancelledTransaction     = errors.New("کاربر از انجام تراکنش منصرف شده است")
	ErrExpiredCard                  = errors.New("تاریخ انقضای کارت گذشته است")
	ErrExceededWithdrawalAmount     = errors.New("مبلغ برداشت وجه بیش از حد مجاز است")
	ErrInvalidCardIssuer            = errors.New("صادر کننده کارت نامعتبر است")
	ErrCardIssuerSwitchError        = errors.New("خطای سوییچ صادر کننده کارت")
	ErrNoResponseFromIssuer         = errors.New("پاسخی از صادر کننده کارت دریافت نشد")
	ErrTransactionNotAllowed        = errors.New("دارنده این کارت مجاز به انجام این تراکنش نیست")
	ErrInvalidMerchant              = errors.New("پذیرنده نامعتبر است")
	ErrSecurityError                = errors.New("خطای امنیتی رخ داده است")
	ErrInvalidMerchantInfo          = errors.New("اطلاعات کاربری پذیرنده نامعتبر است")
	ErrInvalidAmount                = errors.New("مبلغ نامعتبر است")
	ErrInvalidResponse              = errors.New("پاسخ نامعتبر است")
	ErrIncorrectFormat              = errors.New("فرمت اطلاعات وارد شده صحیح نمی‌ باشد")
	ErrInvalidAccount               = errors.New("حساب نامعتبر است")
	ErrSystemError                  = errors.New("خطای سیستمی")
	ErrInvalidDate                  = errors.New("تاریخ نامعتبر است")
	ErrDuplicateRequest             = errors.New("شماره درخواست تکراری است")
	ErrSaleTransactionNotFound      = errors.New("تراکنش Sale یافت نشد")
	ErrAlreadyVerified              = errors.New("قبلا درخواست Verify داده شده است")
	ErrVerifyRequestNotFound        = errors.New("درخواست Verify یافت نشد")
	ErrTransactionSettled           = errors.New("تراکنش Settle شده است")
	ErrTransactionNotSettled        = errors.New("تراکنش Settle نشده است")
	ErrSettlementNotFound           = errors.New("تراکنش Settle یافت نشد")
	ErrTransactionReversed          = errors.New("تراکنش Reverse شده است")
	ErrRefundNotFound               = errors.New("تراکنش Refund یافت نشد")
	ErrInvalidBillID                = errors.New("شناسه قبض نادرست است")
	ErrInvalidPaymentID             = errors.New("شناسه پرداخت نادرست است")
	ErrInvalidBillIssuer            = errors.New("سازمان صادر کننده قبض نامعتبر است")
	ErrTransactionTimeout           = errors.New("مدت زمان مجاز برای انجام تراکنش به پایان رسیده است")
	ErrDataRegistrationError        = errors.New("خطا در ثبت اطلاعات")
	ErrInvalidPayerID               = errors.New("شناسه پرداخت کننده نامعتبر است")
	ErrCustomerDataError            = errors.New("اشکال در تعریف اطلاعات مشتری")
	ErrExceededInputAttempts        = errors.New("تعداد دفعات ورود اطلاعات از حد مجاز گذشته است")
	ErrInvalidIP                    = errors.New("IP نامعتبر است")
	ErrDuplicateTransaction         = errors.New("تراکنش تکراری است")
	ErrReferenceTransactionNotFound = errors.New("تراکنش مرجع موجود نیست")
	ErrInvalidTransaction           = errors.New("تراکنش نامعتبر است")
	ErrDepositError                 = errors.New("خطا در واریز")
	ErrReturnPathError              = errors.New("مسير بازگشت به سايت در دامنه ثبت شده برای پذيرنده قرار ندارد")
	ErrStaticPasswordLimitReached   = errors.New("سقف استفاده از رمز ايستا به پايان رسيده است")
	ErrUnknownError                 = errors.New("خطای ناشناخته")
)

var errorMessages = map[int]error{
	0:   nil,
	11:  ErrInvalidCardNumber,
	12:  ErrInsufficientFunds,
	13:  ErrIncorrectPassword,
	14:  ErrExceededPasswordAttempts,
	15:  ErrInvalidCard,
	16:  ErrExceededWithdrawalAttempts,
	17:  ErrUserCancelledTransaction,
	18:  ErrExpiredCard,
	19:  ErrExceededWithdrawalAmount,
	111: ErrInvalidCardIssuer,
	112: ErrCardIssuerSwitchError,
	113: ErrNoResponseFromIssuer,
	114: ErrTransactionNotAllowed,
	21:  ErrInvalidMerchant,
	23:  ErrSecurityError,
	24:  ErrInvalidMerchantInfo,
	25:  ErrInvalidAmount,
	31:  ErrInvalidResponse,
	32:  ErrIncorrectFormat,
	33:  ErrInvalidAccount,
	34:  ErrSystemError,
	35:  ErrInvalidDate,
	41:  ErrDuplicateRequest,
	42:  ErrSaleTransactionNotFound,
	43:  ErrAlreadyVerified,
	44:  ErrVerifyRequestNotFound,
	45:  ErrTransactionSettled,
	46:  ErrTransactionNotSettled,
	47:  ErrSettlementNotFound,
	48:  ErrTransactionReversed,
	49:  ErrRefundNotFound,
	412: ErrInvalidBillID,
	413: ErrInvalidPaymentID,
	414: ErrInvalidBillIssuer,
	415: ErrTransactionTimeout,
	416: ErrDataRegistrationError,
	417: ErrInvalidPayerID,
	418: ErrCustomerDataError,
	419: ErrExceededInputAttempts,
	421: ErrInvalidIP,
	51:  ErrDuplicateTransaction,
	54:  ErrReferenceTransactionNotFound,
	55:  ErrInvalidTransaction,
	61:  ErrDepositError,
	62:  ErrReturnPathError,
	98:  ErrStaticPasswordLimitReached,
}

func GetBankErrorMessage(code int) error {
	if err, found := errorMessages[code]; found {
		return err
	}
	return ErrUnknownError
}

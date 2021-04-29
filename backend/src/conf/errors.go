package conf

import (
	"fmt"

	"github.com/pkg/errors"
)

// Error Code
// 役割：出力するエラーメッセージのインデックス（目次）
const (
	_ = iota
	ErrBadRequest
	ErrRecordNotFound
	ErrExistSameName
	ErrUsedDesignTemplate
	ErrFailedToServer
	ErrExclusionControl
	ErrUnauthorized
	ErrUsedPlacementRelation
	ErrStandardTemplate
	ErrExistSameAdID
	ErrUsedDeviceModel
	ErrUsedDeviceModelGroup
	ErrUsedOS
	ErrUsedOSGroup
	ErrUsedSDKVersion
	ErrExistSameOuterDealID
	ErrChangeOuterDealID
	ErrUnsupportedPMP
	ErrUsedDeal
	ErrUnsupportedAutomaticRatio
	ErrChangeAdFormat
	ErrChangeDesignTemplateType
	ErrChangeMediaType
	ErrChangePlatform
	ErrExistSameTemplateTypePlatformAdFormat
	ErrNoPermission
	ErrExistEncryptKeyPlatform
)

// Error Message
// 役割：出力するエラーメッセージ内容
var errorText = map[int]string{
	ErrBadRequest:                            "不正な値が設定されています",
	ErrRecordNotFound:                        "データが削除されているか存在しません",
	ErrExistSameName:                         "同じ名前は登録できません",
	ErrUsedDesignTemplate:                    "配信セットで使用されているため、削除できません",
	ErrFailedToServer:                        "予期せぬエラーが発生しました",
	ErrExclusionControl:                      "他のユーザーにより更新されました。再度更新してやり直して下さい",
	ErrUnauthorized:                          "認証エラー",
	ErrUsedPlacementRelation:                 "配信セットで使用されているため、削除できません",
	ErrStandardTemplate:                      "標準テンプレートは、編集/削除できません",
	ErrExistSameAdID:                         "同じADID/IDFAは登録できません",
	ErrUsedDeviceModel:                       "機種グループで使用されているため、削除できません",
	ErrUsedDeviceModelGroup:                  "配信セットで使用されているため、削除できません",
	ErrUsedOS:                                "OSグループで使用されているため、削除できません",
	ErrUsedOSGroup:                           "配信セットまたはプレースメント設定で使用されているため、削除できません",
	ErrUsedSDKVersion:                        "SDKバージョングループで使用されているため、削除できません",
	ErrExistSameOuterDealID:                  "同じDealIDは登録できません",
	ErrChangeOuterDealID:                     "編集時にDealIDの変更はできません",
	ErrUnsupportedPMP:                        "PMP非対応のPFにDealの選択はできません",
	ErrUsedDeal:                              "配信セットで使用されているため、削除できません",
	ErrUnsupportedAutomaticRatio:             "配信比率自動化非対応のPFの選択はできません",
	ErrChangeAdFormat:                        "編集時にフォーマットの変更はできません",
	ErrChangeDesignTemplateType:              "編集時に配信手法の変更はできません",
	ErrChangeMediaType:                       "編集時に配信面の変更はできません",
	ErrExistEncryptKeyPlatform:               "既にkey発行済のプラットフォームです",
	ErrChangePlatform:                        "編集時にプラットフォームの変更はできません",
	ErrExistSameTemplateTypePlatformAdFormat: "同じ配信手法/プラットフォーム/フォーマットは登録できません",
	ErrNoPermission:                          "登録編集する権限がありません",
}

// ErrorText ErrorText関数
// 役割：受け取ったError Codeのインデックス番号に対応したError Messageを返す
func ErrorText(code int) string {
	return errorText[code]
}

// AppError AppError構造体
// 役割：カスタムエラー用の構造体
type AppError struct {
	Code int
	Err  error
}

// NewAppError NewAppError関数
// 役割：code（エラーメッセージのインデックス番号）とerrを元にAppError型のオブジェクトを生成する
func NewAppError(code int, err error) *AppError {
	if tae, ok := errors.Cause(err).(*AppError); ok {
		code = tae.Code
	}
	return &AppError{Code: code, Err: err}
}

// Error Errorメソッド
// 役割：カスタムエラー用のメソッド。自作のAppError型には当然デフォルトではErrorメソッドがなく、
// type assertionの結果に応じてerrorを返すためには、AppError型に「Error() string」という形のメソッドを作成し、
// カスタムエラーを設定する必要がある。
// 参考：https://qiita.com/romukey/items/e49e28b7dcf645ac91c7#%E3%81%8A%E3%81%BE%E3%81%91---%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%A0%E3%82%A8%E3%83%A9%E3%83%BC%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9
// 参考：https://hawksnowlog.blogspot.com/2019/07/access-fields-of-golang-error.html
// 参考：http://akito0107.hatenablog.com/entry/2018/12/14/112717
func (e *AppError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

// Wrap Wrapメソッド
// 役割：任意のエラーメッセージをラップして新規エラーを返す
func (e *AppError) Wrap() (err error) {
	return NewAppError(e.Code, fmt.Errorf(ErrorText(e.Code)))
}
